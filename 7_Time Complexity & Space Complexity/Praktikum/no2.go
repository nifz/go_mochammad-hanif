package main

import "fmt"

func pow(x, n int) int {
	// Kasus dasar
	if n == 0 {
		return 1
	}

	// Jika pangkatnya ganjil, ubah menjadi pangkat genap dengan mengalikan dengan x
	if n%2 == 1 {
		return x * pow(x, n-1)
	}

	// Jika pangkatnya genap, bagi pangkat menjadi dua dan cari hasil pangkat dari setengah pangkat
	halfPow := pow(x, n/2)
	return halfPow * halfPow
}

func main() {
	fmt.Println(pow(2, 3))  // 8
	fmt.Println(pow(5, 3))  // 125
	fmt.Println(pow(10, 2)) // 100
	fmt.Println(pow(2, 5))  // 32
	fmt.Println(pow(7, 3))  // 343
}
