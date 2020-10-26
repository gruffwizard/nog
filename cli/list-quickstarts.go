package cli

import (


  "fmt"

)

func (l *CLI)  ListQuickStarts() error {

  // get quick start list ...
  list,err := FetchSupportedQuickStarts()
  if err!=nil { return err}
  fmt.Println("Nog tested quickstarts")
  for n,qs := range list.QuickStarts {

    fmt.Printf("%s : %s\n",n,qs.Description)
    fmt.Printf("%s\n\n",qs.GuideURL)
  }

  return nil
}
