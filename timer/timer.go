package main

import (
	"log"
	"time"
)

// 创建定时器 func NewTimer(d Duration) *Timer

// 停止定时器 func (t *Timer) Stop() bool
// true: 定时器超时前停止，后续不会再有事件发送;
// false: 定时器超时后停止;

// 重置定时器 func (t *Timer) Reset(d Duration) bool
// 重置的动作实质上是先停掉定时器，再启动。其返回值也即停掉计时器的返回值

// Timer内容总结如下:
// time.NewTimer(d)创建一个Timer;
// timer.Stop()停掉当前Timer;
// timer.Reset(d)重置当前Timer;

func main() {
	//DelayFunc()
	//AfterDemo()
	AfterFuncDemo()
}

// DelayFunc 延迟执行某个方法
func DelayFunc() {
	timer := time.NewTimer(5 * time.Second)

	select {
	case <-timer.C:
		log.Println("Delayed 5s, start to do something.")
	}
}

// AfterDemo 有时我们就是想等指定的时间，没有需求提前停止定时器，
// 也没有需求复用该定时器，那么可以使用匿名的定时器
func AfterDemo() {
	log.Println(time.Now())
	<-time.After(1 * time.Second)
	log.Println(time.Now())
}

// AfterFuncDemo 与上面的例子所不同的是，time.AfterFunc()是异步执行的，
// 所以需要在函数最后sleep等待指定的协程退出，否 则可能函数结束时协程还未执行。
func AfterFuncDemo() {
	log.Println("AfterFuncDemo start: ", time.Now())
	time.AfterFunc(1*time.Second, func() {
		log.Println("AfterFuncDemo end: ", time.Now())
	})
	time.Sleep(2 * time.Second) // 等待协程退出
}
