package main

import (
	"html/template"
	"net/http"
	"time"
)

func formatDate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}

func process(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{"fdate": formatDate}
	t := template.New("D:\\Golang_workspace\\src\\Go_web_programming_instance_code\\chapter5\\custom_function\\layout.html").Funcs(funcMap)
	t, _ = template.ParseFiles("D:\\Golang_workspace\\src\\Go_web_programming_instance_code\\chapter5\\custom_function\\layout.html")
	t.Execute(w, time.Now())
}

func main() {
	server := http.Server{Addr: "127.0.0.1:8080"}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
