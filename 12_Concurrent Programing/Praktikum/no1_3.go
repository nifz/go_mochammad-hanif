package main

import "fmt"

func bil(ch chan int, done chan bool) {
	for i := 1; i <= 10; i++ {
		ch <- i * 3
	}
	done <- true
}

func main() {
	ch := make(chan int, 3) // buat buffered channel
	done := make(chan bool)

	// Spawn goroutine
	go bil(ch, done)

	// Receive values from channel
	for {
		select {
		case val, ok := <-ch:
			if ok {
				fmt.Println(val)
			} else {
				ch = nil // channel selesai
			}
		case <-done:
			return
		}
	}
}
