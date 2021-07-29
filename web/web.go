package web

import (
	"fmt"
	"net/http"
	"strconv"
)

// 开启监听
func Start(port int, handler *Container) {
	if err := http.ListenAndServe(":"+strconv.Itoa(port), handler); err != nil {
		panic(err)
	}
}

// 构建container
func New() *Container {
	handler := &Container{Router: map[string]http.HandlerFunc{}}
	return handler
}

// web容器结构
type Container struct {
	Router map[string]http.HandlerFunc
}

// 增加路由
func (con *Container) AddRoute(method string, url string, handlerFunc http.HandlerFunc) bool {
	var key string
	key = buildKey(method, url)
	//container.Router
	if _, ok := con.Router[key]; ok {
		panic(fmt.Sprintf("route=%s already exists", key))
	} else {
		con.Router[key] = handlerFunc
	}
	return true
}

// url映射关系的key
func buildKey(method string, url string) string {
	return method + "-" + url
}

// 实现 http.Handler
func (con *Container) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	method := request.Method
	url := request.URL.Path

	key := buildKey(method, url)

	if handlerFunc, ok := con.Router[key]; ok {
		handlerFunc(writer, request)
	} else {
		fmt.Printf("404 not fund: %s\n", key)
	}
}
