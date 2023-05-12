package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/sync/singleflight"
)

func main() {
	g := singleflight.Group{}

	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			val, err, shared := g.Do("a", a)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("index: %d, val: %d, shared: %v\n", j, val, shared)
		}(i)
	}

	wg.Wait()

}

var (
	count = int64(0)
)

// 模拟接口方法
func a() (interface{}, error) {
	time.Sleep(time.Millisecond * 500)
	return atomic.AddInt64(&count, 1), nil
}
