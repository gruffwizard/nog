package docker

import (
    "github.com/docker/docker/api/types"
    "io"
    "os"
)

func (nog *NogDockerClient) PullImage(image string) (error) {

  reader, err := nog.cli.ImagePull(nog.ctx, image, types.ImagePullOptions{})
	if err != nil {
		return err
	}

	io.Copy(os.Stdout, reader)

  return nil

}
