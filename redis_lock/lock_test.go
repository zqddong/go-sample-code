package redis_lock_test

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/zqddong/go-sample-code/redis_lock"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func BenchmarkRedisLock(b *testing.B) {
	x := b.N
	wg := sync.WaitGroup{}
	wg.Add(x)
	now := time.Now()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	var n int64
	var fail int64
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	for i := 0; i < x; i++ {
		go func() {
			defer wg.Done()
			rl := redis_lock.New(ctx, rdb, "test_lock")
			if err := rl.Lock(); err != nil {
				atomic.AddInt64(&fail, 1)
				cancel()
				return
			}
			defer rl.Unlock()
			// 模拟耗时业务
			time.Sleep(300 * time.Millisecond)
			atomic.AddInt64(&n, 1)
		}()
	}

	wg.Wait()
	fmt.Printf("测试数：%d\t成功数：%d\t锁失败：%d\t结果：%t\t耗时：%d \n", x, n, fail, x == int(n), time.Since(now).Milliseconds())
}
