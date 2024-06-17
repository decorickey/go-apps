package usecase

import (
	"slices"

	"github.com/decorickey/go-apps/internal/bmonster/domain/entity"
)

type PerformerQueryCommand struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func NewPerformerQueryCommandFromEntity(e entity.Performer) *PerformerQueryCommand {
	return &PerformerQueryCommand{ID: e.ID, Name: e.Name}
}

type PerformerQueryUsecase interface {
	List() []PerformerQueryCommand
}

func NewPerformerQueryUsecase(performerRepo entity.PerformerRepository) PerformerQueryUsecase {
	return &performerQueryUsecase{
		performerRepo: performerRepo,
	}
}

type performerQueryUsecase struct {
	performerRepo entity.PerformerRepository
}

func (u performerQueryUsecase) List() []PerformerQueryCommand {
	performers, _ := u.performerRepo.List()

	queries := make([]PerformerQueryCommand, 0)
	for _, performer := range performers {
		queries = append(queries, *NewPerformerQueryCommandFromEntity(performer))
	}

	slices.SortFunc(queries, func(a, b PerformerQueryCommand) int {
		return a.ID - b.ID
	})
	return queries
}
