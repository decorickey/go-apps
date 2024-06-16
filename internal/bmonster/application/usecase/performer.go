package usecase

import "github.com/decorickey/go-apps/internal/bmonster/domain/entity"

type PerformerQueryCommand struct {
	ID   int
	Name string
}

func NewPerformerQueryCommandFromEntity(e entity.Performer) *PerformerQueryCommand {
	return &PerformerQueryCommand{ID: e.ID, Name: e.Name}
}
