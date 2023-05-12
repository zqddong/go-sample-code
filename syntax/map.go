package main

import "fmt"

func main() {
	// 先声明 再使用
	var m1 map[string]string
	// nil map 不能赋值 需要make分配内存
	m1 = make(map[string]string)
	m1["a"] = "aa"
	m1["b"] = "bb"
	fmt.Println(m1)

	// 直接创建
	m2 := make(map[string]string)
	m2["a"] = "aa"
	m2["b"] = "bb"
	fmt.Println(m2)

	// 初始化+赋值
	m3 := map[string]string{
		"aa": "aa",
		"bb": "bb",
	}
	fmt.Println(m3)

	// 查找
	if v, ok := m1["a"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("key not found")
	}

	// 遍历
	for k, v := range m1 {
		fmt.Println(k, v)
	}
}
