package unit

import (
	"context"
	"errors"
	"testing"

	"github.com/azevedoMairon/decidr-app/internal/entities"
	"github.com/azevedoMairon/decidr-app/internal/services"
	"github.com/azevedoMairon/decidr-app/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetParticipants(t *testing.T) {
	t.Run("should return all participants when no filter is applied", func(t *testing.T) {
		ctx := context.Background()
		expected := []entities.Participant{
			{ID: "1", Name: "Juliette", IsNominated: false, IsEliminated: false},
			{ID: "2", Name: "Gil do Vigor", IsNominated: false, IsEliminated: false},
		}

		mockRepo := new(mocks.MockParticipantRepo)
		mockRepo.On("FindAll", ctx, mock.Anything).Return(expected, nil)

		svc := services.NewParticipantService(mockRepo)

		result, err := svc.GetParticipants(ctx, nil)

		assert.NoError(t, err)
		assert.Equal(t, expected, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return only nominated participants when filter is true", func(t *testing.T) {
		ctx := context.Background()
		nominated := true

		expected := []entities.Participant{
			{ID: "2", Name: "Gil do Vigor", IsNominated: true, IsEliminated: false},
		}

		mockRepo := new(mocks.MockParticipantRepo)
		mockRepo.On("FindAll", ctx, mock.MatchedBy(func(b *bool) bool {
			return b != nil && *b
		})).Return(expected, nil)

		svc := services.NewParticipantService(mockRepo)

		result, err := svc.GetParticipants(ctx, &nominated)

		assert.NoError(t, err)
		assert.Equal(t, expected, result)
		mockRepo.AssertCalled(t, "FindAll", ctx, &nominated)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when repository fails", func(t *testing.T) {
		ctx := context.Background()
		mockRepo := new(mocks.MockParticipantRepo)

		mockRepo.On("FindAll", ctx, mock.Anything).Return([]entities.Participant(nil), errors.New("db failure"))

		svc := services.NewParticipantService(mockRepo)
		result, err := svc.GetParticipants(ctx, nil)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.EqualError(t, err, "db failure")
		mockRepo.AssertExpectations(t)
	})
}
