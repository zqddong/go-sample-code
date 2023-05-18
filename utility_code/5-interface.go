package main

import "fmt"

type Animal interface {
	Run()
	Say()
}

type Dog struct {
	Name string
}

func (d Dog) Run() {
	fmt.Println(d.Name, " Run")
}
func (d Dog) Say() {
	fmt.Println(d.Name, " Say")
}

func main() {
	dog := Dog{Name: "哈士奇"}
	var animal interface{} = dog
	// 检测一个值 v 是否实现了接口
	if _, ok := animal.(Animal); ok {
		fmt.Println("var animal implement Animal interface")
	}

	dog.Say()
	dog.Run()
}

// classifier 使用接口实现一个类型分类函数：
func classifier(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("param #%d is a bool\n", i)
		case float64:
			fmt.Printf("param #%d is a float64\n", i)
		case int, int64:
			fmt.Printf("param #%d is an int\n", i)
		case nil:
			fmt.Printf("param #%d is nil\n", i)
		case string:
			fmt.Printf("param #%d is a string\n", i)
		default:
			fmt.Printf("param #%d’s type is unknown\n", i)
		}
	}
}
