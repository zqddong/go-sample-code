package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
)

var (
	Web   = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)

type Result string
type Search func(ctx context.Context, query string) (Result, error)

func fakeSearch(kind string) Search {
	return func(_ context.Context, query string) (Result, error) {
		return Result(fmt.Sprintf("%s result for %q", kind, query)), nil
	}
}

// https://pkg.go.dev/golang.org/x/sync/errgroup 包地址
// 并行 说明了使用 Group 来同步一个简单的并行任务
func main() {
	Google := func(ctx context.Context, query string) ([]Result, error) {
		g, ctx := errgroup.WithContext(ctx)

		searches := []Search{Web, Image, Video}
		results := make([]Result, len(searches))
		for i, search := range searches {
			i, search := i, search // https://golang.org/doc/faq#closures_and_goroutines
			g.Go(func() error {
				result, err := search(ctx, query)
				if err == nil {
					results[i] = result
				}
				return err
			})
		}
		if err := g.Wait(); err != nil {
			return nil, err
		}
		return results, nil
	}

	results, err := Google(context.Background(), "golang")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	for _, result := range results {
		fmt.Println(result)
	}
}

func main2() {
	var eg errgroup.Group

	//匿名函数将会通过GO关键字启动一个协程
	eg.Go(func() error {
		fmt.Print("task 1\n")
		return nil
	})

	eg.Go(func() error {
		fmt.Print("task 2\n")
		return fmt.Errorf("task 2 error")
	})

	// 使用Wait 等待所有的协程执行完毕后，再进行后面的逻辑，同时可以记录两个协程的错误
	if err := eg.Wait(); err != nil {
		fmt.Printf("some error occur: %s\n", err.Error())
	}

	fmt.Print("over")
}

// JustErrors 说明了使用 Group 代替 sync.WaitGroup 来简化 goroutine 计数和错误处理。此示例派生自
func JustErrors() {
	g := new(errgroup.Group)
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}
	for _, url := range urls {
		// Launch a goroutine to fetch the URL.
		url := url // https://golang.org/doc/faq#closures_and_goroutines
		g.Go(func() error {
			// Fetch the URL.
			resp, err := http.Get(url)
			if err == nil {
				resp.Body.Close()
			}
			return err
		})
	}
	// Wait for all HTTP fetches to complete.
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully fetched all URLs.")
	}
}
