package repositories

import (
	"context"
	"time"

	"github.com/azevedoMairon/decidr-app/internal/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type VoteRepository interface {
	IncrementVote(ctx context.Context, req entities.VoteRequest) (*mongo.UpdateResult, error)
	FindAll(ctx context.Context, byHour *bool) ([]entities.VoteResult, error)
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
	filter := bson.M{
		"participantId": req.ParticipantId,
		"hour":          time.Now().Truncate(time.Hour),
	}

	update := bson.M{
		"$inc": bson.M{"count": 1},
	}

	result, err := r.collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *voteRepository) FindAll(ctx context.Context, byHour *bool) ([]entities.VoteResult, error) {
	var pipeline mongo.Pipeline

	if byHour == nil || !*byHour {
		pipeline = getStandardAggregation()
	} else {
		pipeline = getByHourAggregation()
	}

	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var result []entities.VoteResult
	if err := cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func getStandardAggregation() mongo.Pipeline {
	return mongo.Pipeline{
		{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: "$participantId"},
			{Key: "count", Value: bson.D{{Key: "$sum", Value: "$count"}}},
		}}},
		{{Key: "$project", Value: bson.D{
			{Key: "_id", Value: 0},
			{Key: "participantId", Value: "$_id"},
			{Key: "count", Value: "$count"},
		}}},
	}
}

func getByHourAggregation() mongo.Pipeline {
	return mongo.Pipeline{
		{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: bson.D{
				{Key: "participantId", Value: "$participantId"},
				{Key: "hour", Value: "$hour"},
			}},
			{Key: "count", Value: bson.D{{Key: "$sum", Value: "$count"}}},
		}}},
		{{Key: "$project", Value: bson.D{
			{Key: "_id", Value: 0},
			{Key: "participantId", Value: "$_id.participantId"},
			{Key: "hour", Value: "$_id.hour"},
			{Key: "count", Value: "$count"},
		}}},
		{{Key: "$sort", Value: bson.D{
			{Key: "hour", Value: 1},
		}}},
	}
}
