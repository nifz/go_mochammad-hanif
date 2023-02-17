package main

import (
	"fmt"
	"math"
)

func main() {
	var number int
	fmt.Print("Masukkan sebuah bilangan bulat: ")
	fmt.Scanln(&number)

	if isPrime(number) {
		fmt.Println(number, "adalah bilangan prima")
	} else {
		fmt.Println(number, "bukan bilangan prima")
	}
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	// Mengecek apakah n habis dibagi oleh 2
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	// Mengecek apakah n habis dibagi oleh bilangan ganjil lainnya yang lebih kecil daripada akar n
	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}
