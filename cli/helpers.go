package cli

import (
	"os"
	"path/filepath"
)

func IsLocalDir(file string) (bool, error) {

	abs, _ := filepath.Abs(file)
	info, err := os.Stat(abs)
	if os.IsNotExist(err) {
		return false, nil
	}

	if err != nil {
		panic(err)
	}

	return info.IsDir(), nil

}
