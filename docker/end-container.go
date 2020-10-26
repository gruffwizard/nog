package docker

import (

    "github.com/docker/docker/api/types"
)



func  (nog *NogDockerClient) EndContainer(ID string) (error) {
  return nog.cli.ContainerRemove(nog.ctx,ID, types.ContainerRemoveOptions{})
}
