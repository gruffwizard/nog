package cli

import (
    "github.com/gruffwizard/nog/docker"
    "fmt"
)

/*
The heart of Nog. Creates a container with the relevent mounts, envs etc
and connects to it to run the specified command

At this point there are only two images available a quarkus one which
contains latest Java and Native code tools dependencies. The other is the same
but with the Theia IDE added


*/
func (l *CLI) Run(args []string) error {

  var c Config

  envs:= []string{}
  if Verbose { envs=append(envs,"NOG_VERBOSE=1")}


  image:=QuarkusImage
  mode:="dev"

  if l.IDEMode {
    mode="ide"
    image=TheiaImage
    envs=append(envs,"NOG_START_IDE=1")
  }

  fmt.Printf("Running Nog in %s mode. Enjoy!\n",mode)

  mounts := docker.NewMounts()

  if l.QuickStart!="" { envs=append(envs,"NOG_QUICKSTART="+l.QuickStart) }
  if l.QuickStartOnly { envs=append(envs,"NOG_QUICKSTART_ONLY=1") }


  cmd:=[]string{"/home/nog/tools/nog.sh"}

  if len(args)>0 { cmd=args}

  if l.MvnVol!="" { mounts.AddVolMount(NogMavenHome,l.MvnVol) }
  if l.MvnDir!="" { mounts.AddDirMount(NogMavenHome,l.MvnDir) }

  if l.SrcDir!="" { mounts.AddDirMount("/home/nog/src",l.SrcDir) }
  if l.SrcVol!="" { mounts.AddVolMount("/home/nog/src",l.SrcVol) }


  if Verbose {
    c.Display()
    mounts.List()
    fmt.Printf("envs :%v\n",envs)
  }

  d,err := docker.NewDockerClient()
  d.Verbose=Verbose

  if err!=nil { return err }

	 local,err := d.LocalImage(image)
	 if err != nil { return err }

	 if !local {

		 err := d.PullImage(image)
		 if err != nil { return err }
	}


	id,err := d.CreateContainer(image,cmd,mounts,envs)
	if err != nil { return err }

	l.ActiveID=id

	err = d.JoinContainer(id)
	if err != nil { return err }

	err = d.StartContainer(id)
	if err != nil { return err }


	return  d.WaitForContainer(id)

}
