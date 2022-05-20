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
	"github.com/pkg/errors"
	"os"
)

type FileExistingCheck struct {
	Path string
}

func (f FileExistingCheck) Name() string {
	return "FileExistingCheck"
}

func (FileExistingCheck) Metadata() Metadata {
	return Metadata{
		Description: "Check the given file does is already exist",
		Level:       WarnLevel,
		Explain:     "file existing check means that the file exist on the specific path, if not, maybe your program not run normally",
		Suggestion:  "Maybe you need to check the file path or find why it is not exist",
	}
}

// Validate if the given file already exists.
func (f FileExistingCheck) Validate() (bool, error) {
	if _, err := os.Stat(f.Path); err != nil {
		return false, errors.Errorf("%s doesn't exist", f.Path)
	}
	return true, nil
}
