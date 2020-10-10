package launcher

import (

	"context"
	"github.com/docker/docker/client"
  "github.com/docker/docker/api/types/filters"
  	"os"
    "github.com/docker/docker/api/types"
    "github.com/docker/docker/api/types/volume"
  	"github.com/docker/docker/api/types/container"
		"github.com/docker/docker/api/types/mount"
	 "github.com/docker/go-connections/nat"
"io"
  	"github.com/docker/docker/pkg/stdcopy"
		"fmt"
    "errors"
)

var CommonMounts map[string]string = map[string]string{
    "maven": "/home/nog/.m2",
    "source": "/home/nog/src",
}

type nogDockerClient struct {

  ctx context.Context
  cli *client.Client

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

func NewDockerClient() (*nogDockerClient,error) {

  cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil,err
	}

  nog := new(nogDockerClient)
  nog.cli=cli
  nog.ctx=context.Background()

  return nog,nil

}


func (nog *nogDockerClient) CreateVolume(name string) (error) {

  labels := make(map[string]string)
  labels["nog"]="true"

    _,err := nog.cli.VolumeCreate(nog.ctx, volume.VolumeCreateBody{Name:name,Labels:labels})

    return err
}


func (nog *nogDockerClient) GetVolume(name string) (*NogVolume, error) {

  v,err := nog.cli.VolumeInspect(nog.ctx, name)

  if err!=nil { return nil,err}

  return &NogVolume{Name:v.Name, Labels:v.Labels},nil

}

func (nog *nogDockerClient) DeleteVolume(name string) (error) {

  // volume exits?

    vol,err := nog.GetVolume(name)
    if err!=nil { return err}
    if vol==nil { return errors.New("nog volume "+name+" does not exist")}

    if vol.CreatedByNog() == false { return NewNotNog("volume "+name+" not created by nog. Delete manually") }

    return  nog.cli.VolumeRemove(nog.ctx,name,false)


}
// list Nog volumes

func (nog *nogDockerClient) ListVolumes() ([]NogVolume,error) {

  f :=filters.NewArgs()
  f.Add("label","nog=true")

  list,err := nog.cli.VolumeList(nog.ctx, f)
  if err!=nil { return nil,err }

  noglist := []NogVolume{}
  for _, s := range list.Volumes {

      noglist=append(noglist, NogVolume{Name:s.Name, Labels:s.Labels})
  }


  return noglist,nil

}

func (nog *nogDockerClient) PullImage(image string) (error) {

  reader, err := nog.cli.ImagePull(nog.ctx, image, types.ImagePullOptions{})
	if err != nil {
		return err
	}

	io.Copy(os.Stdout, reader)

  return nil

}

func (nog *nogDockerClient) LocalImage(tag string) (bool,error) {

	images,err := nog.Images()
	if err!=nil { return false,err}
	for _,i := range images {

			for _,t := range i.Tags {
				if t==tag {return true,nil}
			}
	}

	return false,nil

}

func (nog *nogDockerClient) Images() ([]NogImage,error) {

	f :=filters.NewArgs()
  f.Add("label","nog=true")

	imageinfo, err := nog.cli.ImageList(nog.ctx, types.ImageListOptions{Filters:f})
  if err != nil {
		return nil,err
	}

	noglist := []NogImage{}

	for _,s := range imageinfo {

		noglist=append(noglist, NogImage{Name:s.ID, Labels:s.Labels, Tags:s.RepoTags})

	}

  return noglist,nil

}

func  (nog *nogDockerClient) CreateContainer(image string,mounts map[string]*mountPoint) (string,error) {

	labels := make(map[string]string)
  labels["nog"]="true"

	// build mounts
	mountConfig := []mount.Mount{}

	for k,v := range mounts {

		target:=CommonMounts[k]
		if target == ""  { target = k}

		bind := mount.TypeBind
		if v.Type=="volume" { bind=mount.TypeVolume}

		fmt.Printf("mount %v as %v at %v over %v\n",k,v.Type,v.Location,target)

		mountConfig=append(mountConfig,mount.Mount{Type:bind,Source:v.Location,Target:target})
	}

	// build ports
	data := []string{ "8080:8080","8081:8081","9001:9001"}

	_,portsConfig,_:= nat.ParsePortSpecs(data)
	fmt.Printf("\nNat:%v",portsConfig)

	hostConfig := container.HostConfig{
		Mounts: mountConfig,
		PortBindings: portsConfig,
	}


  	resp, err := nog.cli.ContainerCreate(nog.ctx, &container.Config{
  		Image: image,
  		Cmd:   []string{"/home/nog/tools/runner.sh"},
			AttachStdout: true,
      AttachStderr: true,
			StdinOnce:   true,
  		Tty:   false,
			Labels: labels,
  	}, &hostConfig, nil, nil, "")

    if err!=nil { return "",err }

  	return resp.ID,nil
}


func  (nog *nogDockerClient) StartContainer(ID string) (error) {
	return nog.cli.ContainerStart(nog.ctx, ID, types.ContainerStartOptions{});
}


func  (nog *nogDockerClient) JoinContainer(ID string) (error) {

	done := make(chan struct{})
	if body, err := nog.cli.ContainerAttach(nog.ctx,ID, types.ContainerAttachOptions{
		Stream: true,
		Stdout: true,
	}); err != nil {
		panic(err)
	} else {
		go func() {
			defer body.Conn.Close()
			if _, err := stdcopy.StdCopy(os.Stdout, os.Stderr, body.Reader); err != nil {
				panic(err)
			}
			close(done)
		}()
	}

	return nil
}


func  (nog *nogDockerClient) WaitForContainer(ID string) (error) {

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





func  (nog *nogDockerClient) StopContainer(ID string) (error) {

  return nog.cli.ContainerStop(nog.ctx,ID, nil)
	
}

func  (nog *nogDockerClient) EndContainer(ID string) (error) {


  return nog.cli.ContainerRemove(nog.ctx,ID, types.ContainerRemoveOptions{})
}


func  (nog *nogDockerClient) ShowLogs(ID string) (error) {

  	out, err := nog.cli.ContainerLogs(nog.ctx, ID, types.ContainerLogsOptions{ShowStdout: true,ShowStderr:true})
  	if err != nil {
  		return err
  	}

  	stdcopy.StdCopy(os.Stdout, os.Stderr, out)

    return nil
}
