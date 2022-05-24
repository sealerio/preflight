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
	"os"
	"preflight/checker"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

const (
	checkerType = "CHECKER TYPE"
	checkerName = "CHECKER NAME"
	level       = "LEVEL"
	description = "DESCRIPTION"
)

type ListResponse struct {
	Type     string
	Name     string
	Metadata checker.Metadata
}

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "preflight list",
	Long:    "",
	Example: `preflight list`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var resp []ListResponse

		for types, checker := range checker.GetAllCheckers() {
			resp = append(resp, ListResponse{
				Name:     fmt.Sprintf("%s:%s", types, "${arg}"),
				Type:     types,
				Metadata: checker.Metadata(),
			})
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{checkerName, checkerType, level, description})

		for _, r := range resp {
			table.Append([]string{r.Name, r.Type, r.Metadata.Level, r.Metadata.Description})
		}

		table.Render()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
