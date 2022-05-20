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

package app

import (
	"fmt"
	"github.com/spf13/cobra"
	"preflight/checker"
	"preflight/formatter"
	"strconv"
)

var checkCmd = &cobra.Command{
	Use:     "port",
	Short:   "check port",
	Long:    "",
	Example: `check port 22`,
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		p, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}

		checker := checker.PortCheck{Port: p}
		passed, err := checker.Validate()

		if err == nil {
			return nil
		}

		m := checker.Metadata()
		data, err := formatter.FormatDescriptor(formatter.Descriptor{
			IsPassed:     passed,
			ErrorMessage: err.Error(),
			CheckerName:  checker.Name(),
			Metadata:     &m,
		})

		if err != nil {
			return err
		}

		fmt.Println(string(data))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
