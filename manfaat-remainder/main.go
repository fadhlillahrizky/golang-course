package main

import (
	"fmt"
)

func main() {
	angkaPertama := 10
	angkaKedua := 5

	fmt.Println("A ", angkaPertama%1)
	fmt.Println("B ", angkaPertama%2)

	fmt.Println("C ", angkaKedua/3)

	fmt.Println("D ", angkaKedua%3)
	fmt.Println("E ", angkaKedua%4)

	fmt.Println("F ", angkaPertama%3)

	fmt.Println(angkaPertama, " ganjil ? ", angkaPertama%2)
	fmt.Println(angkaKedua, " ganjil ? ", angkaKedua%2)
}
