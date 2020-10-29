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

func newRoot(l *cli.CLI) *cobra.Command {

	root := &cobra.Command{
		Use:   "nog",
		Short: "Containerised Quarkus development environment",
		Long: `
Nog is an interface between local code and a containerised quarkus
development environment. Nog mounts local maven repos, source code etc
into the container running quarkus and allows you to develop in your
IDE and have Quarkus (in a container ) build and run it.

As an added bonus Nog also includes a browser based IDE to allow
for the full quarkus experience without having to install anything.

Well except for Nog - and you've already done that.
	`,
	}

	root.PersistentFlags().BoolVarP(&cli.Verbose, "verbose", "v", false, "verbose")

	return root
}
