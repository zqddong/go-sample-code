package main

import "fmt"

func main() {
	// 创建 map := make(map[keyType]valueType)

	// 初始化
	m := map[string]int{"one": 1, "two": 2}
	// 遍历
	for k, v := range m {
		fmt.Println(k, " -- ", v)
	}

	// 检查键是否存在
	v, has := m["one"]
	fmt.Println(v, " -- ", has)

	// 删除
	delete(m, "one")
}
