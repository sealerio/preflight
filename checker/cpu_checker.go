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
	"fmt"
	"runtime"
	"strings"

	"github.com/pkg/errors"
)

// NumCPUCheck checks if current number of CPUs is not less than required
type NumCPUCheck struct {
	NumCPU int
}

func (c NumCPUCheck) Type() string {
	return "CPU"
}

func (c NumCPUCheck) PrettyName() string {
	return fmt.Sprintf("%s:%d", strings.ToLower(c.Type()), c.NumCPU)
}

func (NumCPUCheck) Metadata() Metadata {
	return Metadata{
		Description: "Check number of CPUs required",
		Level:       FatalLevel,
		Explain:     "more cpu number means more power,less cpu number means that the program may not run normally, or be very slow.",
		Suggestion:  "Maybe you should upgrade your machine",
	}
}

func (c NumCPUCheck) Validate() (bool, error) {
	numCPU := runtime.NumCPU()
	if numCPU < c.NumCPU {
		return false, errors.Errorf(
			"the number of available CPUs %d is less than the required %d", numCPU, c.NumCPU)
	}
	return true, nil
}
