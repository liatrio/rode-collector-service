package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/rode/rode/auth"
	"github.com/rode/rode/config"
	pb "github.com/rode/rode/proto/v1alpha1"
	grafeas_proto "github.com/rode/rode/protodeps/grafeas/proto/v1beta1/grafeas_go_proto"
	grafeas_project_proto "github.com/rode/rode/protodeps/grafeas/proto/v1beta1/project_go_proto"
	"github.com/rode/rode/server"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func main() {
	c, err := config.Build(os.Args[0], os.Args[1:])
	if err != nil {
		log.Fatalf("failed to build config: %v", err)
	}

	logger, err := createLogger(c.Debug)
	if err != nil {
		log.Fatalf("failed to create logger: %v", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", c.Port))
	if err != nil {
		logger.Fatal("failed to listen", zap.Error(err))
	}

	grafeasClientCommon, grafeasClientProjects, err := createGrafeasClients(c.Grafeas.Host)
	if err != nil {
		logger.Fatal("failed to connect to grafeas", zap.String("grafeas host", c.Grafeas.Host), zap.Error(err))
	}

	authenticator := auth.NewAuthenticator(c.Auth)
	s := grpc.NewServer(
		grpc.StreamInterceptor(
			grpc_auth.StreamServerInterceptor(authenticator.Authenticate),
		),
		grpc.UnaryInterceptor(
			grpc_auth.UnaryServerInterceptor(authenticator.Authenticate),
		),
	)
	if c.Debug {
		reflection.Register(s)
	}

	rodeServer, err := server.NewRodeServer(logger.Named("rode"), grafeasClientCommon, grafeasClientProjects)
	if err != nil {
		logger.Fatal("failed to create Rode server", zap.Error(err))
	}
	healthzServer := server.NewHealthzServer(logger.Named("healthz"))

	pb.RegisterRodeServer(s, rodeServer)
	grpc_health_v1.RegisterHealthServer(s, healthzServer)

	go func() {
		if err := s.Serve(lis); err != nil {
			logger.Fatal("failed to serve", zap.Error(err))
		}
	}()

	logger.Info("starting grpc gateway", zap.String("host", lis.Addr().String()))
	go func() {
		if err := createGrpcGateway(context.Background(), lis.Addr().String(), ":50052"); err != nil {
			logger.Fatal("failed to start gateway", zap.Error(err))
		}
	}()

	logger.Info("listening", zap.String("host", lis.Addr().String()))
	healthzServer.Ready()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	terminationSignal := <-sig

	logger.Info("shutting down...", zap.String("termination signal", terminationSignal.String()))
	healthzServer.NotReady()

	s.GracefulStop()
}

func createGrafeasClients(grafeasEndpoint string) (grafeas_proto.GrafeasV1Beta1Client, grafeas_project_proto.ProjectsClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	connection, err := grpc.DialContext(ctx, grafeasEndpoint, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, nil, err
	}

	grafeasClient := grafeas_proto.NewGrafeasV1Beta1Client(connection)
	projectsClient := grafeas_project_proto.NewProjectsClient(connection)

	return grafeasClient, projectsClient, nil
}

func createGrpcGateway(ctx context.Context, grpcAddress, httpPort string) error {
	conn, err := grpc.DialContext(
		context.Background(),
		grpcAddress,
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
	gwmux := runtime.NewServeMux()
	if err := pb.RegisterRodeHandler(ctx, gwmux, conn); err != nil {
		return err
	}

	srv := &http.Server{
		Addr: httpPort,
		Handler: gwmux,
	}

	srv.ListenAndServe()

	return nil
}

func createLogger(debug bool) (*zap.Logger, error) {
	if debug {
		return zap.NewDevelopment()
	}

	return zap.NewProduction()
}
