package main

import (
	"fmt"
	"strconv"
)

func main() {
	var text string
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			text += "Fizz"
		}
		if i%5 == 0 {
			text += "Buzz"
		}
		if i%3 != 0 && i%5 != 0 {
			text += strconv.Itoa(i)
		}
		text += "\n"
	}
	fmt.Println(text)
}
