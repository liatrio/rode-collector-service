// Copyright 2021 The Rode Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	"context"
	"crypto/tls"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/coreos/go-oidc"
)

type Config struct {
	Auth          *AuthConfig
	Elasticsearch *ElasticsearchConfig
	Grafeas       *GrafeasConfig
	Opa           *OpaConfig
	Port          int
	Debug         bool
}

type GrafeasConfig struct {
	Host string
}

type OpaConfig struct {
	Host string
}

type AuthConfig struct {
	Basic *BasicAuthConfig
	JWT   *JWTAuthConfig
}

type BasicAuthConfig struct {
	Username string
	Password string
}

type JWTAuthConfig struct {
	Issuer           string
	RequiredAudience string
	Verifier         *oidc.IDTokenVerifier
}

type ElasticsearchConfig struct {
	Host     string
	Username string
	Password string
	Refresh  RefreshOption
}

func (c ElasticsearchConfig) IsValid() (e error) {
	switch c.Refresh {
	case RefreshTrue, RefreshWaitFor, RefreshFalse:
		break
	default:
		return fmt.Errorf("invalid refresh option %s. valid options are \"true\", \"false\", \"wait_for\"", c.Refresh)
	}

	return nil
}

// RefreshOption is based on https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-refresh.html
type RefreshOption string

func (r RefreshOption) String() string {
	return string(r)
}

const (
	RefreshTrue    = "true"
	RefreshWaitFor = "wait_for"
	RefreshFalse   = "false"
)

func Build(name string, args []string) (*Config, error) {
	flags := flag.NewFlagSet(name, flag.ContinueOnError)

	conf := &Config{
		Auth: &AuthConfig{
			Basic: &BasicAuthConfig{},
			JWT:   &JWTAuthConfig{},
		},
		Elasticsearch: &ElasticsearchConfig{},
		Grafeas:       &GrafeasConfig{},
		Opa:           &OpaConfig{},
	}

	flags.StringVar(&conf.Auth.Basic.Username, "basic-auth-username", "", "when set, basic auth will be enabled for all endpoints, using the provided username. --basic-auth-password must also be set")
	flags.StringVar(&conf.Auth.Basic.Password, "basic-auth-password", "", "when set, basic auth will be enabled for all endpoints, using the provided password. --basic-auth-username must also be set")
	flags.StringVar(&conf.Auth.JWT.Issuer, "jwt-issuer", "", "when set, jwt based auth will be enabled for all endpoints. the provided issuer will be used to fetch the discovery document in order to validate received jwts")
	flags.StringVar(&conf.Auth.JWT.RequiredAudience, "jwt-required-audience", "", "when set, if jwt based auth is enabled, this audience must be specified within the `aud` claim of any received jwts")

	flags.IntVar(&conf.Port, "port", 50051, "the port that the rode gRPC/HTTP API server should listen on")
	flags.BoolVar(&conf.Debug, "debug", false, "when set, debug mode will be enabled")
	flags.StringVar(&conf.Grafeas.Host, "grafeas-host", "localhost:8080", "the host to use to connect to grafeas")
	flags.StringVar(&conf.Opa.Host, "opa-host", "http://localhost:8181", "the host to use to connect to Open Policy Agent")

	flags.StringVar(&conf.Elasticsearch.Host, "elasticsearch-host", "http://elasticsearch-master:9200", "the Elasticsearch endpoint used by Grafeas")
	flags.StringVar(&conf.Elasticsearch.Username, "elasticsearch-username", "", "username for the Grafeas Elasticsearch instance")
	flags.StringVar(&conf.Elasticsearch.Password, "elasticsearch-password", "", "password for the Grafeas Elasticsearch instance")
	var elasticsearchRefresh string
	flags.StringVar(&elasticsearchRefresh, "elasticsearch-refresh", "true", "refresh controls when changes made by a request are made visible to search. Options are \"true\", \"false\", \"wait_for\"")

	err := flags.Parse(args)
	if err != nil {
		return nil, err
	}

	conf.Elasticsearch.Refresh = RefreshOption(elasticsearchRefresh)
	if conf.Elasticsearch.IsValid() != nil {
		return nil, conf.Elasticsearch.IsValid()
	}
	if (conf.Auth.Basic.Username != "" && conf.Auth.Basic.Password == "") || (conf.Auth.Basic.Username == "" && conf.Auth.Basic.Password != "") {
		return nil, errors.New("when using basic auth, both --basic-auth-username and --basic-auth-password must be set")
	}

	if (conf.Elasticsearch.Username != "" && conf.Elasticsearch.Password == "") || (conf.Elasticsearch.Username == "" && conf.Elasticsearch.Password != "") {
		return nil, errors.New("if Elasticsearch auth is configured, both --elasticsearch-username and --elasticsearch-password must be set")
	}

	if conf.Auth.JWT.Issuer != "" {
		httpClient := &http.Client{
			Timeout: 30 * time.Second,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		}
		oidcCtx := oidc.ClientContext(context.Background(), httpClient)

		provider, err := oidc.NewProvider(oidcCtx, conf.Auth.JWT.Issuer)
		if err != nil {
			return nil, fmt.Errorf("error initializing oidc provider: %v", err)
		}

		oidcConfig := &oidc.Config{}
		if conf.Auth.JWT.RequiredAudience != "" {
			oidcConfig.ClientID = conf.Auth.JWT.RequiredAudience
		} else {
			oidcConfig.SkipClientIDCheck = true
		}

		conf.Auth.JWT.Verifier = provider.Verifier(oidcConfig)
	} else if conf.Auth.JWT.RequiredAudience != "" {
		return nil, errors.New("the --jwt-required-audience flag cannot be specified without --jwt-issuer")
	}

	return conf, nil
}
