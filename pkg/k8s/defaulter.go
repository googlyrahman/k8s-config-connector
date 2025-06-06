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

package k8s

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ReconcilerType describes the type of reconciler (controller) being used
type ReconcilerType string

const (
	ReconcilerTypeDirect           ReconcilerType = "direct"
	ReconcilerTypeIAMPolicy        ReconcilerType = "iampolicy"
	ReconcilerTypeIAMPartialPolicy ReconcilerType = "iampartialpolicy"
	ReconcilerTypeIAMPolicyMember  ReconcilerType = "iampolicymember"
	ReconcilerTypeIAMAuditConfig   ReconcilerType = "iamauditconfig"

	ReconcilerTypeTerraform ReconcilerType = "tf"
	ReconcilerTypeDCL       ReconcilerType = "dcl"
)

type Defaulter interface {
	ApplyDefaults(ctx context.Context, reconcilerType ReconcilerType, obj client.Object) (changed bool, err error)
}
