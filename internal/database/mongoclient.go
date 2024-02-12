package database

import (
	"context"
	"fmt"

	interfaces "github.com/LanternNassi/IMSController/internal/Interfaces"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoClient struct {
}

func NewMongoDatabaseClient() (interfaces.MongoDatabaseClient, error) {
	client := MongoClient{}

	return client, nil
}

func (client_db MongoClient) CloseMongo(client *mongo.Client, ctx context.Context) {

	// client provides a method to close
	// a mongoDB connection.
	defer func() {

		fmt.Println("Mongo db closing successfully ....")

		if err := client.Disconnect(ctx); err != nil {
			panic(err)

		}
	}()
}

func (client_db MongoClient) ConnectMongo(uri string) (*mongo.Client, context.Context, error) {

	ctx := context.TODO()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, err
}

func (client_db MongoClient) PingMongo(client *mongo.Client, ctx context.Context) error {

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	fmt.Println("connected successfully")
	return nil
}
