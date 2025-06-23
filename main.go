package main

import (
    "context"
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
	"os"
    "github.com/joho/godotenv"

	"book-mcv-api/config"
    "book-mcv-api/controllers"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }
    // Connect to MongoDB
    client, err := config.ConnectMongoDB(os.Getenv("MONGODB_URI"))
    if err != nil {
        fmt.Println("MongoDB connection error:", err)
        return
    }
    defer client.Disconnect(context.Background())
	// Get the books collection from your database (replace "yourdbname" and "books" as needed)
    books_collection := client.Database("learning_db").Collection("books")
    controllers.SetCollection(books_collection)

    router := mux.NewRouter()
    router.HandleFunc("/books", controllers.CreateBook).Methods("POST")
    router.HandleFunc("/books", controllers.GetBooks).Methods("GET")
    router.HandleFunc("/books/{id}", controllers.GetBook).Methods("GET")
    router.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")
    router.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")

    fmt.Println("Server running at http://localhost:8000")
    log.Fatal(http.ListenAndServe(":8000", router))
}