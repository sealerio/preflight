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

import "strings"

type RunOptions struct {
	// Skips checker Validate by checker type,default is lowercase.
	Skips []string
	// IsTolerable return immediately when an error is reported.
	NotTolerable bool
}

// Option configures a runner list
type Option func(*RunOptions)

var defaultRunOptions = RunOptions{
	Skips:        []string{},
	NotTolerable: false,
}

func WithSkips(skips []string) Option {
	return func(o *RunOptions) {
		var s []string
		if len(skips) > 0 {
			for _, skip := range skips {
				s = append(s, strings.ToLower(skip))
			}
		}
		o.Skips = s
	}
}

func WithToleration(it bool) Option {
	return func(o *RunOptions) {
		o.NotTolerable = it
	}
}
