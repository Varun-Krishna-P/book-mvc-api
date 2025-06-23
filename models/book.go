package models

import (
    "context"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

type Book struct {
    ID     interface{} `bson:"_id,omitempty"`
    Title  string      `bson:"title"`
    Author string      `bson:"author"`
}

// Helper to insert a book
func (b *Book) Insert(collection *mongo.Collection) error {
    b.ID = primitive.NewObjectID()
    _, err := collection.InsertOne(context.Background(), b)
    return err
}

// Helper to get all books
func GetAllBooks(collection *mongo.Collection) ([]Book, error) {
    cursor, err := collection.Find(context.Background(), bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.Background())

    var books []Book
    for cursor.Next(context.Background()) {
        var book Book
        cursor.Decode(&book)
        books = append(books, book)
    }
    return books, nil
}

// Find a book by ID
func GetBookByID(collection *mongo.Collection, id primitive.ObjectID) (*Book, error) {
    var book Book
    err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&book)
    if err != nil {
        return nil, err
    }
    return &book, nil
}

// Update a book by ID
func UpdateBookByID(collection *mongo.Collection, id primitive.ObjectID, book *Book) error {
    update := bson.M{"$set": book}
    _, err := collection.UpdateOne(context.Background(), bson.M{"_id": id}, update)
    return err
}

// Delete a book by ID
func DeleteBookByID(collection *mongo.Collection, id primitive.ObjectID) error {
    _, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
    return err
}