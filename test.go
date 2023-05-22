package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/zqddong/go-sample-code/redis_lock"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	//testChan()
	//testUnSafeMap()
	//go testVar()
	//
	//time.Sleep(10 * time.Second)

	//testArr()
	//te()

	//_ = RemoveSliElement()

	//go rlock(1)
	//go rlock(2)
	//
	//time.Sleep(5 * time.Second)
	//fmt.Println(numberOfSteps(123))
	//fmt.Println(restoreMatrix([]int{3, 8}, []int{4, 7}))
	//strings.Join()

	//var slice []int
	//slice = append(slice, 1, 2, 3)
	//
	//newSlice := AddElement(slice, 4)
	//fmt.Println(&slice[0] == &newSlice[0])
	//
	//fmt.Println("lenth of slice: ", len(slice))
	//fmt.Println("lenth of slice: ", len(newSlice))
	//fmt.Println("capacity of slice: ", cap(slice))
	//fmt.Println("capacity of slice: ", cap(newSlice))

	//
	//ch := make(chan int, 5)
	//for i := 0; i < 5; i++ {
	//	ch <- i
	//	fmt.Println(len(ch))
	//
	//}
	//time.Sleep(time.Second * 3)
	//
	//for i := 0; i < 3; i++ {
	//	getI, ok := <-ch
	//	if ok {
	//		fmt.Println("get i: ", getI, len(ch))
	//	}
	//}
	//
	//fmt.Println(len(ch))
}

func AddElement(slice []int, e int) []int {
	return append(slice, e)
}

// 按顺序打印 cat dog fish 打印10遍
// 利用无缓冲的channel 阻塞 依次打印
func printAnimalByOrder() {
	var wg sync.WaitGroup
	wg.Add(3)

	catCh := make(chan struct{})
	dogCh := make(chan struct{})
	fishCh := make(chan struct{})

	// 打印cat的协程
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			<-catCh
			fmt.Println("cat")
			dogCh <- struct{}{}
		}
	}()

	// 打印dog的协程
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			<-dogCh
			fmt.Println("dog")
			fishCh <- struct{}{}
		}
	}()

	// 打印fish的协程
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			<-fishCh
			fmt.Println("fish")
			if i < 9 {
				catCh <- struct{}{}
			} else {
				// 9 退出
				break
			}
		}
	}()

	// 启动第一个协程
	catCh <- struct{}{}

	wg.Wait()
}

func restoreMatrix(rowSum []int, colSum []int) [][]int {
	n, m := len(rowSum), len(colSum)
	ret := make([][]int, n)
	for i := range ret {
		ret[i] = make([]int, m)
	}
	fmt.Println("--1", ret)

	sum := 0
	for i, x := range colSum {
		sum += x
		ret[0][i] = x
	}
	fmt.Println("--2", ret)
	for i, row := range ret {
		if i == n-1 {
			break
		}
		t := sum - rowSum[i]
		sum -= rowSum[i]
		for j, x := range row {
			if t >= x {
				ret[i][j] = 0
				ret[i+1][j] = x
				t -= x
			} else {

				ret[i][j] = x - t
				ret[i+1][j] = t
				t = 0
			}
		}
	}
	return ret
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func numberOfSteps(num int) int {
	step := 0
	if num == 0 {
		return step
	}
	for {
		if num%2 == 0 {
			num = num / 2
			step++
			continue
		}
		if num%2 != 0 {
			num = num - 1
			step++
		}
		if num <= 0 {
			break
		}
	}

	return step
}

func rlock(g int) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	rl := redis_lock.New(ctx, rdb, "test_lock")
	defer rl.Unlock()
	err := rl.Lock()
	if err != nil {
		fmt.Println("加锁失败:", err, " g: ", g)
		cancel()
		return
	}
	rst := rl.Get("test_lock")
	fmt.Println("加锁成功:", rst, " g: ", g)

	time.Sleep(1 * time.Second)

	//rl.Unlock()

	//rst = rl.Get("test_lock")
	//fmt.Println(rst)
}

// RemoveSliElement 删除切片元素
func RemoveSliElement() error {
	intSlice := []int{1, 2, 3, 4, 5}
	value := 4

	for i, v := range intSlice {
		if value == v {
			intSlice = append(intSlice[:i], intSlice[i+1:]...)
		}
	}
	fmt.Println(intSlice)
	return nil
}

func te() {
	var i = 10
	fmt.Println("newGoRoutine-", i)

	var wg sync.WaitGroup

	ans := int64(0)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go newGoRoutine(&wg, &ans)
	}
	wg.Wait()
	time.Sleep(5 * time.Second)
	fmt.Println("newGoRoutine-", ans)

}

func newGoRoutine(wg *sync.WaitGroup, i *int64) {
	defer wg.Done()
	atomic.AddInt64(i, 1)
	return
}

// 5 协程 打印数组
func testArr() {
	var arr = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		sli := arr[i*2 : i*2+2]
		go func(i int, arr []int) {
			for _, v := range arr {
				fmt.Println("for-", i, "arr-", v)
			}
			wg.Done()
		}(i, sli)
	}

	wg.Wait()

	time.Sleep(10 * time.Second)
}

func test() {
	a := 0
	b := 0
	for {
		n, _ := fmt.Scan(&a, &b)
		fmt.Printf("n: %d \n", n)
		if n == 0 {
			break
		} else {
			fmt.Printf("a+b: %d \n", a+b)
		}
	}
}

func testChan() {
	ch := make(chan int, 10)
	go func() {
		for i := 0; i < 20; i++ {
			ch <- i
		}
	}()

	time.Sleep(2 * time.Second)
	close(ch)

	for {
		select {
		case v, ok := <-ch:
			if ok {
				fmt.Println(v)
			} else {
				fmt.Println("close")
				break
			}
		}
	}

}

var TestMap map[string]string
var TestSafeMap sync.Map

//var l sync.Mutex

func init() {
	TestMap = make(map[string]string, 1)
	TestSafeMap = sync.Map{}
}

func testUnSafeMap() {
	for i := 0; i < 1000; i++ {
		go Write("aaa")
		go Read("aaa")
		go Write("bbb")
		go Read("bbb")
		fmt.Println("--------------", i)
	}
	time.Sleep(5 * time.Second)
}
func Read(key string) {
	//l.Lock()
	//fmt.Println(TestMap[key])
	//l.Unlock()

	fmt.Println(TestSafeMap.Load(key))
}
func Write(key string) {
	//l.Lock()
	//TestMap[key] = key
	//l.Unlock()

	TestSafeMap.Store(key, key)
}

func testVar() {
	// 数组
	var arr1 [5]int
	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr1)
	fmt.Println(arr2)

	// slice
	var slc1 []int
	slc2 := make([]int, 5)
	fmt.Println(slc1)
	fmt.Println(slc2)
	fmt.Println(arr2[:])

	// map
	var map1 map[string]string
	map2 := make(map[string]string)

	fmt.Println(map1)
	fmt.Println(map2)

	return
}
