package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func process(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	var t *template.Template
	if rand.Intn(10) > 5 {
		t, _ = template.ParseFiles("D:\\Golang_workspace\\src\\Go_web_programming_instance_code\\chapter5\\nest1\\layout.html", "D:\\Golang_workspace\\src\\Go_web_programming_instance_code\\chapter5\\nest1\\red_hello.html")
	} else {
		t, _ = template.ParseFiles("D:\\Golang_workspace\\src\\Go_web_programming_instance_code\\chapter5\\nest1\\layout.html")
	}
	t.ExecuteTemplate(w, "layout", "")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
