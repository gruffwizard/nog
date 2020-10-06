/*
Copyright © 2020 Steve Poole  gruff.wizard@yahoo.com

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

// devCmd represents the dev command
var devCmd = &cobra.Command{
	Use:   "dev [git-repo]",
	Short: "quarkus in dev mode",
	Long: `
	Runs the quarkus experience in conjunction with your IDE
	Launches quarkus in dev mode.
	Source and maven repos will be by default be local. To use
	volume managed source or maven repos overide with the -s and -m options
	`,
	Args: func(cmd *cobra.Command, args []string) error {
	    if len(args) > 1 {
	      return errors.New("too many arguments")
	    }

			return nil

	  },
	Run: func(cmd *cobra.Command, args []string) {

		launcher := launcher.NewDevLauncher()
		if mvnLoc=="" { mvnLoc="file:~/.m2" }
		if srcLoc=="" { srcLoc="."}
		
		launcher.SetMaven(mvnLoc)
		launcher.SetSource(srcLoc)

		if Verbose {
			launcher.Display()
		}

		launcher.Run()


	},
}

func init() {
	rootCmd.AddCommand(devCmd)
}
