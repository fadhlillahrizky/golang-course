package main

import (
	"fmt"
)

func main() {
	tahun := 2018

	fmt.Println("tahun ", tahun)
	fmt.Println("tahun ", &tahun) // mendapatkan memory addres

	var sekarang *int = &tahun

	fmt.Println("sekarang ", sekarang)

	fmt.Println("sekarang ", *sekarang)

	*sekarang = 2019

	fmt.Println("sekarang ", sekarang)
	fmt.Println("sekarang ", *sekarang)

}
