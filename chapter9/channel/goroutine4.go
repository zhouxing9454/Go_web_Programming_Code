package main

import "fmt"

func callerA2(c chan string) {
	c <- "hello world!"
	close(c)
}

func callerB2(c chan string) {
	c <- "hola mundo!"
	close(c)
}

func main() {
	a, b := make(chan string), make(chan string)
	go callerA2(a)
	go callerB2(b)
	var msg string
	ok1, ok2 := true, true
	for ok1 || ok2 {
		select {
		case msg, ok1 = <-a:
			if ok1 {
				fmt.Printf("%s from A\n", msg)
			}
		case msg, ok2 = <-b:
			if ok2 {
				fmt.Printf("%s from B\n", msg)
			}
		}
	}
}
