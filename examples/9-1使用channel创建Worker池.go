package main

import (
	"fmt"
	"strconv"
	"time"
)

type Worker struct {
	PoolNum   uint32
	TaskQueue []chan Task
	QueueLen  uint32
}

type Task struct {
	ID   int64
	Type uint8
	Name string
}

func main() {
	worker := NewWorker()

	// 启动 worker 池子
	worker.StartWorkerPool()

	time.Sleep(3 * time.Second)

	// 塞任务
	for i := 0; i < 10; i++ {
		t := Task{ID: int64(i), Type: 1, Name: "Task" + strconv.Itoa(i)}
		worker.SendTaskToWorker(uint32(i), t)
	}

	//time.Sleep(100 * time.Second)
	for {
	}
}

func NewWorker() *Worker {
	// 一个pool 对应 一个 task queue 每个 queue 可缓冲 10个任务
	return &Worker{
		PoolNum:   5,
		TaskQueue: make([]chan Task, 5),
		QueueLen:  2,
	}
}

func (w *Worker) StartWorkerPool() {
	for i := 0; i < int(w.PoolNum); i++ {
		w.TaskQueue[i] = make(chan Task, w.QueueLen)
		go w.StartOneWorker(i, w.TaskQueue[i])
		fmt.Println("== StartWorkerPool Worker ID： ", i, " created ==")
	}
}

func (w *Worker) StartOneWorker(workerId int, taskQueue chan Task) {
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

func (w *Worker) Do(workerId int, t Task) {
	time.Sleep(5 * time.Second)
	fmt.Println("== Do run in workerId: ", workerId, " task detail: ", t, " ==")
}

func (w *Worker) SendTaskToWorker(taskId uint32, t Task) {
	// 判断chan缓存的任务数量 找最少的chan执行任务
	// len(w.TaskQueue[i])

	workerId := taskId % w.PoolNum
	fmt.Println("== SendTaskToWorker workerId: ", workerId, " data: ", t, " ==")
	w.TaskQueue[workerId] <- t
}
