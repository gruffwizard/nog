package docker

import (

  	"github.com/docker/docker/api/types/container"
)


func  (nog *NogDockerClient) WaitForContainer(ID string) (error) {

  	statusCh, errCh := nog.cli.ContainerWait(nog.ctx, ID, container.WaitConditionNotRunning)
  	select {
  	case err := <-errCh:
  		if err != nil {
  			return err
  		}
  	case <-statusCh:
  	}

    return nil
}
