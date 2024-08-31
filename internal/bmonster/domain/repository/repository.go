//go:generate go run go.uber.org/mock/mockgen -source=repository.go -destination=mock_repository.go -package=repository
package repository

import "github.com/decorickey/go-apps/internal/bmonster/domain/entity"

type StudioRepository interface {
	Save([]entity.Studio) error
}

type PerformerRepository interface {
	Save([]entity.Performer) error
}

type ProgramRepository interface {
	Save([]entity.Program) error
}

type ScheduleRepository interface {
	Save([]entity.Schedule) error
}
