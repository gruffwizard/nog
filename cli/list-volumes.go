package cli

import (
	"fmt"
	"github.com/gruffwizard/nog/docker"
)

func (cli *CLI) ListVolumes() error {

	c, err := docker.NewDockerClient()
	if err != nil { return err }

	l, err := c.ListVolumes()
	if err != nil { return err }

	for _, v := range l {

		fmt.Printf("%v\n", v.Name)
	}

	return nil

}
