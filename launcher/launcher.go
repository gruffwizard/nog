package launcher


import (
	"strings"
	"os"
  "fmt"
  "path/filepath"
	"os/user"
	"os/signal"
	"syscall"

)

type launcher struct {

  mode   string
	imageName string
	mounts map[string]*mountPoint
	cli *nogDockerClient
	id string
}

type mountPoint struct {
	Type string
	Location string
}


func NewDevLauncher() *launcher {

	l:=new(launcher)
	l.mounts=make(map[string]*mountPoint)
	l.mode="dev"
	l.imageName="gruffwizard/nog-quarkus-theia:latest"

	client, err := NewDockerClient()
	if (err!=nil) { panic(err)}
	l.cli=client

	l.setupExit()


	return l

}

func (l *launcher) setupExit() {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
	  <- sigs
	  l.cleanUp()
	  os.Exit(0)
	}()
}

func  (l *launcher)  cleanUp() {

	l.cli.StopContainer(l.id)
	l.cli.EndContainer(l.id)

}

func  (l *launcher)  LaunchContainer() {


	 local,err := l.cli.LocalImage(l.imageName)
	 if err != nil { panic(err) }

	 if !local {

		 err := l.cli.PullImage(l.imageName)
		 if err != nil { panic(err) }
	}

	id,err := l.cli.CreateContainer(l.imageName,l.mounts)
	if err != nil { panic(err) }

	l.id=id

	err = l.cli.JoinContainer(id)
	if err != nil { panic(err) }


	err = l.cli.StartContainer(id)
	if err != nil { panic(err) }

	err = l.cli.WaitForContainer(id)
	if err != nil { panic(err) }

}

func (l *launcher) SetMaven(loc string) {

	l.mounts["maven"]=toMountPoint(loc)

}


func (l *launcher) SetSource(loc string) {

		l.mounts["source"]=toMountPoint(loc)
}


/*
	Run sets up  docker
*/
func (l *launcher) Run() {

	// build required local things
	// source code location...
	l.buildLocation(l.mounts["source"])
	l.buildLocation(l.mounts["maven"])

	// now run the nog image

	l.LaunchContainer()


}

func (l *launcher) Display() {

  fmt.Println("Nog config")
  fmt.Printf("mode   : %s\n",l.mode)
	fmt.Printf("image  : %s\n",l.imageName)
	fmt.Printf("maven  : %v\n",l.mounts["maven"])
	fmt.Printf("source : %v\n",l.mounts["source"])


}


func toMountPoint(location string) (*mountPoint) {

	mp := new(mountPoint)

	if  strings.HasPrefix(strings.ToLower(location),"vol:") {
			mp.Location=location[4:]
			mp.Type="volume"
			return mp
	}

	if  strings.HasPrefix(strings.ToLower(location),"file:") {
			location=location[5:]
	}
	if strings.HasPrefix(location,"~") {
			location=location[1:]
			usr,_:= user.Current()
			location=filepath.Join(usr.HomeDir,location)
	}

	mp.Location,_ = filepath.Abs(location)
	mp.Type="bind"

	return mp

}

func IsLocalDir(file string) (bool,error) {

  abs,_ := filepath.Abs(file)
  info, err := os.Stat(abs)
  if os.IsNotExist(err) {
      return false,nil
  }

  if err!=nil {
    panic(err)
  }

  return info.IsDir(),nil

}

func (l *launcher) buildLocation(mp *mountPoint) {

	if mp.Type=="bind" {
			exists,err := IsLocalDir(mp.Location)
			if err!=nil {
				panic(err)
			}
			if !exists {
				os.MkdirAll(mp.Location,os.ModePerm)
			}
	}

	if mp.Type=="volume" {
			err := l.cli.CreateVolume(mp.Location)
			if err!=nil {
						panic(err)
			}
	}
}
