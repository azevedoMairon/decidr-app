package services

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/azevedoMairon/decidr-app/internal/entities"
	"github.com/azevedoMairon/decidr-app/internal/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type VoteService interface {
	PostVote(ctx context.Context, req entities.VoteRequest) (*mongo.UpdateResult, error)
	GetResults(ctx context.Context, byHour *bool) ([]entities.VoteResult, error)
}

type voteService struct {
	voteRepo        repositories.VoteRepository
	participantRepo repositories.ParticipantRepository
}

func NewVoteService(
	voteRepo repositories.VoteRepository,
	participantRepo repositories.ParticipantRepository,
) VoteService {
	return &voteService{
		voteRepo:        voteRepo,
		participantRepo: participantRepo,
	}
}

func (s *voteService) PostVote(ctx context.Context, req entities.VoteRequest) (*mongo.UpdateResult, error) {
	participantId, err := primitive.ObjectIDFromHex(req.ParticipantId)
	if err != nil {
		slog.Warn("[VoteService.PostVote] Failed to convert participant ID", "participant_id", req.ParticipantId, "error", err.Error())
		return nil, fmt.Errorf("failed to convert participantId: %w", err)
	}

	exists, err := s.participantRepo.IsNominated(ctx, participantId)
	if err != nil {
		slog.Error("[VoteService.PostVote] Error checking participant nomination", "participant_id", participantId.Hex(), "error", err.Error())
		return nil, err
	}
	if !exists {
		slog.Warn("[VoteService.PostVote] Vote rejected - participant not nominated", "participant_id", participantId.Hex())
		return nil, errors.New("invalid or not nominated participant")
	}

	return s.voteRepo.IncrementVote(ctx, req)
}

func (s *voteService) GetResults(ctx context.Context, byHour *bool) ([]entities.VoteResult, error) {
	return s.voteRepo.FindAll(ctx, byHour)
}
