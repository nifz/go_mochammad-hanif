/*
	Perubahan yang dilakukan adalah sebagai berikut:
	Dalam penamaan struct diubah menjadi huruf besar untuk memberikan jelasnya bahwa itu adalah tipe data yang berbeda dengan variabel biasa.
	Penamaan variabel diubah dari totalroda menjadi TotalRoda dan kecepatanperjam menjadi KecepatanPerJam untuk memenuhi konvensi penamaan Go yang lebih umum dan mudah dipahami.
	Penamaan function berjalan diubah menjadi Berjalan dan tambahkecepatan diubah menjadi TambahKecepatan karena di golang menggunakan konvensi CamelCase untuk penamaan function.
	Dalam function TambahKecepatan, m.kecepatanperjam + kecepatanbaru diubah menjadi m.KecepatanPerJam += kecepatanBaru agar lebih mudah dipahami dan memenuhi konvensi penamaan variabel yang baik.
	Variabel mobilcepat diubah menjadi mobilCepat dan mobillamban diubah menjadi mobilLambat untuk memenuhi konvensi penamaan variabel yang baik.
	Beberapa spasi tambahan ditambahkan untuk memperbaiki format dan membuat kode lebih mudah dibaca.
*/

package main

import "fmt"

type Kendaraan struct {
	TotalRoda       int
	KecepatanPerJam int
}

type Mobil struct {
	Kendaraan
}

func (m *Mobil) Berjalan() {
	m.TambahKecepatan(10)
}

func (m *Mobil) TambahKecepatan(kecepatanBaru int) {
	m.KecepatanPerJam += kecepatanBaru
}

func main() {
	mobilCepat := Mobil{}
	mobilCepat.Berjalan()
	mobilCepat.Berjalan()
	mobilCepat.Berjalan()

	mobilLambat := Mobil{}
	mobilLambat.Berjalan()

	// Menambahkan format print untuk melihat output dari function yang di jalankan
	fmt.Printf("Kecepatan mobil cepat: %d km/jam\n", mobilCepat.KecepatanPerJam)
	fmt.Printf("Kecepatan mobil lambat: %d km/jam\n", mobilLambat.KecepatanPerJam)
}
