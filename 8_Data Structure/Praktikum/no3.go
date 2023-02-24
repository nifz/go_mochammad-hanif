package main

import "fmt"

func diagonalDifference(arr [][]int32) int32 {
	var leftDiagonal, rightDiagonal rune

	for i := 0; i < len(arr); i++ {
		leftDiagonal += arr[i][i]
		rightDiagonal += arr[i][len(arr)-i-1]
	}

	return abs(leftDiagonal - rightDiagonal)
}

func abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	arr := [][]int32{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9}}
	fmt.Println(diagonalDifference(arr)) // Output: 2
}
