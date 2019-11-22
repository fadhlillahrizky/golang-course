package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Tilte"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Test Title", Desc: "Test Descriptin", Content: "Hello World"},
	}
	fmt.Println("Endpoin Hit : all Article Endpoin")
	json.NewEncoder(w).Encode(articles)
}

func testPostArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test Post endpoint worked")

}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")

}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", testPostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8881", myRouter))
}

func main() {
	handleRequests()
}
