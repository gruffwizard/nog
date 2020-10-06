package launcher

import (
  "github.com/fsouza/go-dockerclient"
  "fmt"
)

func CheckForValidDockerVolume(vol string) (bool,error) {

  return true,nil

}

func ListVolumes() {

    client, err := docker.NewClientFromEnv()
    if err !=nil  { panic(err) }

    var lvo docker.ListVolumesOptions
    lvo.Filters = make(map[string][]string)
    lvo.Filters["label"]=[]string{"nog"}

    vols,error :=client.ListVolumes(lvo)
    if error !=nil  { panic(error) }

    for _, v := range vols {
     fmt.Printf("%s \n",v.Name)
    }

}

func RunContainer() {

    client, err := docker.NewClientFromEnv()
    fmt.Println("resp:%v\n",client)
    fmt.Println("err:%v\n",err)
  }


  func mkDockerVol(l string) (*error ){

    client, err := docker.NewClientFromEnv()
    if err!=nil  {  panic(err) }

    var cvo docker.CreateVolumeOptions
    cvo.Name=l
    cvo.Labels=make(map[string]string)
    cvo.Labels["nog"]="true"

    vol,error := client.CreateVolume(cvo)
    fmt.Printf("\n%v,%v",vol,error)
    return nil
  }
