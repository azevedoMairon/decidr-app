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
	slog.Info("VoteService.PostVote called", "participant_id", req.ParticipantId)

	participantId, err := primitive.ObjectIDFromHex(req.ParticipantId)
	if err != nil {
		slog.Warn("Failed to convert participant ID", "participant_id", req.ParticipantId, "error", err.Error())
		return nil, fmt.Errorf("failed to convert participantId: %w", err)
	}

	slog.Info("Checking if participant is nominated", "participant_id", participantId.Hex())
	exists, err := s.participantRepo.IsNominated(ctx, participantId)
	if err != nil {
		slog.Error("Error checking participant nomination", "participant_id", participantId.Hex(), "error", err.Error())
		return nil, err
	}
	if !exists {
		slog.Warn("Vote rejected - participant not nominated", "participant_id", participantId.Hex())
		return nil, errors.New("invalid or not nominated participant")
	}

	slog.Info("Participant validated, incrementing vote", "participant_id", participantId.Hex())
	result, err := s.voteRepo.IncrementVote(ctx, req)
	if err != nil {
		slog.Error("Failed to increment vote", "participant_id", participantId.Hex(), "error", err.Error())
		return nil, err
	}

	slog.Info("Vote incremented successfully", "participant_id", participantId.Hex(), "matched_count", result.MatchedCount)
	return result, nil
}

func (s *voteService) GetResults(ctx context.Context, byHour *bool) ([]entities.VoteResult, error) {
	slog.Info("VoteService.GetResults called", "by_hour", byHour)

	results, err := s.voteRepo.FindAll(ctx, byHour)
	if err != nil {
		slog.Error("Failed to retrieve results", "error", err.Error())
		return nil, err
	}

	slog.Info("Results retrieved successfully", "result_count", len(results))
	return results, nil
}
