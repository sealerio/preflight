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
	"encoding/json"
	"fmt"
	"preflight/checker"
	"preflight/result"
	"preflight/runner"
	"strings"
	"testing"
)

func TestNewCheckRunnerList(t *testing.T) {
	type ListResponse struct {
		Type     string
		Name     string
		Metadata checker.Metadata
	}

	var resp []ListResponse

	for types, c := range checker.GetAllCheckers() {
		resp = append(resp, ListResponse{
			Name:     fmt.Sprintf("%s:%s", strings.ToLower(types), "${arg}"),
			Type:     types,
			Metadata: c.Metadata(),
		})
	}

	responseJSON, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		t.Errorf("format json failed:%v\n", err)
	}

	fmt.Println(string(responseJSON))
}
func TestNewCheckRunner(t *testing.T) {
	list := []checker.Interface{
		checker.PortCheck{Port: 34442424},
		checker.NumCPUCheck{NumCPU: 2},
		checker.FileExistingCheck{Path: "/code/preflight/cmd/main.go"},
		checker.PortCheck{Port: 90},
	}
	r, err := runner.NewCheckRunner(list)
	if err != nil {
		t.Errorf("failed to init runner err: %s", err)
	}
	resp := result.NewDefaultFormatter(r.Execute()).Format(result.WithIgnores([]string{"Port", "FileExisting"}))

	responseJSON, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		t.Errorf("format json failed:%v\n", err)
	}

	fmt.Println(string(responseJSON))
}

func TestNewCheckRunnerWithIgnore(t *testing.T) {
	list := []checker.Interface{
		checker.PortCheck{Port: 34442424},
		checker.NumCPUCheck{NumCPU: 2},
		checker.FileExistingCheck{Path: "/code/preflight/cmd/main.go"},
		checker.PortCheck{Port: 90},
	}
	r, err := runner.NewCheckRunner(list)
	if err != nil {
		t.Errorf("failed to init runner err: %s", err)
	}
	resp := result.NewDefaultFormatter(r.Execute()).Format(result.WithIgnores([]string{"Port", "FileExisting"}))

	responseJSON, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		t.Errorf("format json failed:%v\n", err)
	}

	fmt.Println(string(responseJSON))
}
