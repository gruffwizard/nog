package cli

import (
	"fmt"
	"github.com/gruffwizard/nog/docker"
)

func (cli *CLI) ListImages() error {

	c, err := docker.NewDockerClient()

	if err != nil {
		return err
	}

	l, err := c.Images()

	for _, v := range l {
		for _, n := range v.Tags {
			fmt.Printf("%v\n", n)
		}
	}

	return nil
}
