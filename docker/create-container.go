package docker

import (
	"github.com/docker/docker/api/types/mount"
	"fmt"
	"strings"
	"errors"
	"github.com/docker/docker/api/types/container"
)


type ContainerDef struct {
	args []string
	envs []string
	Image string
	Cmd []string
	mounts map[string]*mountPoint
}

func (c *ContainerDef) AddEnv(k string,v string) {
				c.envs = append(c.envs, k+"="+v)
}

func (c *ContainerDef) Display() {
		fmt.Printf("Image : %s\n",c.Image)
		fmt.Printf("Cmd   : %v\n",c.Cmd)
		fmt.Printf("Args  : %v\n",c.args)
		fmt.Printf("Envs  : %v\n",c.envs)
		for k, v := range c.mounts {
			fmt.Printf(" %s mount : %s -> %s\n", v.Type, k, v.Location)
		}

}
//image string, cmd []string, m *mounts, envs []string
func (nog *NogDockerClient) CreateContainer(c ContainerDef) (string, error) {

	labels := make(map[string]string)
	labels["nog"] = "true"

	// build mounts
	mountConfig := []mount.Mount{}

	for k, v := range c.mounts {

		bind := mount.TypeBind
		if v.Type == "volume" {
			bind = mount.TypeVolume
		}

		mountConfig = append(mountConfig, mount.Mount{Type: bind, Source: v.Location, Target: k})
	}

	portsConfig, err := nog.ImagePortDetails(c.Image)
	if err != nil {
		return "", err
	}


		for p, b := range portsConfig {
			pno := strings.Split(string(p),"/")
			if ! isPortFree(pno[0]) {
				return "",errors.New("port "+pno[0]+" is already in use")
			}
			if nog.Verbose { fmt.Printf(" port %v : %v\n", p, b[0]) }
		}

	hostConfig := container.HostConfig{
		Mounts:       mountConfig,
		PortBindings: portsConfig,
	}

	resp, err := nog.cli.ContainerCreate(nog.ctx, &container.Config{
		Image:        c.Image,
		Cmd:          c.Cmd,
		AttachStdout: true,
		AttachStderr: true,
		AttachStdin:  true,
		OpenStdin:    true,
		Tty:          true,
		Labels:       labels,
		Env:          c.envs,
	}, &hostConfig, nil, "")

	if err != nil {
		return "", err
	}

	return resp.ID, nil
}
