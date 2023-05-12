package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	taskChan := make(chan bool, 10)

	wg := sync.WaitGroup{}
	defer close(taskChan)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		taskChan <- true
		go func(i int) {
			defer wg.Done()
			// Do something
			//Do()
			fmt.Printf("并行10: %d \n", i)
			time.Sleep(10 * time.Second)
			<-taskChan
		}(i)
	}

	wg.Wait()
}

// 原理：
// WaitGroup主要维护2个计数器，一个是请求计数器v，一个是等待计数器w，二者组成一个64bit的值，请求计数器占高32bit，等待计数器占低32bit
// 每次Add执行，请求计数器v加1，Done方法执行，等待计数器减1，v为0时通过信号量唤醒Wait（）

//type WaitGroup struct {
//	noCopy noCopy
//
//	// 64-bit value: high 32 bits are counter, low 32 bits are waiter count.
//	// 64-bit atomic operations require 64-bit alignment, but 32-bit
//	// compilers only guarantee that 64-bit fields are 32-bit aligned.
//	// For this reason on 32 bit architectures we need to check in state()
//	// if state1 is aligned or not, and dynamically "swap" the field order if
//	// needed.
//	state1 uint64
//	state2 uint32
//}
