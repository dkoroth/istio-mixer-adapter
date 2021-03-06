// Copyright 2018 Google LLC
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

package context

import (
	"istio.io/istio/mixer/pkg/adapter"
	"net/url"
)

// A Context contains all the information needed to communicate with Apigee
// home servers.
type Context interface {
	Log() adapter.Logger
	Organization() string
	Environment() string
	Key() string
	Secret() string

	ApigeeBase() url.URL
	CustomerBase() url.URL
}
