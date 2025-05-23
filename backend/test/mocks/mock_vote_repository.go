package mocks

import (
	"context"

	"github.com/azevedoMairon/decidr-app/internal/entities"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
)

type MockVoteRepo struct {
	mock.Mock
}

func (m *MockVoteRepo) IncrementVote(ctx context.Context, req entities.VoteRequest) (*mongo.UpdateResult, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*mongo.UpdateResult), args.Error(1)
}

func (m *MockVoteRepo) FindAll(ctx context.Context, byHour *bool) ([]entities.VoteResult, error) {
	args := m.Called(ctx)
	return args.Get(0).([]entities.VoteResult), args.Error(1)
}
