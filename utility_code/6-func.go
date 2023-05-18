package main

import "fmt"

func main() {
	protect(test)

	fmt.Println("main done")
}

// 使用内建函数 recover() 终止 panic() 过程
func protect(g func()) {
	defer func() {
		fmt.Println("done")
		if x := recover(); x != nil {
			fmt.Printf("run time panic: %v \n", x)
		}
	}()
	fmt.Println("start")

	g()
}

func test() {
	panic("test panic")
}
