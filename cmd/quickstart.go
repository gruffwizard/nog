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
	"github.com/spf13/cobra"
	"errors"
	"github.com/gruffwizard/nog/cli"

)




func newQuickStart(l *cli.CLI) *cobra.Command {

qs := &cobra.Command{
	Use:   "quickstart",
	Aliases: []string{"qs"},
	Short: "quickstart the Quarkus experience",
	Long:
`Run the Quarkus experience using a named sample.
For a complete list use 'nog quickstart ls'
`,

	Args: func(cmd *cobra.Command, args []string) error {


		if len(args)<1 { return errors.New("must specify quickstart sample name. (Use 'nog qs ls' to find available quickstarts)")}

		err := validate(l)

		if err!=nil { return err}

		return cli.CheckValidQuickStart(args[0])

	},

	RunE: func(cmd *cobra.Command, args []string) error {

		l.QuickStart=args[0]
		return l.Run(args[1:])

	},

}

qs.Flags().BoolVarP(&l.IDEMode, "ide", "i",false, "use containerised ide")
qs.Flags().StringVarP(&l.MvnVol, "mvnvol", "m","", "maven cache volume")
qs.Flags().StringVarP(&l.MvnDir, "mvndir", "d","", "maven directory")


return qs
}
