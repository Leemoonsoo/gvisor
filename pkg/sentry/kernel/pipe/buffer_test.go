// Copyright 2019 The gVisor Authors.
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

package pipe

import (
	"testing"
	"unsafe"

	"gvisor.dev/gvisor/pkg/sentry/usermem"
)

func TestBufferSize(t *testing.T) {
	bufferSize := unsafe.Sizeof(buffer{})
	if bufferSize < usermem.PageSize {
		t.Errorf("buffer is less than a page")
	}
	if bufferSize > (2 * usermem.PageSize) {
		t.Errorf("buffer is greater than two pages")
	}
}