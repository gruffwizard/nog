package docker



func  (nog *NogDockerClient) StopContainer(ID string) (error) {

  return nog.cli.ContainerStop(nog.ctx,ID, nil)

}
