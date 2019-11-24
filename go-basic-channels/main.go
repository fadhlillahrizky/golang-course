package main

import (
	"fmt"
)

func main() {
	fmt.Println("Go Channels Tutorial")

	values := make(chan int)
	defer close(values)
}