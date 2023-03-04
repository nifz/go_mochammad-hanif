package main

import "fmt"

func MaximumBuyProduct(money int, productPrice []int) {
	for i := 0; i < len(productPrice); i++ {
		for j := i + 1; j < len(productPrice); j++ {
			if productPrice[i] > productPrice[j] {
				temp := productPrice[i]
				productPrice[i] = productPrice[j]
				productPrice[j] = temp
			}
		}
	}

	count := 0
	for i := 0; i < len(productPrice); i++ {
		if money >= productPrice[i] {
			money -= productPrice[i]
			count++
		} else {
			break
		}
	}

	fmt.Println(count)
}

func main() {
	MaximumBuyProduct(50000, []int{25000, 25000, 10000, 14000})      // 3
	MaximumBuyProduct(30000, []int{15000, 10000, 12000, 5000, 3000}) // 4
	MaximumBuyProduct(10000, []int{2000, 3000, 1000, 2000, 10000})   // 4
	MaximumBuyProduct(4000, []int{7500, 3000, 2500, 2000})           // 1
	MaximumBuyProduct(0, []int{10000, 30000})                        // 0
}
