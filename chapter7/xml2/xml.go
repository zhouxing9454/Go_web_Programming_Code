package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Post struct {
	XMLName xml.Name `xml:"post"`
	Id      string   `xml:"id,attr"`
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`
	Xml     string   `xml:",innerxml"`
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func main() {
	xmlFile, err := os.Open("D:\\Golang_workspace\\src\\Go_web_Programming_Code\\chapter7\\xml2\\post.xml")
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()
	xmlData, err := io.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Error reading XML file:", err)
		return
	}
	var post Post
	xml.Unmarshal(xmlData, &post)
	fmt.Println(post.Author.Name)
}
