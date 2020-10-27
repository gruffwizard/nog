
package cmd

import (
  "github.com/spf13/cobra"
  "github.com/gruffwizard/nog/cli"
)

func NewCMD(nog *cli.CLI) *cobra.Command  {

    root:=NewRoot(nog)

    root.AddCommand(NewDev(nog))

    vol :=NewVol(nog)
    vol.AddCommand(NewVolLS(nog))
    root.AddCommand(vol)

    image:=NewImage(nog)
    image.AddCommand(NewImageLS(nog))
    root.AddCommand(image)

    qs:=NewQuickStart(nog)
    qs.AddCommand(NewQuickStartLS(nog))
    qs.AddCommand(NewQuickStartCP(nog))

    root.AddCommand(qs)

    root.AddCommand(NewVersion(nog))

    return root
}
