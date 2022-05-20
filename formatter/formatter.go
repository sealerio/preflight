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

package formatter

import (
	"encoding/json"
	"github.com/pkg/errors"
	"preflight/runner"
)

func FormatResults(r runner.Results) ([]byte, error) {
	response := ParesResponse(r)
	responseJSON, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		return nil, errors.Wrap(err, "format json failed")
	}
	return responseJSON, nil
}

func FormatDescriptor(d Descriptor) ([]byte, error) {
	responseJSON, err := json.MarshalIndent(d, "", "    ")
	if err != nil {
		return nil, errors.Wrap(err, "format json failed")
	}
	return responseJSON, nil
}

func ParesResponse(r runner.Results) Response {
	var passedList, failedList, errorList []Descriptor

	if len(r.Passed) > 0 {
		for _, result := range r.Passed {
			d := Descriptor{
				CheckerName: result.Checker.Name(),
			}
			passedList = append(passedList, d)
		}
	}

	if len(r.Failed) > 0 {
		for _, result := range r.Failed {
			m := result.Checker.Metadata()
			d := Descriptor{
				CheckerName:  result.Checker.Name(),
				IsPassed:     result.Passed,
				Metadata:     &m,
				ErrorMessage: result.ErrorMessage,
			}
			failedList = append(failedList, d)
		}
	}

	if len(r.Errors) > 0 {
		for _, result := range r.Errors {
			m := result.Checker.Metadata()
			d := Descriptor{
				CheckerName:  result.Checker.Name(),
				IsPassed:     result.Passed,
				Metadata:     &m,
				ErrorMessage: result.ErrorMessage,
			}
			errorList = append(errorList, d)
		}
	}

	return Response{
		Passed: passedList,
		Failed: failedList,
		Errors: errorList,
	}
}
