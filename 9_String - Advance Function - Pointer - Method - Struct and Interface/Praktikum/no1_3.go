package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {
	for i := len(a); i > 0; i-- {
		for j := 0; j <= len(a)-i; j++ {
			if strings.Contains(b, a[j:j+i]) {
				return a[j : j+i]
			}
		}
	}
	return ""
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))     // AKA
	fmt.Println(Compare("KANGOORO", "KANG"))  // KANG
	fmt.Println(Compare("KI", "KIJANG"))      // KI
	fmt.Println(Compare("KUPU-KUPU", "KUPU")) // KUPU
	fmt.Println(Compare("ILALANG", "ILA"))    // ILA
}
