package main

import (
	"fmt"
	"strconv"
	"time"
)

type Worker struct {
	PoolNum   uint32
	TaskQueue []chan interface{}
	QueueLen  uint32
}

func main() {
	worker := NewWorker()

	// 启动 worker 池子
	worker.StartWorkerPool()

	time.Sleep(3 * time.Second)

	// 塞任务
	for i := 0; i < 100; i++ {
		worker.SendTaskToWorker(uint32(i), "任务 "+strconv.Itoa(i))
	}

	time.Sleep(100 * time.Second)
	//select {}
}

func NewWorker() *Worker {
	// 一个pool 对应 一个 task queue 每个 queue 可缓冲 10个任务
	return &Worker{
		PoolNum:   5,
		TaskQueue: make([]chan interface{}, 5),
		QueueLen:  2,
	}
}

func (w *Worker) StartWorkerPool() {
	for i := 0; i < int(w.PoolNum); i++ {
		w.TaskQueue[i] = make(chan interface{}, w.QueueLen)
		go w.StartOneWorker(i, w.TaskQueue[i])
		fmt.Println("== StartWorkerPool Worker ID： ", i, " created ==")
	}
}

func (w *Worker) StartOneWorker(workerId int, taskQueue chan interface{}) {
	fmt.Println("++ StartOneWorker workerId： ", workerId, " started ++")

	for {
		select {
		case oneTask := <-taskQueue:
			w.Do(workerId, oneTask)
		}
	}
}

//func (w *Worker) Do(workerId int, fn func() (interface{}, error)) {
//	time.Sleep(5 * time.Second)
//	fn()
//}

func (w *Worker) Do(workerId int, data interface{}) {
	time.Sleep(5 * time.Second)
	fmt.Println("== Do run in workerId: ", workerId, " task detail: ", data, " ==")
}

func (w *Worker) SendTaskToWorker(taskId uint32, data interface{}) {
	workerId := taskId % w.PoolNum
	fmt.Println("== SendTaskToWorker workerId: ", workerId, " data: ", data, " ==")
	w.TaskQueue[workerId] <- data
}
