package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	openFile()

}

// 打开文件并读取
func openFile() {
	file, err := os.Open("7-file.go")
	if err != nil {
		fmt.Println("file open err: ", err)
		return
	}

	defer file.Close()
	iReader := bufio.NewReader(file)
	for {
		str, err := iReader.ReadString('\n')
		if err != nil {
			fmt.Println("ReadString err:", err)
			return
		}
		fmt.Printf("the input was: %s", str)
	}
}
