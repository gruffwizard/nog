package cli

import (
  "gopkg.in/yaml.v2"
  "io/ioutil"
  "net/http"
  "errors"
)
type QuickStartSideContainer struct {

  Image string
  Ports map[int]int

}
type QuickStartConfig struct {

  GuideURL         string `yaml:"guide"`
  Description      string `yaml:"desc"`
  SideContainers []QuickStartSideContainer `yaml:"sides"`

}

type QuickStartList struct {
    Version string
    QuickStarts map[string]QuickStartConfig `yaml:"qs"`
}

func CheckValidQuickStart(name string) (error) {

  qs,err:=FetchSupportedQuickStarts()
  if err != nil { return err }

  if _, ok := qs.QuickStarts[name]; !ok {
    return  errors.New("specified quickstart "+name+" is not a recognised Nog friendly one")
  }
  return nil

}
func LoadQuickStarts(file string) (*QuickStartList,error) {


  y, err := ioutil.ReadFile(file)
  if err != nil { return nil,err }

  q := new(QuickStartList)
  err = yaml.Unmarshal(y, q)

  if err != nil { return nil,err }

  return q,nil

}

func SaveQuickStarts(list *QuickStartList) error {

  _,err := yaml.Marshal(list)
  return err
}

func FetchSupportedQuickStarts() (*QuickStartList,error) {

  resp, err := http.Get(QuickStartURL)
	if err != nil { return nil,err }

	defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil { return nil,err }


  q := new(QuickStartList)
  err = yaml.Unmarshal(body, q)

  if err != nil { return nil,err }

  return q,nil


}
