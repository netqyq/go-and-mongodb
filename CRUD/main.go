package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Ping DB, detect
	// err = client.Ping(ctx, readpref.Primary())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// list db
	// databases, err := client.ListDatabaseNames(ctx, bson.M{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(databases)

	// insert
	// Insert(client)

	// readAllDocs(ctx, client)
	// readAllDocsIterate(ctx, client)

	quickstartDatabase := client.Database("quickstart")
	// findOne(ctx, quickstartDatabase)
	// filter(ctx, quickstartDatabase)
	// sorting(ctx, quickstartDatabase)
	// updataOne(ctx, quickstartDatabase)
	// updataMany(ctx, quickstartDatabase)
	// ReplaceOne(ctx, quickstartDatabase)
	// deleteOne(ctx, quickstartDatabase)
	// deleteMany(ctx, quickstartDatabase)
	deleteOne(ctx, quickstartDatabase)

	defer client.Disconnect(ctx)
}
