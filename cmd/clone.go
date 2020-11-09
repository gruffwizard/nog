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
	"github.com/gruffwizard/nog/cli"
	"github.com/spf13/cobra"
	"path/filepath"
)

func newClone(l *cli.CLI) *cobra.Command {

	qs := &cobra.Command{
		Use:     "clone",
		Short:   "Clone git repo and launch quarkus",
		Long: "Run the Quarkus experience using a git repo",

		Args: func(cmd *cobra.Command, args []string) error {

			if len(args) < 1 {
				return errors.New("must specify git repo path")
			}


			if l.IDEMode && l.SrcDir=="" && l.SrcVol=="" { l.SrcVol="nog-"+filepath.Base(args[0])}

			return validate(l)
		},

		RunE: func(cmd *cobra.Command, args []string) error {

			l.Clone = args[0]
			return l.Run(args[1:])

		},
	}

	qs.Flags().BoolVarP(&l.IDEMode, "ide", "i", false, "use containerised ide")
	qs.Flags().StringVarP(&l.MvnVol, "mvnvol", "m", "", "maven cache volume")
	qs.Flags().StringVarP(&l.MvnDir, "mvndir", "d", "", "maven directory")
	qs.Flags().StringVarP(&l.SrcVol, "srcvol", "l", "", "source volume")
	qs.Flags().StringVarP(&l.SrcDir, "srcdir", "s", "", "source directory")


	return qs
}
