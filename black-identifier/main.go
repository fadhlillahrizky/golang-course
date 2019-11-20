package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	_ "time"
)

func main() {
	res, _ := http.Get("https://arkademy.com/")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}
