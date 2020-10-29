package docker

import (
	"github.com/docker/docker/api/types"
)

func (nog *NogDockerClient) StartContainer(ID string) error {
	return nog.cli.ContainerStart(nog.ctx, ID, types.ContainerStartOptions{})
}
