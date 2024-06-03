package usecase

import (
	"errors"
	"fmt"
	"time"

	"github.com/decorickey/go-apps/bmonster/domain/entity"
)

var (
	ErrScheduleRepository = errors.New("schedule repository error")
	ErrValidation         = errors.New("validation error")
)

type ScheduleQuery struct {
	Studio    string
	StartAt   time.Time
	Performer string
	Vol       string
}

type ScheduleQueryUsecase interface {
	FetchAll() ([]ScheduleQuery, error)
}

func NewScheduleQueryUsecase(scheduleRepo entity.ScheduleRepository) ScheduleQueryUsecase {
	return scheduleQueryUsecase{
		scheduleRepo: scheduleRepo,
	}
}

type scheduleQueryUsecase struct {
	scheduleRepo entity.ScheduleRepository
}

func (u scheduleQueryUsecase) FetchAll() ([]ScheduleQuery, error) {
	schedules, err := u.scheduleRepo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("faild to find schedules: %w", err)
	}
	queries := make([]ScheduleQuery, len(schedules))
	for i, schedule := range schedules {
		queries[i] = ScheduleQuery{
			Studio:    schedule.Studio,
			StartAt:   schedule.StartAt,
			Performer: schedule.Performer.Name,
			Vol:       schedule.Vol,
		}
	}
	return queries, nil
}

type ScheduleCommand struct {
	Studio     string
	StartYear  string
	StartMonth string
	StartDay   string
	StartHour  string
	StartMin   string
	Performer  string
	Vol        string
	Err        error
}

type ScheduleCommandUsecase interface {
	BulkUpsert(commands []ScheduleCommand) ([]ScheduleCommand, error)
}

func NewScheduleCommandUsecase(scheduleRepo entity.ScheduleRepository) ScheduleCommandUsecase {
	return scheduleCommandUsecase{
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
		performer, err := entity.NewPerformer(command.Performer)
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

		schedule, err := entity.NewSchedule(*program, command.Studio, time.Now())
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