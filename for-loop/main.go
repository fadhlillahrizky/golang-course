package main

import "fmt"

func main() {
	for i := 0; i <= 7; i++ {
		fmt.Println("i => ", i)
	}

	x := 1
	for x != 10 {
		fmt.Println("aku x ke ", x)

		x++
	}
	// untuk kemamanan code ini di comment
	// for {
	// }
}
