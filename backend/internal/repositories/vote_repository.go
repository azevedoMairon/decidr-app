package repositories

import (
	"context"

	"github.com/azevedoMairon/decidr-app/internal/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type VoteRepository interface {
	IncrementVote(ctx context.Context, req entities.VoteRequest) (*mongo.UpdateResult, error)
	FindAll(ctx context.Context) ([]entities.VoteResult, error)
}

type voteRepository struct {
	collection *mongo.Collection
}

func NewVoteRepository(db *mongo.Database) VoteRepository {
	return &voteRepository{
		collection: db.Collection("votes"),
	}
}

func (r *voteRepository) IncrementVote(ctx context.Context, req entities.VoteRequest) (*mongo.UpdateResult, error) {
	result, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": req.ParticipantId},
		bson.M{"$inc": bson.M{"count": 1}},
		options.Update().SetUpsert(true),
	)

	return result, err
}

func (r *voteRepository) FindAll(ctx context.Context) ([]entities.VoteResult, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var votes []entities.VoteResult
	if err := cursor.All(ctx, &votes); err != nil {
		return nil, err
	}

	return votes, nil
}
