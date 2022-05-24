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

package result

import "preflight/checker"

type CheckResult struct {
	Checker      checker.Interface
	Passed       bool   `json:"passed"`
	ErrorMessage string `json:"error_message"`
}

type RunnerResult struct {
	// result meet the required
	Passed []CheckResult
	// setup checker meet error
	Failed []CheckResult
	// if user ignore the checker result will downgrade the level to Warnings result.
	Warnings []CheckResult
}

type Descriptor struct {
	CheckerName  string            `json:"checker_name"`
	CheckerType  string            `json:"checker_type"`
	IsPassed     bool              `json:"is_passed,omitempty"`
	ErrorMessage string            `json:"error_message,omitempty"`
	Metadata     *checker.Metadata `json:"metadata,omitempty"`
}

// Response used to format for show check report.
type Response struct {
	Passed   []Descriptor `json:"passed,omitempty"`
	Failed   []Descriptor `json:"failed,omitempty"`
	Warnings []Descriptor `json:"warnings,omitempty"`
}
