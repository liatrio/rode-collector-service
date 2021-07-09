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

package common

import "flag"

func SetupRodeClientFlags(flags *flag.FlagSet) *ClientConfig {
	conf := &ClientConfig{
		Rode:      &RodeClientConfig{},
		JWTAuth:   &JWTAuthConfig{},
		BasicAuth: &BasicAuthConfig{},
	}

	flags.StringVar(&conf.Rode.Host, "rode-host", "rode:50051", "the host to use to connect to rode")
	flags.BoolVar(&conf.Rode.DisableTransportSecurity, "rode-insecure-disable-transport-security", false, "when set, the connection to rode will not use transport security")

	flags.StringVar(&conf.JWTAuth.ClientID, "jwt-client-id", "", "the client ID to use when requesting a JWT via the client_credentials grant")
	flags.StringVar(&conf.JWTAuth.ClientSecret, "jwt-client-secret", "", "the client secret to use when requesting a JWT via the client_credentials grant")
	flags.StringVar(&conf.JWTAuth.TokenURL, "jwt-token-url", "", "the URL to use to retrieve an access token via the client_credentials grant")

	flags.StringVar(&conf.BasicAuth.Username, "basic-auth-username", "", "the username to use for basic authentication")
	flags.StringVar(&conf.BasicAuth.Password, "basic-auth-password", "", "the password to use for basic authentication")

	return conf
}