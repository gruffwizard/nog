package docker

func (nog *NogDockerClient) GetVolume(name string) (*NogVolume, error) {

	v, err := nog.cli.VolumeInspect(nog.ctx, name)

	if err != nil {
		return nil, err
	}

	return &NogVolume{Name: v.Name, Labels: v.Labels}, nil

}
