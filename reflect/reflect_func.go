package reflect

import (
	"fmt"
	"reflect"
)

// 测试 var 定义类型
var (
	ApiTypeC1 = ApiType{code: "C1", impl: Api.DoSomething}
	ApiTypeC2 = ApiType{code: "C2", impl: implB{}}
)

type ApiType struct {
	code string
	impl interface{}
}

func getVarType() {
	impl := ApiTypeC1.impl
	fmt.Println(reflect.TypeOf(impl))
	impl = ApiTypeC2.impl
	fmt.Println(reflect.TypeOf(impl))

}
