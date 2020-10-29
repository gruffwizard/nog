package cmd

import (
    "github.com/gruffwizard/nog/helpers"
    "github.com/gruffwizard/nog/cli"
    "errors"
)
func validate(config *cli.CLI) error {


  if config.SrcDir!="" && config.SrcVol!="" {
      return errors.New("srcdir and srcvol options are mutually exclusive")
  }

  if config.MvnDir!="" && config.MvnVol!="" {
    return errors.New("mvndir and mvnvol options are mutually exclusive")}


  if config.MvnDir!="" {
    if err:=helpers.IsSafeDirectory("mvn cache dir",config.MvnDir); err!=nil {
      return err
    }
  }

  if config.SrcDir!="" {
    if err:=helpers.IsSafeDirectory("mvn source dir",config.SrcDir); err!=nil {
      return err
    }
  }

  if config.MvnDir=="" && config.MvnVol=="" {
    m,err:=helpers.MavenLocation()
    if err!=nil { return err }
    config.MvnDir=m
  }

  SrcDir,err:=helpers.CurrentDir();
  if  err!=nil {	return err }
  if  err=helpers.IsSafeDirectory("current directory",SrcDir); err!=nil { return err }

  config.SrcDir=SrcDir

  return nil
}
