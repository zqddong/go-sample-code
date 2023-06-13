package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	test1()
	test2()
	test3()
	test4()
}

func test1() {
	var s string

	for i := 0; i < 10; i++ {
		s += "a"
	}

	fmt.Println(s)
}

// Go 1.10版本后引入了一个strings.Builder类型
// 性能比 += 高三到四个数量级
func test2() {
	var str strings.Builder

	for i := 0; i < 10; i++ {
		str.WriteString("a")
	}

	fmt.Println(str.String())
}

// Go1.10 以前使用的是bytes.Buffer
func test3() {
	var buffer bytes.Buffer

	for i := 0; i < 10; i++ {
		buffer.WriteString("a")
	}

	fmt.Println(buffer.String())
}

// strings.Join  将字符串切片中存在的所有元素连接为单个字符串
// strings.Split 将指定的分隔符切割字符串，并返回切割后的字符串切片
func test4() {
	str := ""

	sli := strings.Split(str, "")
	fmt.Println(sli)

	for i := 0; i < 10; i++ {
		sli = append(sli, "a")
	}
	fmt.Println(strings.Join(sli, ""))
}
