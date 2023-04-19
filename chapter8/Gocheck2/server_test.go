package main

import (
	"encoding/json"
	. "gopkg.in/check.v1"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type PostTestSuit struct {
	mux    *http.ServeMux
	post   *FakePost
	writer *httptest.ResponseRecorder
}

func init() {
	Suite(PostTestSuit{})
}

func Test(t *testing.T) {
	TestingT(t)
}

// SetUpTest 运行每个测试用例之前都会运行一次
// TearDownTest 运行每个测试用例之后都会运行一次
// SetUpSuite 在套件开始运行之前运行一次
// TearDownSuite 在所有测试都运行完毕之后运行一次
func (s *PostTestSuit) SetUpTest(c *C) {
	s.post = &FakePost{}
	s.mux = http.NewServeMux()
	s.mux.HandleFunc("/post/", handleRequest(s.post))
	s.writer = httptest.NewRecorder()
}

func (s *PostTestSuit) TearDownTest(c *C) {
	c.Log("Finished test -", c.TestName())
}

func (s *PostTestSuit) SetUpSuite(c *C) {
	c.Log("Starting Post Test Suite")
}

func (s *PostTestSuit) TearDownSuite(c *C) {
	c.Log("Finishing Post Test Suite")
}

func (s *PostTestSuit) TestHandleGet(c *C) {
	request, _ := http.NewRequest("GET", "/post/1", nil)
	s.mux.ServeHTTP(s.writer, request)

	c.Check(s.writer.Code, Equals, 200)
	var post Post
	json.Unmarshal(s.writer.Body.Bytes(), &post)
	c.Check(post.Id, Equals, 1)
}

func (s *PostTestSuit) TestHandlePut(c *C) {
	jsondata := strings.NewReader(`{"content":"Updated post","author":"Sau sheong"}`)
	request, _ := http.NewRequest("PUT", "/post/1", jsondata)
	s.mux.ServeHTTP(s.writer, request)

	c.Check(s.writer.Code, Equals, 200)
	c.Check(s.post.Id, Equals, 1)
	c.Check(s.post.Content, Equals, "Updated post")
}
