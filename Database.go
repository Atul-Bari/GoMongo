package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoConn struct {
	client *mongo.Client
}

func NewMongoConnection(url string, ctx context.Context) *MongoConn {

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	//"mongodb+srv://ab:ab1234@cluster0.e4s0f.mongodb.net/myFirstDatabase?retryWrites=true&w=majority",
	// "mongodb://localhost:27017",
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal("Ping error: ", err)
	}

	return &MongoConn{client: client}
}

func (conn *MongoConn) Insert(data map[string]interface{}) error {
	bsonStat := bson.M{}
	for k, v := range data {
		bsonStat[k] = v
	}
	collection := conn.client.Database("collage").Collection("student")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, bsonStat)
	if err != nil {
		return err
	}
	return nil

}

func (conn *MongoConn) InsertMillion() error {
	bsonStat := bson.M{}
	data := make(map[string]interface{})

	div := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M"}
	name := []string{"Jon", "Doe", "leo", "sam", "paul", "AB"}

	for i := 0; i < 100; i++ {
		data["_id"] = i
		data["RollNo"] = 3144 + i
		data["Marks"] = rand.Intn(100)
		data["Class"] = div[rand.Intn(13)]
		data["Name"] = name[rand.Intn(6)]

		// fmt.Println(data)
		for k, v := range data {
			bsonStat[k] = v
		}
		collection := conn.client.Database("collage").Collection("s2")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_, err := collection.InsertOne(ctx, bsonStat)
		if err != nil {
			return err
		}
	}
	return nil

}

func (conn *MongoConn) Update() {

}

func (conn *MongoConn) Find(name string) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	collection := conn.client.Database("collage").Collection("s2")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"Name": name}
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, nil
	}
	return result, nil
}
