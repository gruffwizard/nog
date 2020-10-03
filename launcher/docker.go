package launcher

import (
  "github.com/fsouza/go-dockerclient"
  "fmt"
)

func CheckForValidDockerVolume(vol string) (bool,error) {

  return true,nil

}


func RunContainer() {

    client, err := docker.NewClientFromEnv()
    fmt.Println("resp:%v\n",client)
    fmt.Println("err:%v\n",err)
  }
