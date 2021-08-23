package ioc

import "fmt"

type Service struct {
}

func (s *Service) say() {
	fmt.Println("say")
}

func (s *Service) register() {
	context := GetContext()
	add, err := context.add("test", &Service{})
	if err != nil {
		panic(err)
	}

	if !add {
		panic("添加bean失败")
	}
}

func init() {

}
