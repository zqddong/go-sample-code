package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func main() {
	// 创建
	human := new(Human)
	fmt.Println(human)

	// 初始化
	h1 := &Human{"Go", 18}
	fmt.Println(h1)

	// 为每个结构体定义一个构建函数，并推荐使用构建函数初始化结构体
	h2 := NewHuman("Good", 19)
	fmt.Println(h2)
}

func NewHuman(name string, age int) *Human {
	return &Human{name, age}
}
