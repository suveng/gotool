package simple_factory

func main() {
	// 主业务
	// 获取接口实现
	api := CreateApi(ApitypeC1)
	// 接口调用
	api.Doing()
}
