// Copyright 2016 The LUCI Authors.
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

package backend

// Meta is backend metadata about a single configuration file.
type Meta struct {
	// ConfigSet is the item's config set.
	ConfigSet string
	// Path is the item's path within its config set.
	Path string

	// ContentHash is the content hash.
	ContentHash string
	// Revision is the revision string.
	Revision string
	// ViewURL is the URL surfaced for viewing the config.
	ViewURL string
}