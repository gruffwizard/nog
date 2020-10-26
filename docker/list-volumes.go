package docker


import (
  "github.com/docker/docker/api/types/filters"
    

)

func (nog *NogDockerClient) ListVolumes() ([]NogVolume,error) {

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
