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
	"github.com/gruffwizard/nog/cli"
	"github.com/spf13/cobra"
)

func newDev(l *cli.CLI) *cobra.Command {

	dev := &cobra.Command{
		Use:   "dev",
		Short: "quarkus in dev mode",
		Long: `
	Runs the quarkus experience in conjunction with your IDE

	Source is assumed to be in current directory and there should be a pom.xml file present

	Your local maven cache is shared with Nog. This may have side effects if your
	local machine is not linux based and you use platform specific artifacts.

	To use a seperate maven repository specify the -m option to provide an alternative maven
	repository location.

	`,

		Args: func(cmd *cobra.Command, args []string) error {

			return validate(l)

		},

		RunE: func(cmd *cobra.Command, args []string) error {

			return l.Run(args)

		},
	}

	dev.Flags().StringVarP(&l.MvnVol, "mvnvol", "m", "", "maven cache volume")
	dev.Flags().StringVarP(&l.MvnDir, "mvndir", "d", "", "maven directory")
	dev.Flags().StringVarP(&l.SrcVol, "srcvol", "l", "", "source volume")
	dev.Flags().StringVarP(&l.SrcDir, "srcdir", "s", "", "source directory")
	dev.Flags().BoolVarP(&l.Convert, "convert", "c", false, "convert to quarkus")

	dev.Flags().BoolVarP(&l.IDEMode, "ide", "i", false, "use containerised ide")

	return dev
}
