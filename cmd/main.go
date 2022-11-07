package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/funkynunu/go_rest_api/pkg/handlers"
	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome my 1st Go api page!") // displayed in browser
	fmt.Println("What what")                     // display in terminal
}

func handleRequests() {
	// create a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", handlers.GetAllArticles).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	handleRequests()
}