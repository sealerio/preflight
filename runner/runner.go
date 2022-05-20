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
	"github.com/pkg/errors"
	"preflight/checker"
)

type ValidateResult struct {
	Checker      checker.Interface
	Passed       bool   `json:"passed"`
	ErrorMessage string `json:"error_message"`
}

type Results struct {
	Passed []ValidateResult
	Failed []ValidateResult
	Errors []ValidateResult
}

type Runner interface {
	// Execute all given Checkers.
	Execute() Results
}

type CheckRunner struct {
	Checks       []checker.Interface
	CheckResults Results
}

// Execute checker validate and dispatch to different results
func (c CheckRunner) Execute() Results {
	for _, check := range c.Checks {
		// run the validation
		passed, err := check.Validate()
		if err != nil {
			c.CheckResults.Errors = append(c.CheckResults.Errors, ValidateResult{
				Checker:      check,
				Passed:       false,
				ErrorMessage: err.Error(),
			})
			continue
		}

		if !passed {
			c.CheckResults.Failed = append(c.CheckResults.Failed, ValidateResult{
				Checker:      check,
				Passed:       false,
				ErrorMessage: err.Error(),
			})
			continue
		}

		c.CheckResults.Passed = append(c.CheckResults.Passed, ValidateResult{
			Checker: check,
			Passed:  true,
		})
	}

	return c.CheckResults
}

// NewCheckRunner select Checklist via build-in check map.
// if len(Checklist)==0 ,return not specified error.
func NewCheckRunner(checkList []checker.Interface) (Runner, error) {
	if len(checkList) == 0 {
		return nil, errors.New("Checklist could not be nil")
	}
	return CheckRunner{
		Checks: checkList,
	}, nil
}
