
package docker

import (



    "github.com/docker/go-connections/nat"

)

func (nog *NogDockerClient) ImagePortDetails(name string) ( nat.PortMap ,error) {

  i,_,err := nog.cli.ImageInspectWithRaw(nog.ctx,name)
  if err!=nil { return nil,err}

  m:= nat.PortMap{}
  for p := range i.ContainerConfig.ExposedPorts {
    //[]nat.PortBinding
      m[p]=[]nat.PortBinding{nat.PortBinding{HostPort:string(p)}}
  }
  return m,nil
  /*
  var results []string

  for k,_ := range i.ContainerConfig.ExposedPorts {
        bits:=strings.Split(string(k),"/")
        results=append(results,bits[0]+":"+bits[0])
  }
  /*
  _,portsConfig,_:= nat.ParsePortSpecs(data)
  if nog.Verbose  {
    fmt.Printf("Ports:%v\n",portsConfig)
  }
  */
//  return results,nil

}
