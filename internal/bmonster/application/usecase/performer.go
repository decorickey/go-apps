package usecase

import (
	"fmt"

	"github.com/decorickey/go-apps/internal/bmonster/domain/entity"
	"github.com/decorickey/go-apps/internal/bmonster/domain/repository"
)

type PerformerQueryUsecase interface {
	FetchAll() ([]entity.Performer, error)
}

func NewPerformerQueryUsecase(performerRepo repository.PerformerRepository) PerformerQueryUsecase {
	return performerQueryUsecase{
		performerRepo: performerRepo,
	}
}

type performerQueryUsecase struct {
	performerRepo repository.PerformerRepository
}

func (u performerQueryUsecase) FetchAll() ([]entity.Performer, error) {
	performers, err := u.performerRepo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch performers: %w", err)
	}
	return performers, nil
}
