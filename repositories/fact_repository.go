package repositories

import (
	"challenge/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const factCollection = "facts"

type FactRepository interface {
	GetFacts(context.Context) ([]models.Fact, error)
}

type factRepository struct {
	db             *mongo.Database
	factCollection string
}

func (f factRepository) GetFacts(ctx context.Context) (facts []models.Fact, err error) {
	cur, err := f.db.Collection(f.factCollection).Find(ctx, bson.D{{}})
	if err != nil {
		return
	}

	for cur.Next(context.TODO()) {
		f := models.Fact{}

		err := cur.Decode(&f)
		if err != nil {
			return nil, err
		}

		facts = append(facts, f)
	}

	return
}

func NewFactRepository(db *mongo.Database) factRepository {

	return factRepository{
		db:             db,
		factCollection: factCollection,
	}
}
