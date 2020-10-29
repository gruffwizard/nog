package docker

func (nog *NogDockerClient) WaitForContainer(ID string) error {

	_, err := nog.cli.ContainerWait(nog.ctx, ID)
	return err

}
