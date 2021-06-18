package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func deleteOne(ctx context.Context, db *mongo.Database) {
	podcastsCollection := db.Collection("podcasts")
	result, err := podcastsCollection.DeleteOne(ctx, bson.M{"title": "The Nic Raboy Show"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("DeleteOne removed %v document(s)\n", result.DeletedCount)
}
func deleteMany(ctx context.Context, db *mongo.Database) {

	episodesCollection := db.Collection("episodes")
	result, err := episodesCollection.DeleteMany(ctx, bson.M{"duration": 25})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("DeleteMany removed %v document(s)\n", result.DeletedCount)
}

func dropCollection(ctx context.Context, db *mongo.Database) {
	podcastsCollection := db.Collection("podcasts")
	if err := podcastsCollection.Drop(ctx); err != nil {
		log.Fatal(err)
	}
}
