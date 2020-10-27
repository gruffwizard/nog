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

	_ = c.StopContainer(cli.ActiveID)
	_ = c.EndContainer(cli.ActiveID)

  }
  return nil

}
