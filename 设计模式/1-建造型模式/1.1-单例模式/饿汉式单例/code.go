package main

import "sync"

type singleton struct{}

var instance *singleton
var once sync.Once

// GetInstance 懒汉式单例
func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})

	return instance
}

func main() {
	GetInstance()
}
