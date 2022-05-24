//go:build linux
// +build linux

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
	"fmt"
	"strings"
	"syscall"

	"github.com/pkg/errors"
)

type MemCheck struct {
	Mem uint64
}

func (m MemCheck) Type() string {
	return strings.ToLower("Memory")
}

func (m MemCheck) PrettyName() string {
	return fmt.Sprintf("%s:%d", m.Type(), m.Mem)
}

func (MemCheck) Metadata() Metadata {
	return Metadata{
		Description: "Check the number of megabytes of memory required",
		Level:       FatalLevel,
		Explain:     "more memory number means more power,less memory number means that the program may not run normally, or be very slow.",
		Suggestion:  "Maybe you should upgrade your machine",
	}
}

func (m MemCheck) Validate() (bool, error) {
	info := syscall.Sysinfo_t{}
	err := syscall.Sysinfo(&info)
	if err != nil {
		return false, errors.Wrapf(err, "failed to get system info")
	}

	// Total holds the total usable memory. Unit holds the size of a memory unit in bytes. Multiply them and convert to MB
	actual := info.Totalram * uint64(info.Unit) / 1024 / 1024
	if actual < m.Mem {
		return false, errors.Errorf("the system RAM (%d MB) is less than the minimum %d MB", actual, m.Mem)
	}
	return true, nil
}
