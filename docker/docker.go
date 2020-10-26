package docker

import (

	"context"
	"github.com/docker/docker/client"

  	"os"

  "github.com/docker/docker/api/types"



  	"github.com/docker/docker/pkg/stdcopy"


		"fmt"

)


type mountPoint struct {
	Type string
	Location string
}


type mounts struct {
	mounts map[string]*mountPoint

}


func NewMounts() *mounts {

  c:= new(mounts)
  c.mounts=make(map[string]*mountPoint)
  return c

}
func (l *mounts)  List()  {

	for k,v := range l.mounts {

		fmt.Printf(" %s mount : %s -> %s\n",v.Type,k,v.Location)

	}

}


func (l *mounts)  AddVolMount(from string, loc string) {

	mp := new(mountPoint)
	mp.Location=loc
	mp.Type="volume"

	l.mounts[from]=mp

}

func (l *mounts)  AddDirMount(from string, loc string) {

	mp := new(mountPoint)
	mp.Location=loc
	mp.Type="bind"

	l.mounts[from]=mp

}
var CommonMounts map[string]string = map[string]string{
    "maven": "/home/nog/.m2",
    "source": "/home/nog/src",
}

type NogDockerClient struct {

  ctx context.Context
  cli *client.Client
	Verbose bool

}

type NogImage struct {
	  Name string
		Labels map[string]string
		Tags []string
}

type NogVolume struct {
  Name string
  Labels map[string]string
}

func (v *NogVolume) CreatedByNog() (bool) {

    if v.Labels==nil { return false}
    if _, exist := v.Labels["nog"]; exist  { return true}

    return false
}

func NewDockerClient() (*NogDockerClient,error) {

  cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil,err
	}

  nog := new(NogDockerClient)
  nog.cli=cli
  nog.ctx=context.Background()

  return nog,nil

}



// list Nog volumes





func  (nog *NogDockerClient) ShowLogs1(ID string) (error) {

  	out, err := nog.cli.ContainerLogs(nog.ctx, ID, types.ContainerLogsOptions{ShowStdout: true,ShowStderr:true})
  	if err != nil {
  		return err
  	}

  	stdcopy.StdCopy(os.Stdout, os.Stderr, out)

    return nil
}
