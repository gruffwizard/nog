
package docker

import (
  "github.com/docker/docker/api/types/volume"
)


func (nog *NogDockerClient) CreateVolume(name string) (error) {

  labels := make(map[string]string)
  labels["nog"]="true"

    _,err := nog.cli.VolumeCreate(nog.ctx, volume.VolumeCreateBody{Name:name,Labels:labels})

    return err
}
