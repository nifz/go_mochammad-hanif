/*
	Setiap penamaan struct dan method diawal nama diharuskan huruf besar, agar bisa diakses oleh package lain. Jika nama struct atau function ditulis dengan huruf kecil, maka akan dianggap sebagai private dan tidak bisa diakses oleh package lain, kecuali jika package tersebut adalah package yang sama.
	Dengan menggunakan huruf besar pada penamaan struct dan function, maka package lain dapat mengakses dan menggunakan struktur data atau metode yang disediakan oleh package tersebut untuk membangun program yang lebih kompleks. Dengan cara ini, komponen program dapat saling berinteraksi dan saling memanfaatkan satu sama lain, sehingga dapat memudahkan dalam pengembangan dan pemeliharaan kode.

	Ada beberapa kekurangan dalam kode tersebut di antaranya:
	1. Struktur data user tidak konsisten dalam tipe data username dan password yang semestinya berupa string
	2. Struktur data userservice tidak memiliki dokumentasi terkait tujuan dan fungsinya
	3. Metode getallusers tidak memiliki dokumentasi terkait tujuan dan fungsinya
	4. Metode getuserbyid tidak memiliki dokumentasi terkait tujuan dan fungsinya
	Kekurangan-kekurangan tersebut dapat menyebabkan masalah dalam pemeliharaan dan pengembangan kode, serta dapat membuat kode menjadi sulit dipahami oleh programmer lainnya.
*/

package main

import "fmt"

// Struktur data user yang tidak konsisten dalam tipe data username dan password yang semestinya berupa string
type user struct {
	id       int
	username int
	password int
}

// Struktur data userservice tidak memiliki dokumentasi terkait tujuan dan fungsinya
type userservice struct {
	t []user
}

// Metode getallusers tidak memiliki dokumentasi terkait tujuan dan fungsinya
func (u userservice) getallusers() []user {
	return u.t
}

// Metode getuserbyid tidak memiliki dokumentasi terkait tujuan dan fungsinya
func (u userservice) getuserbyid(id int) user {
	for _, r := range u.t {
		if id == r.id {
			return r
		}
	}

	return user{}
}

// Menambahkan function main agar aplikasi bisa jalankan
func main() {
	// buat instance baru dari struct layanan pengguna
	svc := userservice{t: []user{{id: 1, username: 123, password: 123}}}

	// get all users
	users := svc.getallusers()
	fmt.Println(users)

	// get user by ID
	u := svc.getuserbyid(1)
	fmt.Println(u)
}
