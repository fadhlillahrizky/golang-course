package main

import "fmt"

var ar string
var ka int64

var (
	satu, dua, tiga string
	empat           int
)

func main() {
	var de float64
	var my bool

	ar = "1"
	ka = 2
	de = 3.4
	my = true

	kata := "apa"
	angka := 123
	koma := 9.5
	yakin := true
	//sampah := "sampah"

	fmt.Printf("kata \t %v \t => %T \n", kata, kata)
	fmt.Printf("angka \t %v \t => %T \n", angka, angka)
	fmt.Printf("koma \t %v \t => %T \n", koma, koma)
	fmt.Printf("yakin \t %v \t => %T \n", yakin, yakin)

	cetakArka()
	fmt.Print(de, my)
}

func cetakArka() {
	fmt.Print(ar, ka)
}
