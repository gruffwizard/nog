package launcher
import (
  "fmt"
)

func ListVolumes() {


  c,err := NewDockerClient()

  if err!=nil { panic(err)}

  l,err := c.ListVolumes()

  for _,v := range l {

    fmt.Printf("%v",v.Name)
  }

}



func ListImages() {


  c,err := NewDockerClient()

  if err!=nil { panic(err)}

  l,err := c.Images()

  for _,v := range l {
        for _,n := range v.Tags {
        fmt.Printf("%v\n",n)
      }
  }

}
