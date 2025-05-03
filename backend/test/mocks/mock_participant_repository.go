package mocks

import (
	"context"

	"github.com/azevedoMairon/decidr-app/internal/entities"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockParticipantRepo struct {
	mock.Mock
}

func (m *MockParticipantRepo) FindAll(ctx context.Context, isNominated *bool) ([]entities.Participant, error) {
	args := m.Called(ctx, isNominated)
	return args.Get(0).([]entities.Participant), args.Error(1)
}

func (m *MockParticipantRepo) IsNominated(ctx context.Context, id primitive.ObjectID) (bool, error) {
	args := m.Called(ctx, id)
	return args.Bool(0), args.Error(1)
}
