package main

import (
	"fmt"
	"sync"
)

func countFreq(text string, ch chan<- map[rune]int, wg *sync.WaitGroup) {
	freq := make(map[rune]int)
	defer func() {
		ch <- freq
		wg.Done()
	}()
	for _, char := range text {
		freq[char]++
	}
}

func main() {
	text := "Mochammad Hanif"
	chans := make(chan map[rune]int)
	var wg sync.WaitGroup
	wg.Add(1)
	go countFreq(text, chans, &wg)

	go func() {
		wg.Wait()
		close(chans)
	}()

	freq := make(map[rune]int)
	for f := range chans {
		for k, v := range f {
			freq[k] += v
		}
	}

	for k, v := range freq {
		fmt.Printf("%c: %d\n", k, v)
	}
}
