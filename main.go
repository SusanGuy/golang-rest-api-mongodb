package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/SusanGuy/golang-rest-api-mongodb/helper"
	"github.com/SusanGuy/golang-rest-api-mongodb/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

var collection = helper.ConnectDB()

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var books []models.Book = []models.Book{}
	cur, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		helper.GetError(err, w)
		return
	}
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var book models.Book
		err := cur.Decode(&book)
		if err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)
	result, err := collection.InsertOne(context.TODO(), book)
	if err != nil {
		helper.GetError(err, w)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func updateBook(w http.ResponseWriter, r *http.Request) {

}

func deleteBook(w http.ResponseWriter, r *http.Request) {

}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	fmt.Println("Server Started on Port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
