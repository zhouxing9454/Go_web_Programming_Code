package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("D:\\Golang_workspace\\src\\Go_web_programming_instance_code\\chapter5\\iterator\\layout.html")
	daysofweek := []string{"mon", "tue", "wed", "thu", "fri", "sat", "sun", ""}
	t.Execute(w, daysofweek)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
