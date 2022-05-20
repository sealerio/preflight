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

// Interface validates the state of the system or network.
type Interface interface {
	// Validate the asset value for the checker.
	// if validate not pass or encounter error,will return check error.
	Validate() (bool, error)
	// Name of checker
	Name() string
	// Metadata return the checker's Metadata
	Metadata() Metadata
}

//Metadata contains useful information regarding the check
type Metadata struct {
	// short description for checker.
	Description string `json:"description,omitempty"`
	// indicate the check level,it is useful to show if Validate failed.
	// like info,warn,fatal,panic.
	Level string `json:"level,omitempty"`
	// long message for show checker,used to get help message，generally it shows why this check failed.
	Explain string `json:"explain,omitempty"`
	// show Suggestion if Validate failed.
	Suggestion string `json:"suggestion,omitempty"`
}
