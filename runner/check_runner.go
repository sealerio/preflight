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
	"preflight/results"

	"github.com/pkg/errors"
)

type CheckRunner struct {
	Checks  []checker.Interface
	Results results.RunnerResult
}

// Execute checker validate and dispatch to different results
func (c *CheckRunner) Execute() results.RunnerResult {
	for _, check := range c.Checks {
		// run the validation
		passed, err := check.Validate()
		if err != nil {
			c.Results.Failed = append(c.Results.Failed, results.CheckResult{
				Checker:      check,
				Passed:       false,
				ErrorMessage: err.Error(),
			})
			continue
		}

		c.Results.Passed = append(c.Results.Passed, results.CheckResult{
			Checker: check,
			Passed:  passed,
		})
	}

	return c.Results
}

// NewCheckRunnerByList select Checklist via build-in check map.
// if len(Checklist)==0 ,return not specified error.
func NewCheckRunnerByList(checkList []checker.Interface) (Runner, error) {
	if len(checkList) == 0 {
		return nil, errors.New("Checklist could not be nil")
	}
	return &CheckRunner{
		Checks: checkList,
	}, nil
}
