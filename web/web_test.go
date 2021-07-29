package web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestStart(t *testing.T) {
	container := New()
	container.AddRoute("post", "/post", post)
	Start(9999, container)
}

func post(writer http.ResponseWriter, r *http.Request) {
	fmt.Println("post")
}
