package main

import (
	"fmt"
	"math"
)

func main() {
	x := 1.7

	// 上取整
	fmt.Println(ceil(x))
	// 下取整
	fmt.Println(floor(x))
	// 四舍五入
	fmt.Println(round(x))
}

// 上取整
func ceil(x float64) int {
	return int(math.Ceil(x))
}

// 下取整
func floor(x float64) int {
	return int(math.Floor(x))
}

// 四舍五入
func round(x float64) int {
	return int(math.Round(x))
}

// 一个奇葩的四舍五入方法
func round2(x float64) int {
	return int(math.Floor(x + 0.5))
}
