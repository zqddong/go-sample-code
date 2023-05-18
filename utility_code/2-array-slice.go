package main

import "fmt"

func main() {
	// 创建
	arr1 := new([5]int)
	fmt.Println(arr1)

	slice1 := make([]int, 5)
	fmt.Println(slice1)

	// 初始化
	arr2 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)

	arrKeyVal := [5]int{1: 4, 3: 5}
	fmt.Println(arrKeyVal)

	var sli []int = arr2[1:3]
	fmt.Println(sli)

	// 截断数组或者切片的最后一个元素
	fmt.Println(arr2[:len(arr2)-1])

	// for for-range 遍历数组或者切片
	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}

	for k, v := range arr2 {
		fmt.Println(k, " -- ", v)
	}

	// 在一个二维数组或者切片 arr2Dim 中查找一个指定值 v
	V := 99
	arr2Dim := [...][5]int{
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 99},
	}

	found := false
Found:
	for row := range arr2Dim {
		for column := range arr2Dim[row] {
			if arr2Dim[row][column] == V {
				found = true
				break Found
			}
		}
	}
	fmt.Println(found)
}
