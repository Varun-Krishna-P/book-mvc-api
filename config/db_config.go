package config

import (
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "log"
    "time"
)

func ConnectMongoDB(uri string) (*mongo.Client, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
	clientOpts := options.Client().ApplyURI(uri)
	// Set client options
    client, err := mongo.Connect(ctx, clientOpts)
    if err != nil {
        return nil, err
    }
    if err := client.Ping(ctx, nil); err != nil {
        return nil, err
    }
    log.Println("Connected to MongoDB!")
    return client, nil
}