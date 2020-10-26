package cli


import (
    "github.com/gruffwizard/nog/docker"
    "fmt"
)


func  (cli *CLI)  CleanUp() error {

  if cli.ActiveID != "" {

	fmt.Println("Nog closing down")

  c,err := docker.NewDockerClient()
  if err!=nil { return err}

	c.StopContainer(cli.ActiveID)
	c.EndContainer(cli.ActiveID)

  }
  return nil

}
