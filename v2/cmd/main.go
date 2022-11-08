package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/funkynunu/go_rest_api/v2/pkg/db"
	"github.com/funkynunu/go_rest_api/v2/pkg/handlers"
	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome my 1st Go api page with a database!") // displayed in browser
	fmt.Println("What what")                                     // display in terminal
}

func handleRequests(DB *sql.DB) {
	h := handlers.New(DB)
	// create a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", h.GetAllArticles).Methods(http.MethodGet)
	myRouter.HandleFunc("/articles/{id}", h.GetArticle).Methods(http.MethodGet)
	myRouter.HandleFunc("/articles", h.AddArticle).Methods(http.MethodPost)
	myRouter.HandleFunc("/articles/{id}", h.UpdateArticle).Methods(http.MethodPut)
	myRouter.HandleFunc("/articles/{id}", h.DeleteArticle).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	DB := db.Connect()
	db.CreateTable(DB)
	handleRequests(DB)
	db.CloseConnection(DB)
}
