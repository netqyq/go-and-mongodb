package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func updataOne(ctx context.Context, db *mongo.Database) {
	podcastsCollection := db.Collection("podcasts")
	id, _ := primitive.ObjectIDFromHex("60cb0bda582af2d8f1a077b0")
	result, err := podcastsCollection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.D{
			{"$set", bson.D{{"author", "Nic Raboy 1"}}},
		},
	)
	// if id is not matched, or the the content will not be changed
	// result.ModifiedCount will be 0.
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)
}

func updataMany(ctx context.Context, db *mongo.Database) {
	podcastsCollection := db.Collection("podcasts")

	result, err := podcastsCollection.UpdateMany(
		ctx,
		bson.M{"title": "The Polyglot Developer Podcast"},
		bson.D{
			{"$set", bson.D{{"author", "Nicolas Raboy dd"}}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)
}

func ReplaceOne(ctx context.Context, db *mongo.Database) {
	podcastsCollection := db.Collection("podcasts")
	result, err := podcastsCollection.ReplaceOne(
		ctx,
		bson.M{"author": "Lily"},
		bson.M{
			"title":  "The Nic Raboy Show",
			"author": "tom",
		},
	)

	if err != nil {
		log.Panicln(err)
	}

	fmt.Printf("Replaced %v Documents!\n", result.ModifiedCount)
}
