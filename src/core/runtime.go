package core

import (
	"path"
	"runtime"
)

type Runtime interface {
	// 获取运行时的目录路径
	GetRuntimeDir() string
}

type RuntimeUtil struct {
}

func (r *RuntimeUtil) GetRuntimeDir() (string, error) {

	_, filename, _, ok := runtime.Caller(1)
	var cwdPath string
	if ok {
		cwdPath = path.Join(path.Dir(filename), "") // the the main function File directory
	} else {
		cwdPath = "./"
	}
	return cwdPath, nil
}
