package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Post struct {
	XMLName xml.Name `xml:"post"`
	Id      string   `xml:"id,attr"`
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func main() {
	post := Post{
		Id:      "1",
		Content: "HELLO WORLD!",
		Author: Author{
			Id:   "2",
			Name: "Sau Sheong",
		},
	}

	xmlfile, err := os.Create("post.xml")
	if err != nil {
		fmt.Println("Error creating XML file:", err)
		return
	}
	encoder := xml.NewEncoder(xmlfile)
	encoder.Indent("", "\t\t")
	err = encoder.Encode(&post)
	if err != nil {
		fmt.Println("Error writing XML to file:", err)
		return
	}
}
