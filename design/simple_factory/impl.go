package simple_factory

import "fmt"

type ApiImplA struct {
}

type ApiImplB struct {
}

func (a *ApiImplB) Doing() {
	fmt.Println("impl b doing")
}

func (a *ApiImplA) Doing() {
	fmt.Println("impl a doing")
}
