package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const keyCollection = "keys"

type KeyRepository interface {
	GetKey(context.Context, string) (string, error)
}

type keyRepository struct {
	db            *mongo.Database
	keyCollection string
}

func (k keyRepository) GetKey(ctx context.Context, key string) (string, error) {

	findCondations := bson.M{"key": key}

	var keyDocument = struct {
		Key string
	}{}

	err := k.db.Collection(k.keyCollection).FindOne(ctx, findCondations, &options.FindOneOptions{
		Projection: bson.M{"key": 1},
	}).Decode(&keyDocument)

	return keyDocument.Key, err
}

func NewKeyRepository(db *mongo.Database) keyRepository {

	return keyRepository{
		db:            db,
		keyCollection: keyCollection,
	}
}
