package chapter3

import "net/http"

func main() {
	http.ListenAndServe("", nil)
}
