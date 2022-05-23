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
	"encoding/json"
	"fmt"
	"preflight/results"
	"preflight/runner"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type RunArgs struct {
	Skip        []string
	Ignore      []string
	CheckerType string
	CheckerArgs string
}

var runArgs *RunArgs

var runCmd = &cobra.Command{
	Use:     "run",
	Short:   "preflight run",
	Long:    "",
	Example: `preflight run`,
	RunE:    runPreflight,
}

func runPreflight(cmd *cobra.Command, args []string) error {
	r, err := runner.NewRunner(runner.WithSkips(runArgs.Skip))
	if err != nil {
		return errors.Wrap(err, "failed to init runner")
	}

	resp := results.NewDefaultFormatter(r.Execute()).Format(results.WithIgnores(runArgs.Ignore))

	responseJSON, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		return errors.Wrap(err, "format json failed")
	}

	fmt.Println(string(responseJSON))

	return nil
}

func init() {
	runArgs = &RunArgs{}
	runCmd.Flags().StringVarP(&runArgs.CheckerType, "checker", "c", "", "specify checker type")
	runCmd.Flags().StringVar(&runArgs.CheckerArgs, "args", "", "specify checker args when you want run specify checker")
	runCmd.Flags().StringSliceVar(&runArgs.Skip, "skip", []string{}, "run all checkers expect this checker")
	runCmd.Flags().StringSliceVar(&runArgs.Ignore, "ignore-errors", []string{}, "run all checkers ignore this checker error")
	rootCmd.AddCommand(runCmd)
}
