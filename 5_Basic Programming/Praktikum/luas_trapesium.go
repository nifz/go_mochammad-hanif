package main

import "fmt"

func main() {
	var a, b, t float64

	fmt.Println("Menghitung Luas Trapesium")
	fmt.Print("Masukkan panjang sisi sejajar pertama: ")
	fmt.Scanln(&a)
	fmt.Print("Masukkan panjang sisi sejajar kedua: ")
	fmt.Scanln(&b)
	fmt.Print("Masukkan tinggi trapesium: ")
	fmt.Scanln(&t)

	luas := 0.5 * (a + b) * t

	fmt.Println("Luas trapesium adalah", luas)
}
