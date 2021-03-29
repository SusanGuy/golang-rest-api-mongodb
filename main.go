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
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	w.Header().Set("Content-Type", "application/json")
	hash_id := mux.Vars(r)["id"]
	id, err := primitive.ObjectIDFromHex(hash_id)
	if err != nil {
		helper.GetError(err, w)
		return
	}
	var book models.Book
	filter := bson.M{"_id": id}
	json.NewDecoder(r.Body).Decode(&book)
	update := bson.D{
		{"$set", bson.D{
			{"isbn", book.Isbn},
			{"title", book.Title},
			{"author", bson.D{
				{"first_name", book.Author.FirstName},
				{"last_name", book.Author.LastName},
			}},
		}},
	}
	error := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&book)
	fmt.Println(book.Title, book.Isbn)
	if error != nil {
		helper.GetError(error, w)
		return
	}

	book.ID = id
	json.NewEncoder(w).Encode(book)

}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	hash_id := mux.Vars(r)["id"]
	id, err := primitive.ObjectIDFromHex(hash_id)
	if err != nil {
		helper.GetError(err, w)
		return
	}
	filter := bson.M{"_id": id}
	deleteRequest, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		helper.GetError(err, w)
		return
	}
	json.NewEncoder(w).Encode(deleteRequest)
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
