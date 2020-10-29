
package cmd

import (
  "github.com/spf13/cobra"
  "github.com/gruffwizard/nog/cli"
)

func NewCMD(nog *cli.CLI) *cobra.Command  {

    root:=newRoot(nog)

    root.AddCommand(newDev(nog))

    vol :=newVol(nog)
    vol.AddCommand(newVolLS(nog))
    root.AddCommand(vol)

    image:=newImage(nog)
    image.AddCommand(newImageLS(nog))
    root.AddCommand(image)

    qs:=newQuickStart(nog)
    qs.AddCommand(newQuickStartLS(nog))
    qs.AddCommand(newQuickStartCP(nog))

    root.AddCommand(qs)

    root.AddCommand(newVersion(nog))

    return root
}
