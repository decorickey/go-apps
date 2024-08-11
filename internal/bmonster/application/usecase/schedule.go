package usecase

import (
	"github.com/decorickey/go-apps/internal/bmonster/application/dto"
)

type ScheduleQueryUsecase interface {
	FetchAll() (dto.Schedules, error)
}
