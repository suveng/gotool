package main

import (
	//"path"
	"plugin"
)

func main() {
	p, err := plugin.Open("./testplugin_v1.so")
	if err != nil {
		panic(err)
	}
	f, err := p.Lookup("Say")
	if err != nil {
		panic(err)
	}
	f.(func())()
}
