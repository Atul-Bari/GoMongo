package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

var database *MongoConn

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	database = NewMongoConnection("mongodb://localhost:27017", ctx)
	defer database.client.Disconnect(ctx)

	dbs, _ := database.client.ListDatabaseNames(ctx, bson.M{})
	fmt.Println(dbs)

	// fmt.Println(database.InsertMillion())
	data, _ := database.Find("Jon")
	fmt.Println(data)
}
