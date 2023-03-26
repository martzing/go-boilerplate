package mongoHelper

import (
	"context"
	"fmt"

	"github.com/martzing/go-boilerplate/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func connect(config *configs.MongoDbConfig) (*mongo.Client, *mongo.Database) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	db := client.Database(config.Database)

	return client, db
}
