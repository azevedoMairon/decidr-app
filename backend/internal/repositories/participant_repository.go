package repositories

import (
	"context"
	"errors"

	"github.com/azevedoMairon/decidr-app/internal/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ParticipantRepository interface {
	FindAll(ctx context.Context, isNominated *bool) ([]entities.Participant, error)
	IsNominated(ctx context.Context, id string) (bool, error)
}

type participantRepository struct {
	collection *mongo.Collection
}

func NewParticipantRepository(db *mongo.Database) ParticipantRepository {
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

func (r *participantRepository) IsNominated(ctx context.Context, id string) (bool, error) {
	filter := bson.M{"_id": id, "isNominated": true}

	err := r.collection.FindOne(ctx, filter).Err()
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
