package controllers

import (
	"context"
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"

    "book-mcv-api/models"
)


var Collection *mongo.Collection

// Injects the MongoDB collection (called from main)
func SetCollection(c *mongo.Collection) {
    Collection = c
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
    var book models.Book
    json.NewDecoder(r.Body).Decode(&book)
    err := book.Insert(Collection)
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    json.NewEncoder(w).Encode(book)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
    books, err := models.GetAllBooks(Collection)
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    json.NewEncoder(w).Encode(books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := primitive.ObjectIDFromHex(params["id"])
    book, err := models.GetBookByID(Collection, id)
    if err != nil {
        http.Error(w, "Book not found", 404)
        return
    }
    json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := primitive.ObjectIDFromHex(params["id"])
    var book models.Book
    json.NewDecoder(r.Body).Decode(&book)
    err := models.UpdateBookByID(Collection, id, &book)
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    book.ID = id
    json.NewEncoder(w).Encode(book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := primitive.ObjectIDFromHex(params["id"])
    err := models.DeleteBookByID(Collection, id)
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    json.NewEncoder(w).Encode(bson.M{"message": "Book deleted"})
}