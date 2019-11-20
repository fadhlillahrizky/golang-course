package main

import (
	"fmt"
)

const (
	A = iota
	B
	C
	D
	E
	F
)

const nama string = "Arkademy"

const (
	panjang     = 2.9
	lebar   int = 20
	tinggi      = 0
)

func main() {
	const gravity int = 10

	fmt.Printf("panjang: %v tipe nya: %T \n", panjang, panjang)
	fmt.Println("Nama: ", nama)
	fmt.Println("gravitasi: ", gravity)
	fmt.Println("iota: ", F)
}
