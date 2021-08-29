package main

import (
	"fmt"
	"sort"
)

// golang sort包下的内置排序, 类似java的compare
// 只对slice生效
func main() {

	// int 类型排序, 升序
	intSort()

	// string 类型排序, 升序
	stringSort()

	// 倒序
	intReverseSort()

	// 稳定排序, 看相等时, 是否调换位置
	stableSort()

}

// 稳定排序
func stableSort() {

	p := personSlice{
		person{Money: 4},
		person{Money: 1},
		person{Money: 3},
		person{Money: 3},
		person{Money: 9},
	}

	fmt.Println("Before sorted: ", p)

	sort.Stable(p)

	fmt.Println("After sorted: ", p)
}

type person struct {
	Money int
	Age   int
}
type personSlice []person

func (p personSlice) Len() int {
	return len(p)
}

func (p personSlice) Less(i, j int) bool {
	return p[i].Money < p[j].Money
}

func (p personSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// int类型倒序
func intReverseSort() {

	a := []int{4, 3, 2, 1, 5, 9, 8, 7, 6}
	fmt.Println("Before sorted: ", a)
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	fmt.Println("After sorted: ", a)
}

// string 类型排序
func stringSort() {
	str := []string{"d", "b", "c"}
	fmt.Println("Before sorted: ", str)
	sort.Strings(str)
	fmt.Println("After sorted: ", str)
}

// int slice
type IntSlice []int

func (s IntSlice) Len() int { return len(s) }

func (s IntSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }

// 整数排序
func intSort() {

	a := []int{4, 3, 2, 1, 5, 9, 8, 7, 6}
	fmt.Println("Before sorted: ", a)
	sort.Sort(IntSlice(a))
	fmt.Println("After sorted: ", a)
}
