package chapter4

import (
	"fmt"
	"io"
	"net/http"
)

func Process(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024)
	fileHeader := r.MultipartForm.File["uploaded"]
	for _, file := range fileHeader {
		filedata, err := file.Open()
		if err == nil {
			data, err := io.ReadAll(filedata)
			if err == nil {
				fmt.Fprintln(w, string(data))
			}
		}
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", Process)
	server.ListenAndServe()
}

//FormFile
