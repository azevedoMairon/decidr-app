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
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestPostVote(t *testing.T) {
	validID := primitive.NewObjectID()
	validIDHex := validID.Hex()
	req := entities.VoteRequest{ParticipantId: validIDHex}

	t.Run("should return error if ParticipantId is not a valid ObjectID", func(t *testing.T) {
		mockVoteRepo := new(mocks.MockVoteRepo)
		mockParticipantRepo := new(mocks.MockParticipantRepo)
		svc := services.NewVoteService(mockVoteRepo, mockParticipantRepo)

		_, err := svc.PostVote(context.Background(), entities.VoteRequest{ParticipantId: "invalid-hex"})

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to convert participantId")
	})

	t.Run("should return error if participantRepo.IsNominated returns error", func(t *testing.T) {
		mockVoteRepo := new(mocks.MockVoteRepo)
		mockParticipantRepo := new(mocks.MockParticipantRepo)
		mockParticipantRepo.On("IsNominated", mock.Anything, validID).Return(false, errors.New("db error"))
		svc := services.NewVoteService(mockVoteRepo, mockParticipantRepo)

		_, err := svc.PostVote(context.Background(), req)

		assert.EqualError(t, err, "db error")
		mockParticipantRepo.AssertExpectations(t)
	})

	t.Run("should return error if participant is not nominated", func(t *testing.T) {
		mockVoteRepo := new(mocks.MockVoteRepo)
		mockParticipantRepo := new(mocks.MockParticipantRepo)

		mockParticipantRepo.On("IsNominated", mock.Anything, validID).Return(false, nil)
		svc := services.NewVoteService(mockVoteRepo, mockParticipantRepo)

		_, err := svc.PostVote(context.Background(), req)

		assert.EqualError(t, err, "invalid or not nominated participant")
		mockParticipantRepo.AssertExpectations(t)
	})

	t.Run("should return error if voteRepo.IncrementVote fails", func(t *testing.T) {
		mockVoteRepo := new(mocks.MockVoteRepo)
		mockParticipantRepo := new(mocks.MockParticipantRepo)

		mockVoteRepo.On("IncrementVote", mock.Anything, req).Return((*mongo.UpdateResult)(nil), errors.New("write error"))
		mockParticipantRepo.On("IsNominated", mock.Anything, validID).Return(true, nil)

		svc := services.NewVoteService(mockVoteRepo, mockParticipantRepo)

		_, err := svc.PostVote(context.Background(), req)
		assert.EqualError(t, err, "write error")
		mockParticipantRepo.AssertExpectations(t)
		mockVoteRepo.AssertExpectations(t)
	})

	t.Run("should return UpdateResult on success", func(t *testing.T) {
		mockParticipantRepo := new(mocks.MockParticipantRepo)
		mockVoteRepo := new(mocks.MockVoteRepo)

		expected := &mongo.UpdateResult{MatchedCount: 1, ModifiedCount: 1}

		mockParticipantRepo.On("IsNominated", mock.Anything, validID).Return(true, nil)
		mockVoteRepo.On("IncrementVote", mock.Anything, req).Return(expected, nil)

		svc := services.NewVoteService(mockVoteRepo, mockParticipantRepo)

		res, err := svc.PostVote(context.Background(), req)
		assert.NoError(t, err)
		assert.Equal(t, expected, res)

		mockParticipantRepo.AssertExpectations(t)
		mockVoteRepo.AssertExpectations(t)
	})
}
