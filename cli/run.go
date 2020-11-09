package cli

import (
	"fmt"
	"github.com/gruffwizard/nog/docker"
		"strings"
	//"github.com/pkg/browser"
)

func ensureImagePresent(d *docker.NogDockerClient,image string) error {

		local, err := d.LocalImage(image)
		if err != nil {
			return err
		}

		if !local {

			err := d.PullImage(image)
			if err != nil {
				return err
			}
		}

		return nil
}
func (l *CLI) buildDef(args []string) (docker.ContainerDef) {

	var c docker.ContainerDef

	if Verbose { c.AddEnv("NOG_VERBOSE","1") }

	c.Image = QuarkusImage

	if l.IDEMode {

		c.Image = TheiaImage
		c.AddEnv("NOG_START_IDE","1")
	}

	fmt.Println("Running Nog - Enjoy!")

	if l.QuickStart != "" { c.AddEnv("NOG_QUICKSTART",l.QuickStart) }
	if l.QuickStartOnly   { c.AddEnv("NOG_QUICKSTART_ONLY","1")     }

	if l.Clone != ""      { c.AddEnv("NOG_CLONE",l.Clone) }
	if l.Convert          { c.AddEnv("NOG_CONVERT","1") }

	c.Cmd = []string{"/home/nog/tools/nog.sh"}

	if len(args) > 0 {
		c.Cmd = args
	}

	if l.MvnVol != "" {
		c.AddVolMount(NogMavenHome, l.MvnVol)
	}
	if l.MvnDir != "" {
		c.AddDirMount(NogMavenHome, l.MvnDir)
	}

	if l.SrcDir != "" {
		c.AddDirMount("/home/nog/src", l.SrcDir)
	}
	if l.SrcVol != "" {
		c.AddVolMount("/home/nog/src", l.SrcVol)
	}

	return c
}

/*
The heart of Nog. Creates a container with the relevent mounts, envs etc
and connects to it to run the specified command

At this point there are only two images available a quarkus one which
contains latest Java and Native code tools dependencies. The other is the same
but with the Theia IDE added


*/
func (l *CLI) Run(args []string) error {

	c:=l.buildDef(args)

	if Verbose { c.Display() }


	d, err := docker.NewDockerClient()

	if err != nil {
		return err
	}

	d.Verbose = Verbose
	d.Scanner=func(s string) {

			s=strings.ToLower(s)
			if strings.Contains(s,"quarkus main thread") &&
			 	 strings.Contains(s,"listening on:") {
					 //browser.OpenURL("http://localhost:8080")
		}

	}
	d.Scanner=nil

	if err=ensureImagePresent(d,c.Image); err!=nil { return err}


	id, err := d.CreateContainer(c)
	if err != nil {
		return err
	}

	l.ActiveID = id


	err = d.JoinContainer(id)
	if err != nil {
		return err
	}

	err = d.StartContainer(id)
	if err != nil {
			println(err)
		return err
	}

	return d.WaitForContainer(id)

}
