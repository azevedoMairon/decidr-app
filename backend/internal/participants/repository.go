package participants

import (
	"context"

	"github.com/azevedoMairon/decidr-app/core/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	FindAll(ctx context.Context) ([]entities.Participant, error)
}

type repository struct {
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database) Repository {
	return &repository{
		collection: db.Collection("participants"),
	}
}

func (r *repository) FindAll(ctx context.Context) ([]entities.Participant, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var participants []entities.Participant
	if err := cursor.All(ctx, &participants); err != nil {
		return nil, err
	}

	return participants, nil
}
