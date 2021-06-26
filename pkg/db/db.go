package db

import (
	"challenge/pkg/config"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Connect(config config.DBConfig) (database *mongo.Database, err error) {
	client, err := mongo.NewClient(options.Client().
		ApplyURI(fmt.Sprintf("mongodb://%s:%d", config.Host, config.Port)).
		SetAuth(options.Credential{
			Username: config.Username,
			Password: config.Password,
		}))

	if err != nil {
		return
	}

	if err = client.Connect(context.TODO()); err != nil {
		return
	}

	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return
	}

	database = client.Database(config.Database)
	return
}
