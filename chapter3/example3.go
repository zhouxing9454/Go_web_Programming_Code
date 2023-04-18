package chapter3

import "net/http"

func main() {
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}
	server.ListenAndServeTLS("cert.pem", "key.pem")
	//前者是SSL证书，后者是服务器的私钥
	//可以使用标准库crypto生成
}
