package goroutine

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func printNumbers1() {
	for i := 0; i < 10; i++ {
		//fmt.Printf("%d ", i)
	}
}

func printLetters1() {
	for i := 'A'; i < 'A'+10; i++ {
		//fmt.Printf("%c ", i)
	}
}

func print1() {
	printNumbers1()
	printLetters1()
}

func print2() {
	printNumbers2(&wg)
	printLetters2(&wg)
}

func goPrint1() {
	go printNumbers1()
	go printLetters1()
}

func goPrint2() {
	go printNumbers2(&wg)
	go printLetters2(&wg)
}

func printNumbers2(group *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Microsecond) // 微秒
		fmt.Printf("%d ", i)
	}
	group.Done()
}

func printLetters2(group *sync.WaitGroup) {
	for i := 'A'; i < 'A'+10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%c ", i)
	}
	group.Done()
}

func main() {
	wg.Add(2)
	go printNumbers2(&wg)
	go printLetters2(&wg)
	wg.Wait()
}
