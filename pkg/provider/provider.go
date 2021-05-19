package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rode/es-index-manager/indexmanager"
	"github.com/rode/grafeas-elasticsearch/go/v1beta1/storage/esutil"
	"github.com/rode/grafeas-elasticsearch/go/v1beta1/storage/filtering"
	"github.com/rode/rode/auth"
	"github.com/rode/rode/config"
	"github.com/rode/rode/opa"
	"github.com/rode/rode/pkg/resource"
	pb "github.com/rode/rode/proto/v1alpha1"
	grafeas_proto "github.com/rode/rode/protodeps/grafeas/proto/v1beta1/grafeas_go_proto"
	grafeas_project_proto "github.com/rode/rode/protodeps/grafeas/proto/v1beta1/project_go_proto"
	"github.com/rode/rode/server"
	"github.com/soheilhy/cmux"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
	"os"
	"time"
)

func NewConfig() (*config.Config, error) {
	return config.Build(os.Args[0], os.Args[1:])
}

func NewLogger(c *config.Config) (*zap.Logger, error) {
	if c.Debug {
		return zap.NewDevelopment()
	}

	return zap.NewProduction()
}

func NewEsClient(lc fx.Lifecycle, c *config.Config, logger *zap.Logger) (*elasticsearch.Client, error) {
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{
			c.Elasticsearch.Host,
		},
		Username: c.Elasticsearch.Username,
		Password: c.Elasticsearch.Password,
	})

	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			res, err := client.Info()

			if err != nil {
				return err
			}

			var r struct {
				Version struct {
					Number string `json:"number"`
				} `json:"string"`
			}
			if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
				return err
			}

			logger.Debug("Successful Elasticsearch connection", zap.String("ES Server version", r.Version.Number))

			return nil
		},
	})

	return client, nil
}

func NewGrafeasClients(c *config.Config) (grafeas_proto.GrafeasV1Beta1Client, grafeas_project_proto.ProjectsClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	connection, err := grpc.DialContext(ctx, c.Grafeas.Host, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, nil, err
	}

	grafeasClient := grafeas_proto.NewGrafeasV1Beta1Client(connection)
	projectsClient := grafeas_project_proto.NewProjectsClient(connection)

	return grafeasClient, projectsClient, nil
}

type CmuxResult struct {
	fx.Out
	Mux cmux.CMux

	GrpcListener net.Listener `name:"grpc"`
	HttpListener net.Listener `name:"http"`
}

func NewCmux(lc fx.Lifecycle, c *config.Config) (CmuxResult, error) {
	address := fmt.Sprintf(":%d", c.Port)
	// TODO: this is a side effect, would ideally happen during an OnStart hook
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return CmuxResult{}, err
	}

	mux := cmux.New(lis)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return mux.Serve()
		},
	})

	return CmuxResult{
		Mux:          mux,
		GrpcListener: mux.Match(cmux.HTTP2()),
		HttpListener: mux.Match(cmux.HTTP1()),
	}, nil
}

type grpcParams struct {
	fx.In
	Config       *config.Config
	Logger       *zap.Logger
	GrpcListener net.Listener `name:"grpc"`
}

func NewGrpcServer(lc fx.Lifecycle, params grpcParams) *grpc.Server {
	authenticator := auth.NewAuthenticator(params.Config.Auth)

	s := grpc.NewServer(
		grpc.StreamInterceptor(
			grpc_auth.StreamServerInterceptor(authenticator.Authenticate),
		),
		grpc.UnaryInterceptor(
			grpc_auth.UnaryServerInterceptor(authenticator.Authenticate),
		),
	)

	if params.Config.Debug {
		reflection.Register(s)
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return s.Serve(params.GrpcListener)
		},
		OnStop: func(ctx context.Context) error {
			s.GracefulStop()

			return nil
		},
	})

	return s
}

type gatewayParams struct {
	fx.In
	Config       *config.Config
	Logger       *zap.Logger
	HttpListener net.Listener `name:"http"`
}

// startup: call pb.RegisterRodeHandler(ctx, gwmux, conn)
func NewGrpcGateway(lc fx.Lifecycle, params gatewayParams) (*http.Server, error) {
	address := fmt.Sprintf(":%d", params.Config.Port)
	conn, err := grpc.DialContext(
		context.Background(),
		address,
		grpc.WithInsecure(),
	)

	if err != nil {
		return nil, err
	}

	gwmux := runtime.NewServeMux()
	httpMux := http.NewServeMux()
	httpMux.Handle("/", http.Handler(gwmux))

	httpServer := &http.Server{
		Handler: httpMux,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if err := pb.RegisterRodeHandler(ctx, gwmux, conn); err != nil {
				return err
			}

			if err := httpServer.Serve(params.HttpListener); err != http.ErrServerClosed {
				return err
			}

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return httpServer.Shutdown(ctx)
		},
	})

	return httpServer, nil
}

func NewHealthzServer(lc fx.Lifecycle, l *zap.Logger) server.HealthzServer {
	s := server.NewHealthzServer(l.Named("healthz"))

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			s.NotReady() // TODO: does this need to be called earlier?

			return nil
		},
	})

	return s
}

func NewEsutilClient(logger *zap.Logger, client *elasticsearch.Client) esutil.Client {
	return esutil.NewClient(logger.Named("ESClient"), client)
}

func NewIndexManager(logger *zap.Logger, client *elasticsearch.Client) indexmanager.IndexManager {
	return indexmanager.NewIndexManager(logger.Named("IndexManager"), client, &indexmanager.Config{
		IndexPrefix:  "rode",
		MappingsPath: "mappings",
	})
}

func NewResourceManager(
	c *config.Config,
	logger *zap.Logger,
	client esutil.Client,
	indexManager indexmanager.IndexManager,
) resource.Manager {
	return resource.NewManager(
		logger.Named("Resource Manager"),
		client,
		c.Elasticsearch,
		indexManager,
	)
}

func NewOpaClient(c *config.Config, logger *zap.Logger) opa.Client {
	return opa.NewClient(logger.Named("opa"), c.Opa.Host, c.Debug)
}

type rodeParams struct {
	fx.In
	Config          *config.Config
	Logger          *zap.Logger
	GrafeasCommon   grafeas_proto.GrafeasV1Beta1Client
	GrafeasProjects grafeas_project_proto.ProjectsClient
	Opa             opa.Client
	EsClient        *elasticsearch.Client
	//Filterer filtering.Filterer,
	//ElasticsearchConfig *config.ElasticsearchConfig
	ResourceManager resource.Manager
	IndexManager    indexmanager.IndexManager
}

func NewRodeServer(lc fx.Lifecycle, params rodeParams) (pb.RodeServer, error) {
	s, err := server.NewRodeServer(
		params.Logger,
		params.GrafeasCommon,
		params.GrafeasProjects,
		params.Opa,
		params.EsClient,
		filtering.NewFilterer(),
		params.Config.Elasticsearch,
		params.ResourceManager,
		params.IndexManager,
	)

	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// TODO: don't cast this
			return s.(*server.RodeServer).Initialize(ctx)
		},
	})

	return s, nil
}
