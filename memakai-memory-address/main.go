package main

import (
	"fmt"
)

const detailDalamMenit int64 = 60

func main() {
	var menit int64

	fmt.Print("masukkan menit: ")
	fmt.Scan(&menit)

	detik := menit * detailDalamMenit

	fmt.Println(menit, " menit adalah ", detik, "detik")
}
