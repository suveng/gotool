package runtime

import (
	"runtime"
	"testing"
)

func TestGetRuntimeInfo(t *testing.T) {
	// 获取 go root
	println(runtime.GOROOT())

	// 获取cpu核数
	cpu := runtime.NumCPU()
	println(cpu)

	// 获取当前系统平台
	goos := runtime.GOOS
	println(goos)
	goarch := runtime.GOARCH
	println(goarch)

	// Golang 默认所有任务都运行在一个 cpu 核里，
	//如果要在 goroutine 中使用多核，可以使用 runtime.GOMAXPROCS 函数修改，当参数小于 1 时使用默认值。
	gomaxprocs := runtime.GOMAXPROCS(1)
	println(gomaxprocs)

}
