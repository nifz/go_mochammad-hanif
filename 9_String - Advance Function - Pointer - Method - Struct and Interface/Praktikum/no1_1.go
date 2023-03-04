package main

import "fmt"

type Car struct {
	Type   string
	FuelIn float64 // Bahan bakar dalam liter
}

func (c *Car) Calc() float64 {
	fuelConsumption := 7.5 // Rata-rata konsumsi bahan bakar per 100 km
	return (c.FuelIn * 100) / fuelConsumption
}

func main() {
	myCar := Car{Type: "sedan", FuelIn: 1.5} // Buat instance dari struct Car dengan tipe sedan dan 20 liter bahan bakar
	distance := myCar.Calc()                 // Hitung jarak yang bisa ditempuh
	fmt.Printf("Mobil %s dengan %v L bahan bakar bisa menempuh jarak sekitar %v km\n", myCar.Type, myCar.FuelIn, distance)
}
