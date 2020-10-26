package docker

import (

  "github.com/docker/docker/api/types/mount"
 //"github.com/docker/go-connections/nat"
 "fmt"
 "github.com/docker/docker/api/types/container"
)

func  (nog *NogDockerClient) CreateContainer(image string,cmd []string,m *mounts,envs []string) (string,error) {

	labels := make(map[string]string)
  labels["nog"]="true"

	// build mounts
	mountConfig := []mount.Mount{}

	for k,v := range m.mounts {

		bind := mount.TypeBind
		if v.Type=="volume" { bind=mount.TypeVolume}

		mountConfig=append(mountConfig,mount.Mount{Type:bind,Source:v.Location,Target:k})
	}


  portsConfig,err :=nog.ImagePortDetails(image)
  if err!=nil { return "",err}

  if nog.Verbose {
    for p,b := range portsConfig {
      fmt.Printf(" port %v : %v\n",p,b[0])
    }
  }

	hostConfig := container.HostConfig{
		Mounts: mountConfig,
		PortBindings: portsConfig,
	}


  	resp, err := nog.cli.ContainerCreate(nog.ctx, &container.Config{
  		Image: image,
  		Cmd:   cmd,
			AttachStdout: true,
	    AttachStderr: true,
			AttachStdin: true,
			OpenStdin:   true,
			Tty:   true,
			Labels: labels,
	    Env: envs,
  	}, &hostConfig, nil, nil, "")

    if err!=nil { return "",err }

  	return resp.ID,nil
}
