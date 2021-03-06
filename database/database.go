package database

import (
	"context"
	"newsapp/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB(ctx context.Context, conf config.MongoConfiguration) *mongo.Database {
	connection := options.Client().ApplyURI(conf.Server)
	client, err := mongo.Connect(context.TODO(), connection)

	if err != nil {
		panic(err)
	}

	return client.Database(conf.Database)

}
