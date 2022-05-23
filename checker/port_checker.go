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

	"github.com/pkg/errors"

	"net"
)

type PortCheck struct {
	Port int
}

func (p PortCheck) Type() string {
	return "Port"
}

func (p PortCheck) PrettyName() string {
	return fmt.Sprintf("%s:%d", strings.ToLower(p.Type()), p.Port)
}

func (PortCheck) Metadata() Metadata {
	return Metadata{
		Description: "Check the the port is available",
		Level:       FatalLevel,
		Explain:     "an open port is a network port that accepts incoming packets from remote locations",
		Suggestion:  "Maybe you should check your machine of the port is available for use",
	}
}

// Validate if the port is available.
func (p PortCheck) Validate() (bool, error) {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", p.Port))
	if err != nil {
		return false, errors.Errorf("Port %d is in use,%v", p.Port, err)
	}
	if ln != nil {
		if err = ln.Close(); err != nil {
			return false, errors.Errorf("when closing port %d, encountered %v", p.Port, err)
		}
	}

	return true, nil
}
