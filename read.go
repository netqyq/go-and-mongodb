package main

// https://www.mongodb.com/blog/post/quick-start-golang--mongodb--how-to-read-documents
import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Reading All Documents from a Collection
func readAllDocs(ctx context.Context, client *mongo.Client) {
	quickstartDatabase := client.Database("quickstart")
	episodesCollection := quickstartDatabase.Collection("episodes")

	cursor, err := episodesCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var episodes []bson.M
	if err = cursor.All(ctx, &episodes); err != nil {
		log.Fatal(err)
	}
	fmt.Println(episodes)
}

func readAllDocsIterate(ctx context.Context, client *mongo.Client) {
	quickstartDatabase := client.Database("quickstart")
	episodesCollection := quickstartDatabase.Collection("episodes")

	cursor, err := episodesCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var episode bson.M
		if err = cursor.Decode(&episode); err != nil {
			log.Fatal(err)
		}
		fmt.Println(episode)
	}
}

func findOne(ctx context.Context, db *mongo.Database) {
	podcastsCollection := db.Collection("podcasts")
	var podcast bson.M

	if err := podcastsCollection.FindOne(ctx, bson.M{}).Decode(&podcast); err != nil {
		log.Fatal(err)
	}
	fmt.Println(podcast)
}

func filter(ctx context.Context, db *mongo.Database) {
	episodesCollection := db.Collection("episodes")

	filterCursor, err := episodesCollection.Find(ctx, bson.M{"duration": 25})
	if err != nil {
		log.Fatal(err)
	}
	var episodesFiltered []bson.M
	if err = filterCursor.All(ctx, &episodesFiltered); err != nil {
		log.Fatal(err)
	}
	fmt.Println(episodesFiltered)
}

func sorting(ctx context.Context, db *mongo.Database) {
	episodesCollection := db.Collection("episodes")
	opts := options.Find()
	opts.SetSort(bson.D{{"duration", -1}})
	sortCursor, err := episodesCollection.Find(ctx, bson.D{{"duration", bson.D{{"$gt", 24}}}}, opts)
	if err != nil {
		log.Fatal(err)
	}
	var episodesSorted []bson.M
	if err = sortCursor.All(ctx, &episodesSorted); err != nil {
		log.Fatal(err)
	}
	fmt.Println(episodesSorted)
}
