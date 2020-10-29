package cmd

import (
	"github.com/gruffwizard/nog/cli"
	"github.com/spf13/cobra"
)

func NewCMD(nog *cli.CLI) *cobra.Command {

	root := newRoot(nog)

	root.AddCommand(newDev(nog))

	vol := newVol(nog)
	vol.AddCommand(newVolLS(nog))
	root.AddCommand(vol)

	image := newImage(nog)
	image.AddCommand(newImageLS(nog))
	root.AddCommand(image)

	qs := newQuickStart(nog)
	qs.AddCommand(newQuickStartLS(nog))
	qs.AddCommand(newQuickStartCP(nog))

	root.AddCommand(qs)

	root.AddCommand(newVersion(nog))

	return root
}
