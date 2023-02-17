package main

import "fmt"

func main() {
	var number int
	fmt.Print("Masukkan sebuah bilangan bulat: ")
	fmt.Scanln(&number)

	if number%2 == 0 {
		fmt.Println(number, "adalah bilangan genap")
	} else {
		fmt.Println(number, "adalah bilangan ganjil")
	}
}
