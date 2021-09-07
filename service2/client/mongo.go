package client

import (
	"context"
	"log"
	"time"

	"github.com/service2/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	parkingCollection *mongo.Collection
	ctx, _            = context.WithTimeout(context.Background(), 10*time.Second)

	mongoPort = config.Conf.MongoPort
)

func init() {
	clientOptions := options.Client().ApplyURI(mongoPort)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	parkingCollection = client.Database("test_db").Collection("parking")
}

func GetParkingCollection() *mongo.Collection {
	return parkingCollection
}
