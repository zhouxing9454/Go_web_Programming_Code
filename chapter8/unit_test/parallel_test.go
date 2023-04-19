package main

import (
	"testing"
	"time"
)

func TestParallel_1(t *testing.T) {
	t.Parallel() //调用Parallel函数，以并行方式运行测试用例
	time.Sleep(1 * time.Second)
}

func TestParallel_2(t *testing.T) {
	t.Parallel()
	time.Sleep(2 * time.Second)
}

func TestParallel_3(t *testing.T) {
	t.Parallel()
	time.Sleep(3 * time.Second)
}

//go test -v -cover -short -parallel 3
//- parallel指示go以并行方式运行测试用例，参数3表示最多并行运行3个测试用例
