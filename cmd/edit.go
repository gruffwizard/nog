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
	"errors"
	"github.com/spf13/cobra"
	"github.com/gruffwizard/nog/launcher"
)



// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit [ git-repo ]",
	Short: "quarkus in full container mode",
	Long: `
Runs the quarkus experience fully containers
Lauches a browser based editor based on Eclipse Thea with quarkus enabled.
Source and maven repos will be by default managed in volumes. To use
local source or maven repos overide with the -s and -m options
`,
Args: func(cmd *cobra.Command, args []string) error {
    if len(args) > 1 {
      return errors.New("too many arguments")
    }
		return nil

  },
	Run: func(cmd *cobra.Command, args []string) {


			launcher.RunContainer()

	},
}

func init() {
	rootCmd.AddCommand(editCmd)

}
