package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/alscaldeira/twitter/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UserGet(username string) model.User {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	defer cancel()
	defer client.Disconnect(ctx)

	if err != nil {
		panic(err)
	}

	twitterDb := client.Database("twitter")
	userCollection := twitterDb.Collection("user")

	cursor, err := userCollection.Find(ctx, bson.D{{Key: "username", Value: username}})

	if err != nil {
		fmt.Println(err.Error())
	}

	defer cursor.Close(ctx)

	if err != nil {
		panic(err)
	}

	var result model.User
	for cursor.Next(ctx) {
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
	}
	return result
}

func UserPost(username string, password string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	defer cancel()
	defer client.Disconnect(ctx)

	if err != nil {
		panic(err)
	}

	twitterDb := client.Database("twitter")
	userCollection := twitterDb.Collection("user")

	_, err = userCollection.InsertOne(ctx, bson.D{
		{Key: "username", Value: username},
		{Key: "password", Value: password},
	})

	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}
}
