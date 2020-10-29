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
	"fmt"
	"github.com/spf13/cobra"
	"github.com/gruffwizard/nog/cli"
)


func newVersion(l *cli.CLI) *cobra.Command {

versionCmd := &cobra.Command{
	Use:   "version",
	Short: "Display Nog version",
	Long: `Displays version info. Combine with -v to see extended data`,
	Example: "nog version",
	RunE: func(cmd *cobra.Command, args []string) (error) {

		fmt.Println(l.NogVersion)


		if cli.Verbose {
				fmt.Printf("commit  : %s\n",l.NogCommit)
				fmt.Printf("built   : %s\n",l.NogBuiltDate)
				fmt.Printf("builtby : %s\n",l.NogBuiltBy)
		}

		return nil
	},
}
return versionCmd

}
