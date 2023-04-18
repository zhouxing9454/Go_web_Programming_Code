package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

func main() {
	csvfile, err := os.Create("posts.csv")
	if err != nil {
		panic(err)
	}
	defer csvfile.Close()
	allPosts := []Post{
		{Id: 1, Content: "Hello World!", Author: "Sau Sheong"},
		{Id: 2, Content: "Bonjour Monde!", Author: "Pierre"},
		{Id: 3, Content: "Hola Mundo!", Author: "Pedro"},
		{Id: 4, Content: "Greetings Earthlings!", Author: "Sau Sheong"},
	}

	writer := csv.NewWriter(csvfile)
	for _, post := range allPosts {
		line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	//确保缓冲区数据都写入文件
	writer.Flush()

	file, err := os.Open("posts.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//设置负值，record缺少字段时，读取进程不会中断
	//正数表示每条记录读取的字段数量，世界读取少于该值会报错
	//0，表示第一条记录的字段数量用作FieldsPerRecord值
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var posts []Post
	for _, item := range record {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		post := Post{
			Id:      int(id),
			Content: item[1],
			Author:  item[2],
		}
		posts = append(posts, post)
	}
	fmt.Println(posts[0].Id)
	fmt.Println(posts[0].Content)
	fmt.Println(posts[0].Author)
}
