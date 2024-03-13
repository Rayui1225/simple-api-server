package main

import (
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "log"
)

func connectDB() *mongo.Client {
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    // 连接到MongoDB
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }
    return client
}