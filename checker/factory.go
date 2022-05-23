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

package checker

import (
	"github.com/pkg/errors"
)

// Register all checks
var memNumCheck Interface = &MemCheck{}
var cpuNumCheck Interface = &NumCPUCheck{}
var fileExistingCheck Interface = &FileExistingCheck{}
var portInuseCheck Interface = &PortCheck{}

var nameToChecksMap = map[string]Interface{
	memNumCheck.Type():       memNumCheck,
	cpuNumCheck.Type():       cpuNumCheck,
	fileExistingCheck.Type(): fileExistingCheck,
	portInuseCheck.Type():    portInuseCheck,
}

func GetAllCheckers() map[string]Interface {
	return nameToChecksMap
}

func GetAllCheckerTypes() []string {
	all := make([]string, len(nameToChecksMap))
	for k := range nameToChecksMap {
		all = append(all, k)
	}
	return all
}

func GetCheckersByType(checkType string) (Interface, error) {
	if check, exists := nameToChecksMap[checkType]; exists {
		return check, nil
	}
	return nil, errors.Errorf("checker %s not found", checkType)
}
