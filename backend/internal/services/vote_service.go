package services

import (
	"context"
	"errors"

	"github.com/azevedoMairon/decidr-app/internal/entities"
	"github.com/azevedoMairon/decidr-app/internal/repositories"
	"go.mongodb.org/mongo-driver/mongo"
)

type VoteService interface {
	PostVote(ctx context.Context, req entities.VoteRequest) (*mongo.UpdateResult, error)
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
	exists, err := s.participantRepo.IsNominated(ctx, req.ParticipantId)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("invalid or not nominated participant")
	}

	return s.voteRepo.IncrementVote(ctx, req)
}
