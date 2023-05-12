package gotest_test

import (
	"github.com/zqddong/go-sample-code/gotest"
	"testing"
)

// 编写并执行性能测试是非常简单的，只需要遵循一些规则:
// 文件名必须以“_test.go”结尾;
// 函数名必须以“BenchmarkXxx”开始;
// 使用命令“go test -bench=.”即可开始性能测试;
// go test -v -bench=. -benchmem -count=3
// go test -v -bench=BenchmarkRedisLock  -benchmem
func BenchmarkMakeSliceWithoutAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gotest.MakeSliceWithoutAlloc()
	}
}
func BenchmarkMakeSliceWithPreAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gotest.MakeSliceWithPreAlloc()
	}
}
