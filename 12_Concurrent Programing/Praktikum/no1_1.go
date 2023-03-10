package main

import (
	"fmt"
	"time"
)

func bil(x int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", i, x, i*x)
		time.Sleep(3 * time.Second)
	}
}

func main() {
	go bil(5)
	fmt.Println("Press enter to exit")
	fmt.Scanln()
}
