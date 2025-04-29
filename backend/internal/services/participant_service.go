package services

import (
	"context"

	"github.com/azevedoMairon/decidr-app/internal/entities"
	"github.com/azevedoMairon/decidr-app/internal/repositories"
)

type ParticipantService interface {
	GetParticipants(ctx context.Context, isNominated *bool) ([]entities.Participant, error)
}

type participantService struct {
	repo repositories.ParticipantRepository
}

func NewService(repo repositories.ParticipantRepository) ParticipantService {
	return &participantService{repo: repo}
}

func (s *participantService) GetParticipants(ctx context.Context, isNominated *bool) ([]entities.Participant, error) {
	return s.repo.FindAll(ctx, isNominated)
}
