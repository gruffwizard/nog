package docker



import (
    "github.com/docker/docker/api/types"
    "github.com/docker/docker/api/types/filters"
)



func (nog *NogDockerClient) Images() ([]NogImage,error) {

	f :=filters.NewArgs()
  f.Add("label","nog=true")

	imageinfo, err := nog.cli.ImageList(nog.ctx, types.ImageListOptions{Filters:f})
  if err != nil {return nil,err}

	noglist := []NogImage{}

	for _,s := range imageinfo {

		noglist=append(noglist, NogImage{Name:s.ID, Labels:s.Labels, Tags:s.RepoTags})

	}

  return noglist,nil

}
