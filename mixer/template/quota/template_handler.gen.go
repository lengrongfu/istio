// Copyright 2017 Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// THIS FILE IS AUTOMATICALLY GENERATED.

package quota

import (
	"context"

	"istio.io/istio/mixer/pkg/adapter"
)

// The `quota` template represents an item for which to check quota.
//
// Example config:
//
// ```
// apiVersion: "config.istio.io/v1alpha2"
// kind: quota
// metadata:
//   name: requestcount
//   namespace: istio-system
// spec:
//   dimensions:
//     source: source.labels["app"] | source.service | "unknown"
//     sourceVersion: source.labels["version"] | "unknown"
//     destination: destination.labels["app"] | destination.service | "unknown"
//     destinationVersion: destination.labels["version"] | "unknown"
// ```

// Fully qualified name of the template
const TemplateName = "quota"

// Instance is constructed by Mixer for the 'quota' template.
//
// The `quota` template represents a piece of data to check Quota for.
//
// When writing the configuration, the value for the fields associated with this template can either be a
// literal or an [expression](https://istio.io/docs//reference/config/policy-and-telemetry/expression-language.html). Please note that if the datatype of a field is not istio.policy.v1beta1.Value,
// then the expression's [inferred type](https://istio.io/docs//reference/config/policy-and-telemetry/expression-language.html#type-checking) must match the datatype of the field.
type Instance struct {
	// Name of the instance as specified in configuration.
	Name string

	// The unique identity of the particular quota to manipulate.
	Dimensions map[string]interface{}
}

// HandlerBuilder must be implemented by adapters if they want to
// process data associated with the 'quota' template.
//
// Mixer uses this interface to call into the adapter at configuration time to configure
// it with adapter-specific configuration as well as all template-specific type information.
type HandlerBuilder interface {
	adapter.HandlerBuilder

	// SetQuotaTypes is invoked by Mixer to pass the template-specific Type information for instances that an adapter
	// may receive at runtime. The type information describes the shape of the instance.
	SetQuotaTypes(map[string]*Type /*Instance name -> Type*/)
}

// Handler must be implemented by adapter code if it wants to
// process data associated with the 'quota' template.
//
// Mixer uses this interface to call into the adapter at request time in order to dispatch
// created instances to the adapter. Adapters take the incoming instances and do what they
// need to achieve their primary function.
//
// The name of each instance can be used as a key into the Type map supplied to the adapter
// at configuration time via the method 'SetQuotaTypes'.
// These Type associated with an instance describes the shape of the instance
type Handler interface {
	adapter.Handler

	// HandleQuota is called by Mixer at request time to deliver instances to
	// to an adapter.
	HandleQuota(context.Context, *Instance, adapter.QuotaArgs) (adapter.QuotaResult, error)
}
