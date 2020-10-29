package launcher

/*
func (l *launcher) Run() {


	 local,err := l.cli.LocalImage(l.Image)
	 if err != nil { panic(err) }

	 if !local {

		 err := l.cli.PullImage(l.Image)
		 if err != nil { panic(err) }
	}

	id,err := l.cli.CreateContainer(l.Image,l.Cmd,l.mounts)
	if err != nil { panic(err) }

	l.id=id

	err = l.cli.JoinContainer(id)
	if err != nil { panic(err) }


	err = l.cli.StartContainer(id)
	if err != nil { panic(err) }

	err = l.cli.WaitForContainer(id)
	if err != nil { panic(err) }

}

*/

/*

func (l *launcher) buildLocation1(mp *mountPoint) {

	if mp.Type=="bind" {
			exists,err := IsLocalDir(mp.Location)
			if err!=nil {
				panic(err)
			}
			if !exists {
				os.MkdirAll(mp.Location,os.ModePerm)
			}
	}

	if mp.Type=="volume" {
			err := l.cli.CreateVolume(mp.Location)
			if err!=nil {
						panic(err)
			}
	}
}
*/
