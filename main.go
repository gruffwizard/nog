/*
Copyright © 2020 Steve Poole  gruff.wizard@yahoo.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"github.com/gruffwizard/nog/cli"
	"github.com/gruffwizard/nog/cmd"
	"os"
	"os/signal"
	"syscall"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	builtBy = "unknown"
)

func main() {

	cli := new(cli.CLI)
	registerTearDown(cli)

	cli.NogVersion = version
	cli.NogCommit = commit
	cli.NogBuiltDate = date
	cli.NogBuiltBy = builtBy
	
	cmd := cmd.NewCMD(cli)

	cmd.Execute()
}

func registerTearDown(cli *cli.CLI) {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		cli.CleanUp()
		os.Exit(0)
	}()
}
