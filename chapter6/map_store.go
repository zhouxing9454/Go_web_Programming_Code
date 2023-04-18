package main

import "fmt"

type Post struct {
	Id      int
	Content string
	Author  string
}

var (
	PostById     = make(map[int]*Post)
	PostByAuthor = make(map[string][]*Post)
)

func store(post *Post) {
	PostById[post.Id] = post
	PostByAuthor[post.Author] = append(PostByAuthor[post.Author], post)
}

func main() {
	post1 := &Post{Id: 1, Content: "Hello World!", Author: "Sau Sheong"}
	post2 := &Post{Id: 2, Content: "Bonjour Monde!", Author: "Pierre"}
	post3 := &Post{Id: 3, Content: "Hola Mundo!", Author: "Pedro"}
	post4 := &Post{Id: 4, Content: "Greetings Earthlings!", Author: "Sau Sheong"}

	store(post1)
	store(post2)
	store(post3)
	store(post4)

	fmt.Println(PostById[1])
	fmt.Println(PostById[2])

	for _, post := range PostByAuthor["Sau Sheong"] {
		fmt.Println(post)
	}

	for _, post := range PostByAuthor["Pierre"] {
		fmt.Println(post)
	}
}
