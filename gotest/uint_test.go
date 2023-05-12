package gotest_test

import (
	"github.com/zqddong/go-sample-code/gotest"
	"testing"
)

// 编写一个单元测试并执行是非常方便的，只需要遵循一定的规则:
// 测试文件名必须以”_test.go”结尾;
// 测试函数名必须以“TestXxx”开始;
// 命令行下使用”go test”即可启动测试;
func TestAdd(t *testing.T) {
	var a = 1
	var b = 2
	var expected = 3

	actual := gotest.Add(a, b)
	if actual != expected {
		t.Errorf("Add(%d, %d) = %d; expected: %d", a, b, actual, expected)
	}
}
