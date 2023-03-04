package main

import "fmt"

func Frog(jumps []int) int {
	n := len(jumps)
	frog := make([]int, n)

	frog[0] = 0
	frog[1] = jumps[1] - jumps[0]
	if frog[1] < 0 {
		frog[1] = -frog[1]
	}

	for i := 2; i < n; i++ {
		diff1 := jumps[i] - jumps[i-1]
		if diff1 < 0 {
			diff1 = -diff1
		}
		diff2 := jumps[i] - jumps[i-2]
		if diff2 < 0 {
			diff2 = -diff2
		}
		if frog[i-1]+diff1 < frog[i-2]+diff2 {
			frog[i] = frog[i-1] + diff1
		} else {
			frog[i] = frog[i-2] + diff2
		}
	}

	return frog[n-1]
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))         // 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40
}
