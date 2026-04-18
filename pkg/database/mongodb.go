package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB(uri string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("MONGODB CONNECTING ERROR: ", err) // Nếu hiện dòng này, hãy bật MongoDB Compass lên kiểm tra
	}

	log.Println("Successfully connected to MongoDB")
	return client, nil
}
