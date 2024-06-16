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

// func NewPerformerQueryUsecase() PerformerQueryUsecase {}

func NewMockPerformerQueryUsecase(su ScrapingUsecase) PerformerQueryUsecase {
	return &mockPerformerQueryUsecase{su: su}
}

type mockPerformerQueryUsecase struct {
	su ScrapingUsecase
}

func (u mockPerformerQueryUsecase) List() []PerformerQueryCommand {
	res, _ := u.su.Performers()
	slices.SortFunc(res, func(a, b PerformerQueryCommand) int {
		return a.ID - b.ID
	})
	return res
}
