package docker

import (
	"context"
	"github.com/docker/docker/client"

	"os"

	"github.com/docker/docker/api/types"

	"github.com/docker/docker/pkg/stdcopy"
	"net"
	"fmt"

)

type mountPoint struct {
	Type     string
	Location string
}

type mounts struct {
	mounts map[string]*mountPoint
}

func (l *mounts) List() {

	for k, v := range l.mounts {

		fmt.Printf(" %s mount : %s -> %s\n", v.Type, k, v.Location)

	}

}

func (c *ContainerDef) AddVolMount(from string, loc string) {

	if c.mounts==nil { c.mounts=make(map[string]*mountPoint) }
	mp := new(mountPoint)
	mp.Location = loc
	mp.Type = "volume"

	c.mounts[from] = mp

}

func (c *ContainerDef) AddDirMount(from string, loc string) {

	if c.mounts==nil { c.mounts=make(map[string]*mountPoint) }
	mp := new(mountPoint)
	mp.Location = loc
	mp.Type = "bind"

	c.mounts[from] = mp

}

var CommonMounts map[string]string = map[string]string{
	"maven":  "/home/nog/.m2",
	"source": "/home/nog/src",
}

type NogDockerClient struct {
	ctx     context.Context
	cli     *client.Client
	Verbose bool
}

type NogImage struct {
	Name   string
	Labels map[string]string
	Tags   []string
}

type NogVolume struct {
	Name   string
	Labels map[string]string
}

func (v *NogVolume) CreatedByNog() bool {

	if v.Labels == nil {
		return false
	}
	if _, exist := v.Labels["nog"]; exist {
		return true
	}

	return false
}

func NewDockerClient() (*NogDockerClient, error) {

	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}

	nog := new(NogDockerClient)
	nog.cli = cli
	nog.ctx = context.Background()

	return nog, nil

}

// list Nog volumes

func (nog *NogDockerClient) ShowLogs1(ID string) error {

	out, err := nog.cli.ContainerLogs(nog.ctx, ID, types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true})
	if err != nil {
		return err
	}

	_, _ = stdcopy.StdCopy(os.Stdout, os.Stderr, out)

	return nil
}

func isPortFree(port string) bool {

	ln, err := net.Listen("tcp", ":" + port)

  if err != nil {
    return false
  }

  ln.Close()

	return true
}
