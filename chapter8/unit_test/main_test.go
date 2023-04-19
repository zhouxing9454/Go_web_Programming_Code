package main

import (
	"testing"
	"time"
)

func TestDecode(t *testing.T) {
	post, err := decode("D:\\Golang_workspace\\src\\Go_web_Programming_Code\\chapter8\\unit_test\\post.json")
	if err != nil {
		t.Error(err)
	}
	if post.Id != 1 {
		t.Error("Wrong id,was expecting 1 but got", post.Id)
	}
	if post.Content != "Hello World!" {
		t.Error("wrong content,was expecting 'Hello World!' but got", post.Content)
	}
}

func TestEncode(t *testing.T) {
	t.Skip("Skipping encoding for now")
}

/*
使用testing.Short()函数检查是否在短模式下运行测试。
如果是，则跳过长时间运行的测试并输出“Skipping long-running test in short mode”的信息。
如果不在短模式下运行测试，则使用time包中的Sleep()函数暂停程序执行10秒钟，以模拟长时间运行的测试。
这个测试函数没有实际的测试操作，但是它可以用来测试测试运行的时间。
*/
func TestLongRunningTest(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping long-running test in short mode")
	}
	time.Sleep(10 * time.Second)
}
