package main

import "fmt"

func primeX(number int) int {
	count, n := 0, 2
	for {
		isPrime := true
		for i := 2; i <= n/2; i++ {
			if n%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			count++
			if count == number {
				return n
			}
		}
		n++
	}
}

func main() {
	fmt.Println(primeX(1))  // 2
	fmt.Println(primeX(5))  // 11
	fmt.Println(primeX(8))  // 19
	fmt.Println(primeX(9))  // 23
	fmt.Println(primeX(10)) // 29
}
