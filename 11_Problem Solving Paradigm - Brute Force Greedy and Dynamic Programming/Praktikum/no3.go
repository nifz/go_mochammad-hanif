package main

import "fmt"

func intToRoman(num int) string {
	var values = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	var symbols = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var result string

	for i := 0; i < len(values); i++ {
		for num >= values[i] {
			num -= values[i]
			result += symbols[i]
		}
	}
	return result
}

func main() {
	fmt.Println(intToRoman(4))    // IV
	fmt.Println(intToRoman(9))    // IX
	fmt.Println(intToRoman(23))   // XXIII
	fmt.Println(intToRoman(2021)) // MMXXI
	fmt.Println(intToRoman(1646)) // MDCXLVI
}
