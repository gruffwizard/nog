
package cli

import (

    "github.com/gruffwizard/nog/docker"

)

func (l *CLI) CreateVolume(name string) error {


  c,err := docker.NewDockerClient()

  if err!=nil { return err }


  return  c.CreateVolume(name)

}
