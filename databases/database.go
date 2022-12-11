package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectDB() *mongo.Client {
	Mongo_URL := "mongodb+srv://bmajor414:Najiri41214.@cluster0.gtixszz.mongodb.net/?retryWrites=true&w=majority"
	client, err := mongo.NewClient(options.Client().ApplyURI(Mongo_URL))
	
	if err != nil {
	log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	err = client.Connect(ctx)
	defer cancel()

	if err != nil {
		log.Fatal(err)
	}

	
	
	err = client.Ping(ctx, readpref.Primary())
		if err != nil {
			log.Fatal(err)
		}
		
		databases, err := client.ListDatabaseNames(ctx, bson.M{})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(databases)


fmt.Println("MongoDB connection established")
	return client
}