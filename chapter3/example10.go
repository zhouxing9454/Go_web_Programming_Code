package chapter3

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Hellorouter(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "hello,%s!\n", p.ByName("name")) //获取具名参数
}

func main() {
	mux := httprouter.New()              //创建一个多路复用器
	mux.GET("/hello/:name", Hellorouter) //具名参数
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux, //绑定
	}
	server.ListenAndServe()
}
