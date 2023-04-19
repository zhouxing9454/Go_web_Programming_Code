package main

import (
	"encoding/json"
	. "gopkg.in/check.v1"
	"net/http"
	"net/http/httptest"
	"testing"
)

type PostTestSuit struct {
}

func init() {
	Suite(&PostTestSuit{})
}

func Test(t *testing.T) {
	TestingT(t)
}

func (s *PostTestSuit) TestHandleGet(c *C) {
	mux := http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest(&FakePost{}))
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/post/1", nil)
	mux.ServeHTTP(writer, request)

	c.Check(writer.Code, Equals, 200)
	var post Post
	json.Unmarshal(writer.Body.Bytes(), &post)
	c.Check(post.Id, Equals, 1)
}

//go test -check.vv
//https://zhuanlan.zhihu.com/p/45570168
