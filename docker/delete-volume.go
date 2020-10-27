
package docker

import (

   nogerrs "github.com/gruffwizard/nog/errors"
   "errors"
)


func (nog *NogDockerClient) DeleteVolume(name string) (error) {


    vol,err := nog.GetVolume(name)
    if err!=nil { return err}
    if vol==nil { return errors.New("nog volume "+name+" does not exist")}

    if !vol.CreatedByNog()  { return nogerrs.NewNotNog("volume "+name+" not created by nog. Delete manually") }

    return  nog.cli.VolumeRemove(nog.ctx,name,false)


}
