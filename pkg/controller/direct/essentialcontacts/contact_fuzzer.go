// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +tool:fuzz-gen
// proto.message: google.cloud.essentialcontacts.v1.Contact
// api.group: essentialcontacts.cnrm.cloud.google.com

package essentialcontacts

import (
	pb "cloud.google.com/go/essentialcontacts/apiv1/essentialcontactspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(essentialContactsContactFuzzer())
}

func essentialContactsContactFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Contact{},
		EssentialContactsContactSpec_FromProto, EssentialContactsContactSpec_ToProto,
		EssentialContactsContactObservedState_FromProto, EssentialContactsContactObservedState_ToProto,
	)

	f.SpecFields.Insert(".email")
	f.SpecFields.Insert(".notification_category_subscriptions")
	f.SpecFields.Insert(".language_tag")
	f.SpecFields.Insert(".validate_time")

	f.StatusFields.Insert(".validation_state")

	f.UnimplementedFields.Insert(".name") // special field
	return f
}
