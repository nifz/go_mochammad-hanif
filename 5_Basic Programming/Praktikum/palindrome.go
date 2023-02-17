package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Print("Apakah Palindrome?\nmasukkan kata: ")
	fmt.Scanln(&input)
	fmt.Print("captureed: ")

	if isPalindrome(input) {
		fmt.Println(input, "\nPalindrome")
	} else {
		fmt.Println(input, "\nBukan Palindrome")
	}
}

func isPalindrome(input string) bool {
	input = strings.ToLower(input)             // Ubah ke huruf kecil semua
	input = strings.ReplaceAll(input, " ", "") // Hapus spasi

	// Iterasi setengah kata dan periksa apakah karakter di kedua ujung sama
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}

	return true
}
