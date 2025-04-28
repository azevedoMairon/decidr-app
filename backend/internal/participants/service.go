package participants

import (
	"context"

	"github.com/azevedoMairon/decidr-app/core/entities"
)

type Service interface {
	GetParticipants(ctx context.Context) ([]entities.Participant, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetParticipants(ctx context.Context) ([]entities.Participant, error) {
	return s.repo.FindAll(ctx)
}
