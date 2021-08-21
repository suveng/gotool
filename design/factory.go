package design

import (
	"errors"
	"fmt"
)

// 工厂方法
// 根据传入参数构建出接口的实现类, 内部可以控制初始化的逻辑
type Restaurant interface {
	GetFood()
}

type Sx struct {
}

func (s *Sx) GetFood() {
	fmt.Println("sha xian xiao chi factory")
}

type Sh struct {
}

func (s *Sh) GetFood() {
	fmt.Println("shang hai factory")
}

// factory build
func build(name string) (Restaurant, error) {
	switch name {
	case "1":
		{
			fmt.Println(1)
			return &Sx{}, nil
		}
	case "2":
		{
			fmt.Println(2)
			return &Sh{}, nil
		}
	}

	return nil, errors.New("type not found")
}
