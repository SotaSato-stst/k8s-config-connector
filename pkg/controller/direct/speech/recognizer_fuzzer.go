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
// proto.message: google.cloud.speech.v2.Recognizer
// api.group: speech.cnrm.cloud.google.com

package speech

import (
	pb "cloud.google.com/go/speech/apiv2/speechpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(speechRecognizerFuzzer())
}

func speechRecognizerFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Recognizer{},
		SpeechRecognizerSpec_FromProto, SpeechRecognizerSpec_ToProto,
		SpeechRecognizerObservedState_FromProto, SpeechRecognizerObservedState_ToProto,
	)

	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".default_recognition_config")
	f.SpecFields.Insert(".annotations")

	f.StatusFields.Insert(".uid")
	f.StatusFields.Insert(".default_recognition_config")
	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".delete_time")
	f.StatusFields.Insert(".expire_time")
	f.StatusFields.Insert(".etag")
	f.StatusFields.Insert(".reconciling")
	f.StatusFields.Insert(".kms_key_name")
	f.StatusFields.Insert(".kms_key_version_name")

	f.UnimplementedFields.Insert(".name") // special field
	f.UnimplementedFields.Insert(".language_codes")
	f.UnimplementedFields.Insert(".model")

	return f
}
