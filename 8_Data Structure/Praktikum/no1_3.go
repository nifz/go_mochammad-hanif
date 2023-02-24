package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int {
	var freq = make(map[rune]int)
	for _, digit := range angka {
		freq[digit]++
	}

	var result []int
	for _, digit := range angka {
		if freq[digit] == 1 {
			digitInt, _ := strconv.Atoi(string(digit))
			result = append(result, digitInt)
		}
	}
	return result
}

func main() {
	fmt.Println(munculSekali("1234123"))    // [4]
	fmt.Println(munculSekali("76523752"))   // [6, 3]
	fmt.Println(munculSekali("12345"))      // [1, 2, 3, 4, 5]
	fmt.Println(munculSekali("1122334455")) // []
	fmt.Println(munculSekali("0872504"))    // [8, 7, 2, 5, 4]
}
