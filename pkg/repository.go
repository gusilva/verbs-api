package pkg

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type VerbRepository interface {
	FindVerb(ctx context.Context, verbName string) (*VerbConjugation, error)
	SearchVerb(ctx context.Context, queryWord string) ([]string, error)
}

type VerbRepo struct {
	db *mongo.Collection
}

func CreateVerbRepository(db *mongo.Collection) *VerbRepo {
	return &VerbRepo{db}
}

func (repo *VerbRepo) FindVerb(ctx context.Context, verbName string) (*VerbConjugation, error) {
	var verbConjugation VerbConjugation

	if findError := repo.db.FindOne(ctx, bson.M{"verbo": verbName}).Decode(&verbConjugation); findError != nil {
		return nil, fmt.Errorf("verb not found: %w", findError)
	}

	return &verbConjugation, nil
}

func (repo *VerbRepo) SearchVerb(ctx context.Context, queryWord string) ([]string, error) {
	query := bson.M{"verbo": bson.M{"$regex": queryWord, "$options": "im"}}
	opts := options.Find().SetProjection(bson.D{{"_id", 0}, {"verbo", 1}})

	cursor, errorFind := repo.db.Find(ctx, query, opts)

	if errorFind != nil {
		return nil, fmt.Errorf("error on find: %w", errorFind)
	}

	var verbsResult []SearchResult
	if errorDecode := cursor.All(context.TODO(), &verbsResult); errorDecode != nil {
		return nil, fmt.Errorf("error decoding db result: %w", errorDecode)
	}

	var verbNames []string
	if len(verbsResult) != 0 {
		verbNames = make([]string, len(verbsResult))
		for i := 0; i < len(verbsResult); i++ {
			verbNames[i] = verbsResult[i].Verbo
		}
	}

	return verbNames, nil
}
