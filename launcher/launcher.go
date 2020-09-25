package launcher


import (
	"os"
  "fmt"
  "path/filepath"
)

type launcher struct {

  Mode   string
  Src string
  SrcType string
}

func (l *launcher) Run() {

  fmt.Println("launching")
  fmt.Printf("mode: %s\n",l.Mode)
  fmt.Printf("source %s at %s\n",l.SrcType,l.Src)


}



func NewEditModeLauncher(srcLoc string) (*launcher,error) {

        l := new(launcher)

        l.Mode="edit"

  			var exists bool
  			var err error
  			// where's the source?
  			if srcLoc!="" {
  				// the assumption is that this is a volume name
  				// if it is then we are done. If not it might be
  				exists,err= CheckForValidDockerVolume(srcLoc)

  				if err!=nil {return nil,err}

  				if exists {
  					l.Src=srcLoc
            l.SrcType="vol"
  				} else {
  					// is it a local dir?
  					exists,err =IsLocalDir(srcLoc)
  					if err!=nil {return nil,err}

  					l.Src=srcLoc
            l.SrcType="dir"

  				}
  				// a local file dir so check there too

  			}

        return l,nil

}
func ValidateEditMode() {


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
