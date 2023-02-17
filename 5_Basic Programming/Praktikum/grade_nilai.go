package main

import "fmt"

func main() {
	var nilai int
	fmt.Print("Masukkan sebuah nilai: ")
	fmt.Scanln(&nilai)

	if nilai < 0 || nilai > 100 {
		fmt.Println("Nilai Invalid")
	} else if nilai >= 80 {
		fmt.Println("Nilai:", nilai, "= Grade: A")
	} else if nilai >= 65 {
		fmt.Println("Nilai:", nilai, "= Grade: B")
	} else if nilai >= 50 {
		fmt.Println("Nilai:", nilai, "= Grade: C")
	} else if nilai >= 35 {
		fmt.Println("Nilai:", nilai, "= Grade: D")
	} else {
		fmt.Println("Nilai:", nilai, "= Grade: E")
	}
}
