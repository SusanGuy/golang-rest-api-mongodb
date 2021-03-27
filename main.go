package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getBooks(w http.ResponseWriter, r *http.Request) {

}

func getBook(w http.ResponseWriter, r *http.Request) {

}

func createBook(w http.ResponseWriter, r *http.Request) {

}

func updateBook(w http.ResponseWriter, r *http.Request) {

}

func deleteBook(w http.ResponseWriter, r *http.Request) {

}

//collection := helper.ConnectDB()

func main() {

	r := mux.NewRouter()
	r.HandleFunc("api/books/", getBooks).Methods("GET")
	r.HandleFunc("api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))
}
