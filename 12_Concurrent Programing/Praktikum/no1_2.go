package main

import (
	"fmt"
)

func bil(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i * 3
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go bil(ch)
	for i := range ch {
		fmt.Println(i)
	}
}
