package main

import (
	"fmt"
	"strconv"
)

func getBinaryRepresentation(n int) []int {
	ans := []int{}
	for i := 0; i <= n; i++ {
		binaryStr := strconv.FormatInt(int64(i), 2) //convert to binary representation
		binaryInt, _ := strconv.Atoi(binaryStr)
		ans = append(ans, binaryInt)
	}
	return ans
}

func main() {
	// Contoh penggunaan
	ans := getBinaryRepresentation(2)
	fmt.Println(ans) // Output: [0 1 10]
	// Contoh penggunaan
	ans2 := getBinaryRepresentation(3)
	fmt.Println(ans2) // Output: [0 1 10 11]
}
