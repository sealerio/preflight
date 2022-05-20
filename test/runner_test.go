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

package test

import (
	"fmt"
	"preflight/checker"
	"preflight/formatter"
	"preflight/runner"
	"testing"
)

func TestNewCheckRunner(t *testing.T) {
	list := []checker.Interface{
		checker.PortCheck{Port: 34442424},
		checker.NumCPUCheck{NumCPU: 2},
		checker.FileExistingCheck{Path: "/code/preflight/cmd/main.go"},
	}
	checkers, err := runner.NewCheckRunner(list)
	if err != nil {
		t.Errorf("failed to init runner err: %s", err)
	}
	data, err := formatter.FormatResults(checkers.Execute())

	if err != nil {
		t.Errorf("formatter err: %s", err)
	}

	fmt.Println(string(data))
}
