package helpers

import (
  "path/filepath"
  "os"
  "os/user"
  "errors"
)

func IsSafeDirectory(desc string,dir string) error {

  if dir==""  { return errors.New(desc+" is empty") }

  path,err:=filepath.Abs(dir)

  info, err := os.Stat(path)
  if os.IsNotExist(err) { return err }

  if !info.IsDir() {  return errors.New(desc+" "+dir+" is not a directory") }

  home,err:=HomeDir()
  if home==path { return errors.New(desc+ " "+dir+" is user home directory and cannot be used as a mount point")}

  return nil

}
func FileExists(path string,name string) bool {

    filename := filepath.Join(path,name)

    info, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false
    }

    return !info.IsDir()
}

func HomeDir() (string,error) {

  user, err := user.Current()
  if err != nil {return "",err }

  return user.HomeDir,nil

}


func CurrentDir() (string,error) {

  current,err := os.Getwd()

  if err!=nil {return "",err }

  return current,nil

}

func MavenLocation() (string,error) {


  current, err := HomeDir()
  if err != nil {return "",err }

  mvnrepo:=filepath.Join(current,".m2")

  info, err := os.Stat(mvnrepo)
  if os.IsNotExist(err) {
      return "",err
  }

  if !info.IsDir() { return "",errors.New("maven location ~/.m2 is not a directory")}

  return mvnrepo,nil

}
