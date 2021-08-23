package simple_factory

var (
	ApiTypeC1 = ApiType{code: "C1", impl: ApiImplA{}}
	ApiTypeC2 = ApiType{code: "C2", impl: ApiImplB{}}
)

type ApiType struct {
	code string
	impl interface{}
}

// 这里可以使用反射优化
func CreateApi(api ApiType) Api {
	switch api {
	case ApiTypeC1:
		return &ApiImplA{}
	case ApiTypeC2:
		return &ApiImplB{}
	}
	return nil
}
