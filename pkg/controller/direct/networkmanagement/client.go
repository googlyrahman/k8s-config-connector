// Copyright 2024 Google LLC
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

// +tool:controller-client
// proto.service: google.cloud.networkmanagement.v1.ReachabilityService

package networkmanagement

import (
	"context"
	"fmt"

	api "cloud.google.com/go/networkmanagement/apiv1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
)

func newReachabilityClient(ctx context.Context, config *config.ControllerConfig) (*api.ReachabilityClient, error) {
	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := api.NewReachabilityRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building networkmanagement reachability client: %w", err)
	}
	return client, err
}
