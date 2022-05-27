package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Podcast represents the schema for the "Podcasts" collection
type Podcast struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Title  string             `bson:"title,omitempty"`
	Author string             `bson:"author,omitempty"`
	Tags   []string           `bson:"tags,omitempty"`
}

// Episode represents the schema for the "Episodes" collection
type Episode struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Podcast     primitive.ObjectID `bson:"podcast,omitempty"`
	Title       string             `bson:"title,omitempty"`
	Description string             `bson:"description,omitempty"`
	Duration    int32              `bson:"duration,omitempty"`
}

// PodcastEpisode represents an aggregation result-set for two collections
type PodcastEpisode struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Podcast     Podcast            `bson:"podcast,omitempty"`
	Title       string             `bson:"title,omitempty"`
	Description string             `bson:"description,omitempty"`
	Duration    int32              `bson:"duration,omitempty"`
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)

	database := client.Database("quickstart")
	episodesCollection := database.Collection("episodes")

	id, _ := primitive.ObjectIDFromHex("60cc4994be6d99bf2cea8896")

	// in this case a matching stage and grouping stage.
	// In the matching stage we are matching all documents
	// that have the podcast field in question. In the grouping
	// stage we are grouping the matches by the podcast field
	// because it is non-distinct, and then we are summing each
	// of the duration fields into a new total field.
	matchStage := bson.D{{"$match", bson.D{{"podcast", id}}}}
	groupStage := bson.D{{"$group", bson.D{{"_id", "$podcast"}, {"total", bson.D{{"$sum", "$duration"}}}}}}

	showInfoCursor, err := episodesCollection.Aggregate(ctx, mongo.Pipeline{matchStage, groupStage})
	if err != nil {
		panic(err)
	}
	var showsWithInfo []bson.M
	if err = showInfoCursor.All(ctx, &showsWithInfo); err != nil {
		panic(err)
	}
	fmt.Println(showsWithInfo)

	//
	lookupStage := bson.D{{"$lookup", bson.D{{"from", "podcasts"}, {"localField", "podcast"}, {"foreignField", "_id"}, {"as", "podcast"}}}}
	unwindStage := bson.D{{"$unwind", bson.D{{"path", "$podcast"}, {"preserveNullAndEmptyArrays", false}}}}

	showLoadedCursor, err := episodesCollection.Aggregate(ctx, mongo.Pipeline{lookupStage, unwindStage})
	if err != nil {
		panic(err)
	}
	// var showsLoaded []bson.M
	var showsLoadedStruct []PodcastEpisode

	// if err = showLoadedCursor.All(ctx, &showsLoaded); err != nil {
	if err = showLoadedCursor.All(ctx, &showsLoadedStruct); err != nil {
		panic(err)
	}
	fmt.Println(showsLoadedStruct)

}
