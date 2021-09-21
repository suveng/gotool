package main

import "fmt"

func init() {
	fmt.Println("world")
	//我们还可以做其他更高阶的事情，比如 platform.RegisterPlugin({"func": Hello}) 之类的，向插件平台自动注册该插件的函数
}

func Say() {
	fmt.Println("plugin say")
}

func Fly() {
	fmt.Println("plugin fly")
}
