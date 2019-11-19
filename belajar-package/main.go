package main

import (
	"belajar-package/greeting"
	selamat "belajar-package/greeting"

	s "belajar-package/greeting"
)

func main() {
	salam()
}

func salam() {
	greeting.Sore()

	selamat.Pagi()

	s.Siang()
	selamat.Malam()

	selamat.Tidur()
}
