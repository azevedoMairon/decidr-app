package repositories

import (
	"context"
	"errors"
	"log/slog"

	"github.com/azevedoMairon/decidr-app/internal/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ParticipantRepository interface {
	FindAll(ctx context.Context, isNominated *bool) ([]entities.Participant, error)
	IsNominated(ctx context.Context, id primitive.ObjectID) (bool, error)
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
		slog.Info("[ParticipantRepository.FindAll] Fetching nominated participants", "filter", filter)
	} else {
		slog.Info("[ParticipantRepository.FindAll] Fetching all participants")
	}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		slog.Error("[ParticipantRepository.FindAll] Failed to execute Find() on participants", "error", err.Error(), "filter", filter)
		return nil, err
	}
	defer cursor.Close(ctx)

	var participants []entities.Participant
	if err := cursor.All(ctx, &participants); err != nil {
		slog.Error("[ParticipantRepository.FindAll] Failed to decode participant documents", "error", err.Error())
		return nil, err
	}

	return participants, nil
}

func (r *participantRepository) IsNominated(ctx context.Context, id primitive.ObjectID) (bool, error) {
	filter := bson.M{"_id": id, "isNominated": true}

	err := r.collection.FindOne(ctx, filter).Err()
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			slog.Info("[ParticipantRepository.IsNominated] Participant not nominated or not found", "participant_id", id.Hex())
			return false, nil
		}
		slog.Error("[ParticipantRepository.IsNominated] Error checking nomination status", "participant_id", id.Hex(), "error", err.Error())
		return false, err
	}

	return true, nil
}
