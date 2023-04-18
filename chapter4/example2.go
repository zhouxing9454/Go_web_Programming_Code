package main

import (
	"fmt"
	"net/http"
)

func body(w http.ResponseWriter, r *http.Request) {
	length := r.ContentLength
	body := make([]byte, length)
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/body", body)
	server.ListenAndServe()
}

//curl -id "first_name=zhou&last_name=xing" 127.0.0.1:8080/body

/*
result:
HTTP/1.1 200 OK
Date: Mon, 17 Apr 2023 06:43:58 GMT
Content-Length: 31
Content-Type: text/plain; charset=utf-8

first_name=zhou&last_name=xing
*/
