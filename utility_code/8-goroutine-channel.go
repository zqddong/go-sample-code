package main

import "fmt"

// 出于性能考虑的建议：
// 实践经验表明，为了使并行运算获得高于串行运算的效率，在协程内完成的工作量
// 必须远远高于协程的创建和相互通信的开销
// 出于性能考虑建议使用带缓存的通道：
// 		使用带缓存的通道可以轻易成倍提高吞吐量，某些场景其性能可以提高至10倍甚至
// 		更多，通过调整通道的容量，可以进一步优化性能
//	 限制一个通道的数据量并将他们封装成一个数组：
//		如果使用通道传递大量单独的数据，那么通道将变成性能瓶颈，然后将数据块打包
//		封装成数组，在接收端解压数据时，性能可以提高至10倍
func main() {
	// for for-range 遍历通道
	ch := make(chan int, 10)
	ch <- 1
	ch <- 2
	ch <- 3

	for v := range ch {
		fmt.Println(v)
	}

	// 检测 ch 是否关闭
	//for {
	//	if input, open := <-ch; !open {
	//		break
	//	} else {
	//		fmt.Println(input)
	//	}
	//}

	// 通过一个通道让主程序等待直到协程完成（信号量模式）：
	ch2 := make(chan int)

	// 在goroutine中启动一些东西;当它完成时，在频道上发出信号
	go func() {
		do()
		// 如果希望程序一直阻塞，省略 ch <- 1 即可
		ch2 <- 1 // 发送信号;价值并不重要。
	}()
	doSomethingElseForAWhile()
	<-ch2 // 等待goroutine完成;丢弃发送值。

	// 通道的工厂模板
	pump()

	//
	//
}

// 通道工厂，启动一个匿名函数作为协程以生产通道
func pump() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func do() {

}

func doSomethingElseForAWhile() {

}

// 限制同时处理的请求数量 -- start
func ctlMaxTasks() {
	svc := make(chan *Request)
	go server(svc)
}

const MaxReq = 50

var sem = make(chan int, MaxReq)

type Request struct {
	a, b  int
	reply chan int
}

func process(r *Request) {
	// do something
}

func handler(r *Request) {
	sem <- 1 // 占位
	process(r)
	<-sem // 移除
}

func server(service chan *Request) {
	for {
		request := <-service
		go handler(request)
	}
}

// 限制同时处理的请求数量 -- end

// 单向channel用法 -- start

func testChan() {
	// 默认 channel 是双向的
	//var ch1 chan int
	//ch1 := make(chan int)

	// 单向写
	// var sendCh chan <- int
	// sendCh := make(chan <- int)

	// 单向读
	// var recvCh <- chan int
	// recvCh := make(<-chan int)
}

// chan<- 只写
func write(ch chan<- int, number int) {
	ch <- number
}

// <-chan 只读
func read(ch <-chan int) {
	for num := range ch {
		fmt.Println(num)
	}
}

// 单向channel用法 -- end
