// Copyright 2019 Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package visualize

import (
	"fmt"
	"github.com/gardener/test-infra/pkg/util"
	"github.com/gardener/test-infra/pkg/util/dot"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

var (
	path   string
	output string

	status bool
	spec   bool
)

// AddCommand adds run-testrun to a command.
func AddCommand(cmd *cobra.Command) {
	cmd.AddCommand(vizCmd)
}

var vizCmd = &cobra.Command{
	Use:   "visualize",
	Short: "Run the testrunner with a helm template containing testruns",
	Aliases: []string{
		"viz",
	},
	Run: func(cmd *cobra.Command, args []string) {
		var file []byte
		var err error

		if path != "" {
			file, err = ioutil.ReadFile(path)
			if err != nil {
				log.Fatalf("unable to read from pipe: %s", err.Error())
			}
		} else {
			// try to read from stdin
			file, err = ioutil.ReadAll(os.Stdin)
			if err != nil {
				log.Fatalf("unable to read from pipe: %s", err.Error())
			}
		}

		tr, err := util.ParseTestrun(file)
		if err != nil {
			log.Fatalf("unable to parse testrun: %s", err.Error())
		}

		var graph string
		if status {
			graph, err = dot.GenerateDotFileFromStatus(&tr)
			if err != nil {
				log.Fatalf("unable to generate graph: %s", err.Error())
			}
		} else if spec {
			graph, err = dot.GenerateDotFileFromSpec(&tr)
			if err != nil {
				log.Fatalf("unable to generate graph: %s", err.Error())
			}
		}

		if output != "" {
			if err := ioutil.WriteFile(output, []byte(graph), os.ModePerm); err != nil {
				log.Fatalf("unable to write fiel to %s: %s", output, err.Error())
			}
			log.Infof("Successfully written graph to %s", output)
		}
		fmt.Print(graph)
	},
}

func init() {
	// configuration flags
	vizCmd.Flags().StringVar(&path, "path", "", "Path to the testrun")
	vizCmd.Flags().StringVarP(&output, "output", "o", "", "Path to the testrun")
	vizCmd.Flags().BoolVar(&spec, "spec", true, "Generate graph from spec.testflow")
	vizCmd.Flags().BoolVar(&status, "status", false, "Generate graph from status")
}
