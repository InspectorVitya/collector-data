package db

import (
	"context"
	config "github.com/gusleein/goconfig"
	log "github.com/gusleein/golog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoDB struct {
	client     *mongo.Client
	collection *mongo.Collection
}

var mongodb *mongoDB

func Init() {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.GetString("mongoUrl")))
	if err != nil {
		log.Fatal(err.Error())
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}

	demoDB := client.Database(config.GetString("dbName"))

	cNames, err := demoDB.ListCollectionNames(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err.Error())
	}
	cExist := false
	cName := config.GetString("collection")
	for i := range cNames {
		if cNames[i] == cName {
			cExist = true
		}
	}
	if !cExist {
		err = demoDB.CreateCollection(context.Background(), cName)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	collection := demoDB.Collection(cName)

	mongodb = &mongoDB{
		client:     client,
		collection: collection,
	}
}

func Stop() {
	if err := mongodb.client.Disconnect(context.Background()); err != nil {
		log.Error(err.Error())
	}
	log.Info("db connection stop...")
}
