// Copyright © 2022 Alibaba Group Holding Ltd.
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

package runner

import (
	"preflight/checker"
)

func BuildInitCheckers() []checker.Interface {
	return []checker.Interface{
		checker.PortCheck{Port: 6443},
		checker.NumCPUCheck{NumCPU: 2},
		checker.MemCheck{Mem: 1700},
		checker.OsCheck{
			OSType:         "linux",
			OSDistribution: []string{"ubuntu", "centos"},
			// Requires 3.10+, or newer
			KernelVersions: []string{`^3\.[1-9][0-9].*$`, `^([4-9]|[1-9][0-9]+)\.([0-9]+)\.([0-9]+).*$`}},
	}
}
