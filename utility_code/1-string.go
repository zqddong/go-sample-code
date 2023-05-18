package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	// 1. 如何修改字符串的一个字符
	str := "hello"
	c := []byte(str)
	c[0] = 'c'
	s2 := string(c) // s2 == "cello"
	fmt.Println(s2)

	// 2. 如何获取字符串的子串
	n := 1
	m := 3
	substr := str[n:m]
	fmt.Println(substr)

	// 3. for for-range 遍历一个字符串
	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
	}

	for index, chr := range str {
		fmt.Println(index, " -- ", chr)
	}

	// 4. 获取一个字符串的字节数 len(str)
	fmt.Println(len(str))

	// 4.1 获取一个字符串的字符数 utf8.RuneCountInString(str) 速度快
	fmt.Println(utf8.RuneCountInString(str))

	// 4.2 获取一个字符串的字节数 len([]rune(str))
	fmt.Println(len([]rune(str)))

	// 5.1 连接字符串 bytes.Buffer
	var buffer bytes.Buffer
	buffer.WriteString("H")
	buffer.WriteString("e")
	buffer.WriteString("l")
	buffer.WriteString("l")
	buffer.WriteString("o")

	// 这种实现方式比使用 += 要更节省内存和 CPU，尤其是要串联的字符串数目特别多的时候
	fmt.Print(buffer.String(), "\n")

	// 5.2 连接字符串 +=
	str1 := "Hello "
	str2 := "Go"
	str1 += str2
	fmt.Println(str1)

	// 5.3 连接字符串 String.Join
	str3 := "Hello Go World"
	strSli := strings.Split(str3, " ")

	fmt.Println(strings.Join(strSli, "--"))
}
