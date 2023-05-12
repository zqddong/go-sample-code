package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
)

func main() {
	//test1()
	//test2()
	test3()
}

// --------------------------------------------------------- 3
// 利用无缓冲channel与任务发送/执行分离方式
func test3() {
	ch := make(chan int, 3)

	goCnt := 3 //启动goroutine的数量
	for i := 0; i < goCnt; i++ {
		go do3(ch)
	}

	num := math.MaxInt64
	for t := 0; t < num; t++ {
		sendTask(t, ch)
	}

	wg.Wait()
}

func sendTask(task int, ch chan int) {
	wg.Add(1)
	// channel 缓冲为3 满了就会阻塞
	ch <- task
}

func do3(ch chan int) {
	for t := range ch {
		fmt.Println("go func, ", t, "goroutine count: ", runtime.NumGoroutine())
		wg.Done()
	}
}

// --------------------------------------------------------- 3

// --------------------------------------------------------- 2
// channel与sync同步组合方式

var wg = sync.WaitGroup{}

func test2() {
	num := math.MaxInt64
	//num = 10

	ch := make(chan bool, 3)

	for i := 0; i < num; i++ {
		wg.Add(1)
		ch <- true

		go do2(ch, i)
	}

	wg.Wait()
}

func do2(ch chan bool, i int) {
	fmt.Println("go func, ", i, "goroutine count: ", runtime.NumGoroutine())

	<-ch
	wg.Done()
}

// --------------------------------------------------------- 2

// --------------------------------------------------------- 1
// 利用 有缓冲的 channel 限制
func test1() {
	num := math.MaxInt64

	// 有缓冲的 channel 控制
	ch := make(chan bool, 3)

	for i := 0; i < num; i++ {
		ch <- true
		go do(ch, i)
	}
}

func do(ch chan bool, i int) {
	fmt.Println("go func, ", i, "goroutine count: ", runtime.NumGoroutine())
	<-ch
}

// --------------------------------------------------------- 1
