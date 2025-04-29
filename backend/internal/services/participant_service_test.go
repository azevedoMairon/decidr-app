package services

import (
	"context"
	"testing"

	"github.com/azevedoMairon/decidr-app/internal/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockRepo struct {
	mock.Mock
}

func (m *mockRepo) FindAll(ctx context.Context) ([]entities.Participant, error) {
	args := m.Called(ctx)
	return args.Get(0).([]entities.Participant), args.Error(1)
}

func TestGetParticipants_ReturnsList(t *testing.T) {
	ctx := context.Background()

	expected := []entities.Participant{
		{ID: "1", Name: "Juliette", IsNominated: false, IsEliminated: false},
		{ID: "2", Name: "Gil do Vigor", IsNominated: false, IsEliminated: false},
	}

	mockRepo := new(mockRepo)
	mockRepo.On("FindAll", ctx).Return(expected, nil)

	svc := NewService(mockRepo)

	result, err := svc.GetParticipants(ctx)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}
