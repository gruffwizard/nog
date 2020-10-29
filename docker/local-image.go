package docker

func (nog *NogDockerClient) LocalImage(tag string) (bool, error) {

	images, err := nog.Images()
	if err != nil {
		return false, err
	}
	for _, i := range images {

		for _, t := range i.Tags {
			if t == tag {
				return true, nil
			}
		}
	}

	return false, nil

}
