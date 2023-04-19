package main

import "testing"

// benchmarking the decode function
func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		decode("D:\\Golang_workspace\\src\\Go_web_Programming_Code\\chapter8\\unit_test\\post.json")
	}
}

//goos: windows
//goarch: amd64
//pkg: Go_web_Programming_Code/chapter8/unit_test
//cpu: Intel(R) Core(TM) i5-10300H CPU @ 2.50GHz
//BenchmarkDecode
//BenchmarkDecode-8          12014            100019 ns/op

func BenchmarkUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unmarshal("D:\\Golang_workspace\\src\\Go_web_Programming_Code\\chapter8\\unit_test\\post.json")
	}
}
