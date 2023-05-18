package main

import (
	"fmt"
	"go.uber.org/atomic"
	"golang.org/x/sync/errgroup"
	"time"
)

func main() {
	NewWorkerBuild().Start()
}

type TaskDetail struct {
	ID   int64
	Type uint8
	Name string
}

type WorkerBuild struct {
	interval time.Duration // 任务时间间隔
	p        int           // 控制并行的任务数量
	running  atomic.Int32
	receiver chan *TaskDetail
	stop     chan struct{}
}

type Processor struct {
	task *TaskDetail
}

func NewWorkerBuild() *WorkerBuild {
	w := &WorkerBuild{
		interval: 10 * time.Second,
		p:        2,
		receiver: make(chan *TaskDetail, 10),
		stop:     make(chan struct{}, 2),
	}

	go w.fetchTask()

	return w
}

func (w *WorkerBuild) fetchTask() {
	for {
		if w.running.Load() >= int32(w.p) {
			time.Sleep(10 * time.Second)
			continue
		}
		task, err := GetOneTask()
		if err != nil {
			time.Sleep(w.interval)
		}

		w.receiver <- task

		// 防止同一节点，接收到大量任务，均匀分散
		time.Sleep(5 * time.Second)
	}
}

func (w *WorkerBuild) Start() error {
	g := errgroup.Group{}

Loop:
	for {
		select {
		case task := <-w.receiver:
			p := NewProcessor(task)
			g.Go(func() error {
				w.running.Inc()
				defer w.running.Dec()

				w.handleProcessor(p)
				return nil
			})
		case <-w.stop:
			break Loop
		}
	}

	return g.Wait()
}

func NewProcessor(t *TaskDetail) *Processor {
	p := &Processor{
		task: t,
	}

	return p
}

func (w *WorkerBuild) handleProcessor(p *Processor) {
	err := p.Run()
	if err != nil {
		return
	}
}

func (p *Processor) Run() error {
	time.Sleep(1000 * time.Second)
	return nil
}

func GetOneTask() (*TaskDetail, error) {
	// TODO 走DB使用事务 获取 并更新状态 保证每次仅一条任务进来

	now := time.Now()
	//now := time.Now().Unix() // 时间戳
	date := now.Format("2006-01-02 15:04:05")

	t := &TaskDetail{
		ID:   1,
		Type: 1,
		Name: "测试: " + date,
	}
	fmt.Println("Get One Task: ", t)
	return t, nil
}
