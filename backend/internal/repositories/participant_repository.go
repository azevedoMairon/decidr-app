package repositories

import (
	"context"

	"github.com/azevedoMairon/decidr-app/internal/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ParticipantRepository interface {
	FindAll(ctx context.Context, isNominated *bool) ([]entities.Participant, error)
}

type participantRepository struct {
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database) ParticipantRepository {
	return &participantRepository{
		collection: db.Collection("participants"),
	}
}

func (r *participantRepository) FindAll(ctx context.Context, isNominated *bool) ([]entities.Participant, error) {
	filter := bson.M{}

	if isNominated != nil {
		filter["isNominated"] = *isNominated
	}

	cursor, err := r.collection.Find(ctx, filter)
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
