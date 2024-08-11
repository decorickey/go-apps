package usecase

import (
	"fmt"

	"github.com/decorickey/go-apps/internal/bmonster/domain/entity"
	"github.com/decorickey/go-apps/internal/bmonster/domain/repository"
)

type StudioQueryUsecase interface {
	FetchAll() ([]entity.Studio, error)
}

func NewStudioQueryUsecase(studioRepo repository.StudioRepository) StudioQueryUsecase {
	return studioQueryUsecase{
		studioRepo: studioRepo,
	}
}

type studioQueryUsecase struct {
	studioRepo repository.StudioRepository
}

func (u studioQueryUsecase) FetchAll() ([]entity.Studio, error) {
	performers, err := u.studioRepo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch performers: %w", err)
	}
	return performers, nil
}
