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

func (m *mockRepo) FindAll(ctx context.Context, isNominated *bool) ([]entities.Participant, error) {
	args := m.Called(ctx, isNominated)
	return args.Get(0).([]entities.Participant), args.Error(1)
}

func (m *mockRepo) IsNominated(ctx context.Context, id string) (bool, error) {
	args := m.Called(ctx, id)
	return args.Bool(0), args.Error(1)
}

func Test_GetParticipants_ShouldReturnDbResponse(t *testing.T) {
	ctx := context.Background()

	expected := []entities.Participant{
		{ID: "1", Name: "Juliette", IsNominated: false, IsEliminated: false},
		{ID: "2", Name: "Gil do Vigor", IsNominated: false, IsEliminated: false},
	}

	mockRepo := new(mockRepo)
	mockRepo.On("FindAll", ctx, mock.Anything).Return(expected, nil)

	svc := NewParticipantService(mockRepo)

	result, err := svc.GetParticipants(ctx, nil)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func Test_GetParticipants_ShouldSendNominatedFilterValue(t *testing.T) {
	ctx := context.Background()
	nominated := true

	expected := []entities.Participant{
		{ID: "2", Name: "Gil do Vigor", IsNominated: true, IsEliminated: false},
	}

	mockRepo := new(mockRepo)
	mockRepo.On("FindAll", ctx, mock.MatchedBy(func(b *bool) bool {
		return b != nil && *b == true
	})).Return(expected, nil)

	svc := NewParticipantService(mockRepo)
	result, err := svc.GetParticipants(ctx, &nominated)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertCalled(t, "FindAll", ctx, &nominated)
}
