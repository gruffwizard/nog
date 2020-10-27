package cli


import (
    "github.com/gruffwizard/nog/docker"
    "fmt"
)


func  (cli *CLI)  CleanUp()  {

  if cli.ActiveID != "" {

	fmt.Println("Nog closing down")

  c,err := docker.NewDockerClient()
  if err!=nil { return }

	_ = c.StopContainer(cli.ActiveID)
	_ = c.EndContainer(cli.ActiveID)

  }


}
