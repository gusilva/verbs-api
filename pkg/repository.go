package pkg

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type VerbRepository interface {
	FindVerb(ctx context.Context, verbName string) (*VerbConjugation, error)
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
