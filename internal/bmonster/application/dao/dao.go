package dao

import (
	"time"

	"github.com/decorickey/go-apps/internal/bmonster/application/dto"
)

type StudioDAO interface {
	FetchAll() ([]dto.Studio, error)
}

type PerformerDAO interface {
	FetchAll() ([]dto.Performer, error)
}

type ScheduleDAO interface {
	FetchByDateAndStudio(time.Time, dto.Studio) (dto.TimeTable, error)
	FetchByDateAndPerformer(time.Time, dto.Performer) (dto.TimeTable, error)
}
