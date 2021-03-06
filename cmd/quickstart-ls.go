/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/gruffwizard/nog/cli"
	"github.com/spf13/cobra"
)

func newQuickStartLS(cli *cli.CLI) *cobra.Command {

	qslsCmd := &cobra.Command{
		Use:   "ls",
		Short: "list quickstart samples",
		Long:  `list quickstart samples`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cli.ListQuickStarts()
		},
	}

	return qslsCmd

}
