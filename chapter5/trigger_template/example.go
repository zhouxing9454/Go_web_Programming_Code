package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("D:\\Golang_workspace\\src\\Go_web_programming_instance_code\\chapter5\\trigger_template\\layout.html")
	t.Execute(w, "hello world!")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}

//ParseFiles
//ParseGlob
//New,Parse
//Execute
//ExecuteTemplate
