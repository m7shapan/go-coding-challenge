package repositories

import (
	"challenge/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const factCollection = "facts"

type FactRepository interface {
	GetFacts(context.Context, *models.Filters) ([]models.Fact, int64, error)
}

type factRepository struct {
	db             *mongo.Database
	factCollection string
}

func (f factRepository) GetFacts(ctx context.Context, filter *models.Filters) (facts []models.Fact, total int64, err error) {

	findCondations := bson.D{}
	if filter.Search != "" {
		findCondations = append(findCondations, bson.E{Key: "text", Value: primitive.Regex{Pattern: filter.Search, Options: "i"}})
	}

	total, err = f.db.Collection(f.factCollection).CountDocuments(ctx, findCondations)
	if err != nil {
		return
	}

	cur, err := f.db.Collection(f.factCollection).Find(ctx, findCondations, &options.FindOptions{
		Limit: &filter.Limit,
		Skip:  &filter.Skip,
	})
	if err != nil {
		return
	}

	for cur.Next(context.TODO()) {
		f := models.Fact{}

		err = cur.Decode(&f)
		if err != nil {
			return
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
