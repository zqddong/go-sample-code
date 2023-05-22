package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	err := errors.New("test err")
	//err := new(error)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err != nil {
		panic(err)
	}
}
