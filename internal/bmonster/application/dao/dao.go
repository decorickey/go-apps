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

type TimetableDAO interface {
	FetchByStudioIDAndDate(uint, time.Time) (dto.Timetable, error)
	FetchByPerformerIDAndDate(uint, time.Time) (dto.Timetable, error)
}