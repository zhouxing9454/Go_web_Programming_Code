package chapter3

import (
	"fmt"
	"net/http"
)

type helloHandler struct {
}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello zhouxing!")
}

func Log(h http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Printf("Handler called - %T\n", h)
		h.ServeHTTP(writer, request)
	})
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	hello := helloHandler{}
	http.Handle("/hello", Log(&hello))
	server.ListenAndServe()
}
