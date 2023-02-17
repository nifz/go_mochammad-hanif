package main

import "fmt"

func main() {
	var number int
	fmt.Print("Masukkan sebuah bilangan bulat: ")
	fmt.Scanln(&number)

	if number%7 == 0 {
		fmt.Println(number, "adalah kelipatan 7")
	} else {
		fmt.Println(number, "bukan kelipatan 7")
	}
}
