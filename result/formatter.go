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

type Formatter interface {
	Format(opts ...Option) Response
}

type defaultFormatter struct {
	result RunnerResult
}

func (d defaultFormatter) Format(opts ...Option) Response {
	options := defaultFormatOptions
	for _, opt := range opts {
		opt(&options)
	}
	return d.ParesToResponse(options.Ignores)
}

func (d defaultFormatter) ParesToResponse(ignores []string) Response {
	var passedList, failedList, warnings []Descriptor

	if len(d.result.Passed) > 0 {
		for _, result := range d.result.Passed {
			desc := Descriptor{
				CheckerType: result.Checker.Type(),
				CheckerName: result.Checker.PrettyName(),
			}
			passedList = append(passedList, desc)
		}
	}

	if len(d.result.Failed) > 0 {
		for _, result := range d.result.Failed {
			t := result.Checker.Type()
			m := result.Checker.Metadata()
			name := result.Checker.PrettyName()
			desc := Descriptor{
				CheckerType:  t,
				CheckerName:  name,
				IsPassed:     result.Passed,
				Metadata:     &m,
				ErrorMessage: result.ErrorMessage,
			}
			if NotIn(t, ignores) {
				failedList = append(failedList, desc)
			} else {
				warnings = append(warnings, desc)
			}
		}
	}

	return Response{
		Passed:   passedList,
		Failed:   failedList,
		Warnings: warnings,
	}
}

func NewDefaultFormatter(r RunnerResult) Formatter {
	return defaultFormatter{
		result: r,
	}
}
