package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
)

var database *MongoConn

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	database = NewMongoConnection(os.Getenv("MONGO_URI"), ctx)
	defer database.client.Disconnect(ctx)

	dbs, _ := database.client.ListDatabaseNames(ctx, bson.M{})
	log.Println(dbs)

	// fmt.Println(database.InsertMillion())
	data, _ := database.Find("Jon")
	log.Println(data)
}
