package launcher


import (
	"strings"
	"os"
  "fmt"
  "path/filepath"
	"os/user"
)

type launcher struct {

  mode   string
	imageName string
	maven *mountPoint
	source *mountPoint

}

type mountPoint struct {
	Type string
	Location string
}

func NewDevLauncher() *launcher {

	l:=new(launcher)
	l.mode="dev"
	l.imageName="gruffwizard/nog-quarkus:latest"

	return l

}

func (l *launcher) SetMaven(loc string) {

	l.maven=toMountPoint(loc)

}


func (l *launcher) SetSource(loc string) {

		l.source=toMountPoint(loc)
}


/*
	Run sets up  docker
*/
func (l *launcher) Run() {

	// build required local things
	// source code location...
	buildLocation(l.source)
	buildLocation(l.maven)


}

func (l *launcher) Display() {

  fmt.Println("Nog config")
  fmt.Printf("mode   : %s\n",l.mode)
	fmt.Printf("image  : %s\n",l.imageName)
	fmt.Printf("maven  : %v\n",l.maven)
	fmt.Printf("source : %v\n",l.source)


}

func toMountPoint(location string) (*mountPoint) {

	mp := new(mountPoint)

	if  strings.HasPrefix(strings.ToLower(location),"vol:") {
			mp.Location=location[4:]
			mp.Type="vol"
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
	mp.Type="file"

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

func buildLocation(mp *mountPoint) {

	if mp.Type=="file" {
			exists,err := IsLocalDir(mp.Location)
			if err!=nil {
				panic(err)
			}
			if !exists {
				os.MkdirAll(mp.Location,os.ModePerm)
			}
	}

	if mp.Type=="vol" {
			err := mkDockerVol(mp.Location)
			if err!=nil {
						panic(err)
			}
	}
}
