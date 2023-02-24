package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	// gabungkan kedua array
	result := append(arrayA, arrayB...)

	// buat map untuk menyimpan setiap elemen yang sudah ditemukan
	uniqueMap := make(map[string]bool)
	uniqueSlice := []string{}

	// periksa setiap elemen pada array hasil gabungan
	// jika elemen belum ditemukan, tambahkan ke uniqueSlice dan tambahkan ke uniqueMap
	for _, element := range result {
		if !uniqueMap[element] {
			uniqueSlice = append(uniqueSlice, element)
			uniqueMap[element] = true
		}
	}

	return uniqueSlice
}

func main() {
	// Test cases
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	// ["king", "devil jin", "akuma", "eddie", "steve", "geese"]
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	// ["sergei", "jin", "steve", "bryan"]
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	// ["alisa", "yoshimitsu", "devil jin", "law"]
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	// ["devil jin", "sergei"]
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	// ["hwoarang"]
	fmt.Println(ArrayMerge([]string{}, []string{}))
	// []
}
