package _interface

import "fmt"

type Behavior interface {
	Run() string
	Eat() string
}

// 通过 接口定义变量
func testBehavior() {
	// 接口定义变量 方式一
	var b Behavior
	b = new(Dog)
	b.Run()

	// 方式二
	dog := new(Dog)
	action(dog)

	cat := new(Cat)
	action(cat)
}

// 多态 根据传入实例执行相应的动作
func action(b Behavior) string {
	b.Run()
	b.Eat()
	return ""
}

type Animal struct {
	Color string
}

type Dog struct {
	Animal // 组合
	ID     int
	Name   string
	Age    int
}

// 初始化结构体
func test() {
	//方式一
	//var dog Dog
	//dog.Age = 3
	//dog.Name = "Tt"
	//dog.ID = 10

	//方式二
	//dog :=Dog{ ID:   1, Name: "df", Age:  12}

	//方式三 返回指针
	dog := new(Dog)
	dog.ID = 10
	dog.Name = "Cd"
	dog.Age = 11
	dog.Color = "red"

	fmt.Println("dog: ", dog)
}

func (d *Dog) Run() string {
	fmt.Println("ID: ", d.ID)
	return ""
}

func (d *Dog) Eat() string {
	fmt.Println("Eat: ")
	return ""
}

type Cat struct {
	Animal // 组合
	ID     int
	Name   string
	Age    int
}

func (d *Cat) Run() string {
	fmt.Println("ID: ", d.ID)
	return ""
}

func (d *Cat) Eat() string {
	fmt.Println("Eat: ")
	return ""
}
