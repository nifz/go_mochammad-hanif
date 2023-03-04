package main

import "fmt"

type student struct {
	name       string
	nameEncode string
	score      int
}

type Cipher interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var nameEncode string
	cipherMap := map[rune]rune{'a': 'z', 'b': 'y', 'c': 'x', 'd': 'w', 'e': 'v', 'f': 'u', 'g': 't', 'h': 's', 'i': 'r', 'j': 'q', 'k': 'p', 'l': 'o', 'm': 'n', 'n': 'm', 'o': 'l', 'p': 'k', 'q': 'j', 'r': 'i', 's': 'h', 't': 'g', 'u': 'f', 'v': 'e', 'w': 'd', 'x': 'c', 'y': 'b', 'z': 'a'}
	for _, char := range s.name {
		encodedChar, ok := cipherMap[char]
		if !ok {
			// jika karakter tidak ada di cipherMap maka menggunakan karakter asli
			encodedChar = char
		}
		nameEncode += string(encodedChar)
	}
	s.nameEncode = nameEncode
	return nameEncode
}

func (s *student) Decode() string {
	var nameDecode string
	cipherMap := map[rune]rune{'a': 'z', 'b': 'y', 'c': 'x', 'd': 'w', 'e': 'v', 'f': 'u', 'g': 't', 'h': 's', 'i': 'r', 'j': 'q', 'k': 'p', 'l': 'o', 'm': 'n', 'n': 'm', 'o': 'l', 'p': 'k', 'q': 'j', 'r': 'i', 's': 'h', 't': 'g', 'u': 'f', 'v': 'e', 'w': 'd', 'x': 'c', 'y': 'b', 'z': 'a'}
	for _, char := range s.nameEncode {
		decodedChar, ok := cipherMap[char]
		if !ok {
			// jika karakter tidak ada di cipherMap maka menggunakan karakter asli
			decodedChar = char
		}
		nameDecode += string(decodedChar)
	}
	s.name = nameDecode
	return nameDecode
}

func main() {
	var menu int
	var a student
	var c Cipher = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")
	fmt.Scan(&menu)

	if menu == 1 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of student’s name " + a.name + " is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Encoded Student Name: ")
		fmt.Scan(&a.nameEncode)
		fmt.Print("\nDecode of student’s name " + a.nameEncode + " is : " + c.Decode())
	}
}
