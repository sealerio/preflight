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
	"preflight/pkg/system"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

// OsCheck machine system information
type OsCheck struct {
	// OSType: freebsd, linux
	OSType string
	// valid os distribution list : ubuntu, centos
	OSDistribution []string
	// KernelVersions define supported kernel version. It is a group of regexps.
	KernelVersions []string
}

func (a OsCheck) Type() string {
	return strings.ToLower("OS")
}

func (a OsCheck) PrettyName() string {
	return a.Type()
}

func (a OsCheck) Validate() (bool, error) {
	info, err := system.GetHostInfo()
	if err != nil {
		return false, errors.Wrapf(err, "failed to get system info")
	}

	if a.OSType != info.OS {
		return false, errors.Errorf("required os type is %s,but got %s", a.OSType, info.OS)
	}

	if !a.validateDistribution(info.OSDistribution) {
		return false, errors.Errorf("required os distribution list is %s,but got %s",
			a.OSDistribution, info.OSDistribution)
	}

	return a.validateKernel(info.Kernel)
}

func (a OsCheck) validateDistribution(distribution string) bool {
	for _, s := range a.OSDistribution {
		if distribution == s {
			return true
		}
	}
	return false
}

func (a OsCheck) validateKernel(version string) (bool, error) {
	for _, versionRegexp := range a.KernelVersions {
		r := regexp.MustCompile(versionRegexp)
		if r.MatchString(version) {
			return true, nil
		}
	}

	return false, errors.Errorf("required os Kernel 3.10+, or newer,but got %s", version)
}

func (OsCheck) Metadata() Metadata {
	return Metadata{
		Description: "Check host operating system info",
		Level:       PanicLevel,
		Explain:     "",
		Suggestion:  "",
	}
}
