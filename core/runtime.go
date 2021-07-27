package util

import (
	"path"
	"runtime"
)

func GetRuntimeDir() (string, error) {

	_, filename, _, ok := runtime.Caller(1)
	var cwdPath string
	if ok {
		cwdPath = path.Join(path.Dir(filename), "") // the the main function File directory
	} else {
		cwdPath = "./"
	}
	return cwdPath, nil
}
