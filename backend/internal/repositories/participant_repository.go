package repositories

import (
	"context"

	"github.com/azevedoMairon/decidr-app/internal/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ParticipantRepository interface {
	FindAll(ctx context.Context) ([]entities.Participant, error)
}

type participantRepository struct {
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database) ParticipantRepository {
	return &participantRepository{
		collection: db.Collection("participants"),
	}
}

func (r *participantRepository) FindAll(ctx context.Context) ([]entities.Participant, error) {
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
