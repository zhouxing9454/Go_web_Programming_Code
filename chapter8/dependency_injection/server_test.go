package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetPost(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest(&FakePost{}))

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/post/1", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
	var post Post
	json.Unmarshal(writer.Body.Bytes(), &post)
	if post.Id != 1 {
		t.Error("Cannot retrieve JSON post")
	}
}

func TestPutPost(t *testing.T) {
	mux := http.NewServeMux()
	post := &FakePost{}
	mux.HandleFunc("/post/", handleRequest(post))

	writer := httptest.NewRecorder()
	jsondata := strings.NewReader(`{"content":"Updated post","author":"Sau sheong"}`)
	request, _ := http.NewRequest("PUT", "/post/1", jsondata)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}

//依赖注入
/*
type Greeter interface {
    Greet() string
}

type EnglishGreeter struct {}

func (eg EnglishGreeter) Greet() string {
    return "Hello, World!"
}

type GreeterClient struct {
    greeter Greeter
}

func NewGreeterClient(greeter Greeter) *GreeterClient {
    return &GreeterClient{greeter: greeter}
}

func (gc GreeterClient) SayHello() string {
    return gc.greeter.Greet()
}

func main() {
    greeter := EnglishGreeter{}
    greeterClient := NewGreeterClient(greeter)
    fmt.Println(greeterClient.SayHello())
}
*/
