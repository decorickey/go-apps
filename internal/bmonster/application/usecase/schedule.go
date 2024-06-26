package usecase

import (
	"errors"
	"fmt"
	"time"

	"github.com/decorickey/go-apps/internal/bmonster/domain/entity"
	"github.com/decorickey/go-apps/pkg/timeutil"
)

var (
	ErrScheduleRepository = errors.New("schedule repository error")
	ErrValidation         = errors.New("validation error")
)

type ScheduleQuery struct {
	Studio    string                `json:"studio"`
	StartAt   time.Time             `json:"startAt"`
	Performer PerformerQueryCommand `json:"performer"`
	Vol       string                `json:"vol"`
}

func NewScheduleQueryFromEntity(e entity.Schedule) *ScheduleQuery {
	return &ScheduleQuery{
		Studio:    e.Studio,
		StartAt:   e.StartAt,
		Performer: *NewPerformerQueryCommandFromEntity(e.Performer),
		Vol:       e.Vol,
	}
}

type ScheduleQueryUsecase interface {
	All() ([]ScheduleQuery, error)
}

func NewScheduleQueryUsecase(scheduleRepo entity.ScheduleRepository) ScheduleQueryUsecase {
	return &scheduleQueryUsecase{
		scheduleRepo: scheduleRepo,
	}
}

type scheduleQueryUsecase struct {
	scheduleRepo entity.ScheduleRepository
}

func (u scheduleQueryUsecase) All() ([]ScheduleQuery, error) {
	schedules, err := u.scheduleRepo.All()
	if err != nil {
		return nil, fmt.Errorf("faild to find schedules: %w", err)
	}
	queries := make([]ScheduleQuery, len(schedules))
	for i, schedule := range schedules {
		queries[i] = *NewScheduleQueryFromEntity(schedule)
	}
	return queries, nil
}

type ScheduleCommand struct {
	Studio        string
	StartYear     int
	StartMonth    time.Month
	StartDay      int
	StartHour     int
	StartMin      int
	PerformerID   int
	PerformerName string
	Vol           string
	Err           error
}

func (c ScheduleCommand) StartAt() time.Time {
	return time.Date(c.StartYear, c.StartMonth, c.StartDay, c.StartHour, c.StartMin, 0, 0, timeutil.JST)
}

func (c ScheduleCommand) ToQuery() *ScheduleQuery {
	return &ScheduleQuery{
		Studio:    c.Studio,
		StartAt:   c.StartAt(),
		Performer: PerformerQueryCommand{ID: c.PerformerID, Name: c.PerformerName},
		Vol:       c.Vol,
	}
}

type ScheduleCommandUsecase interface {
	BulkUpsert(commands []ScheduleCommand) ([]ScheduleCommand, error)
}

func NewScheduleCommandUsecase(scheduleRepo entity.ScheduleRepository) ScheduleCommandUsecase {
	return &scheduleCommandUsecase{
		scheduleRepo: scheduleRepo,
	}
}

type scheduleCommandUsecase struct {
	scheduleRepo entity.ScheduleRepository
}

func (u scheduleCommandUsecase) BulkUpsert(commands []ScheduleCommand) ([]ScheduleCommand, error) {
	var hasErr bool
	schedules := make([]entity.Schedule, len(commands))

	for i, command := range commands {
		performer, err := entity.NewPerformer(command.PerformerID, command.PerformerName)
		if err != nil {
			hasErr = true
			commands[i].Err = ErrValidation
			continue
		}

		program, err := entity.NewProgram(*performer, command.Vol)
		if err != nil {
			hasErr = true
			commands[i].Err = ErrValidation
			continue
		}

		schedule, err := entity.NewSchedule(*program, command.Studio, command.StartAt())
		if err != nil {
			hasErr = true
			commands[i].Err = ErrValidation
			continue
		}

		schedules[i] = *schedule
	}

	if hasErr {
		return commands, errors.New("invalid commands")
	}

	err := u.scheduleRepo.Save(schedules)
	if err != nil {
		for i := range commands {
			commands[i].Err = ErrScheduleRepository
		}
		return commands, fmt.Errorf("faild to save schedules: %w", err)
	}
	return commands, nil
}
