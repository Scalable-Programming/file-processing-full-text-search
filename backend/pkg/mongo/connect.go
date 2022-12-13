package mongo_db

import (
	"context"
	"fmt"
	"time"

	config "github.com/Scalable-Programming/file-processing-full-text-search/backend/pkg/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func close(client *mongo.Client, ctx context.Context,
	cancel context.CancelFunc) {

	defer cancel()

	defer func() {

		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func connect(uri string) (*mongo.Client, context.Context,
	context.CancelFunc, error) {

	ctx, cancel := context.WithTimeout(context.Background(),
		30*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

func ping(client *mongo.Client, ctx context.Context) error {
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	fmt.Println("connected successfully")
	return nil
}

func connect_mongodb() *mongo.Client {
	client, ctx, _, err := connect(config.AppConfig.MongoUri)
	if err != nil {
		panic(err)
	}

	// TODO (Matej)
	//defer close(client, ctx, cancel)

	ping(client, ctx)

	return client
}

var client = connect_mongodb()
var Db = client.Database(config.AppConfig.MongoDatabase)
