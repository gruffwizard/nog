package docker

import (
	"bufio"
	"fmt"
	"github.com/docker/docker/api/types"
	"io"
	"os"
	"strings"
// "github.com/msoap/byline"
)

func checkQuarkusReady(s string) {

		s=strings.ToLower(s)
	//	if //strings.Contains(s,"quarkus main thread") &&
		//	 strings.Contains(s,"listening on:") {
				 println(">>>>"+s)
	//}

}

func (nog *NogDockerClient) JoinContainer(ID string) error {

	var inout chan []byte

	if body, err := nog.cli.ContainerAttach(nog.ctx, ID, types.ContainerAttachOptions{
		Stream: true,
		Stdout: true,
		Stderr: true,
		Stdin:  true,
	}); err != nil {
		panic(err)
	} else {

		// stdin handler and closer

		inout = make(chan []byte)
		go func() {
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				inout <- []byte(scanner.Text())
			}
		}()

		go func(w io.WriteCloser) {
			for {
				data, ok := <-inout
				if !ok {
					fmt.Println("!ok")
					w.Close()
					return
				}

				_, _ = w.Write(append(data, '\n'))
			}
		}(body.Conn)

		// everything else ..


		go io.Copy(os.Stdout, body.Reader)
		go io.Copy(os.Stderr, body.Reader)
	}

	return nil
}
