package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comment"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	jsonfile, err := os.Open("D:\\Golang_workspace\\src\\Go_web_Programming_Code\\chapter7\\json2\\post.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonfile.Close()

	decoder := json.NewDecoder(jsonfile)
	for {
		var post Post
		err := decoder.Decode(&post)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error decoding JSON file:", err)
			return
		}
		fmt.Println(post)
	}
}
