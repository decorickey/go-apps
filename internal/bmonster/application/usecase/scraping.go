package usecase

import (
	"fmt"

	"github.com/decorickey/go-apps/internal/bmonster/domain/entity"
	"github.com/decorickey/go-apps/internal/bmonster/domain/repository"
)

type ScrapingUsecase interface {
	Handle() ([]entity.Studio, []entity.Performer, []entity.Program, []entity.Schedule, error)
	Save([]entity.Studio, []entity.Performer, []entity.Program, []entity.Schedule) error
}

func NewScrapingUsecase(
	scrapingRepo repository.ScrapingRepository,
	studioRepo repository.StudioRepository,
	performerRepo repository.PerformerRepository,
	programRepo repository.ProgramRepository,
	scheduleRepo repository.ScheduleRepository,
) ScrapingUsecase {
	return &scrapingUsecase{
		scrapingRepo:  scrapingRepo,
		studioRepo:    studioRepo,
		performerRepo: performerRepo,
		programRepo:   programRepo,
		scheduleRepo:  scheduleRepo,
	}
}

type scrapingUsecase struct {
	scrapingRepo  repository.ScrapingRepository
	studioRepo    repository.StudioRepository
	performerRepo repository.PerformerRepository
	programRepo   repository.ProgramRepository
	scheduleRepo  repository.ScheduleRepository
}

func (u scrapingUsecase) Handle() ([]entity.Studio, []entity.Performer, []entity.Program, []entity.Schedule, error) {
	studios, err := u.scrapingRepo.FetchStudios()
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("fetch studios: %w", err)
	}

	performers, programs, schedules, err := u.scrapingRepo.FetchSchedulesByStudios(studios)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("fetch schedules: %w", err)
	}

	return studios, performers, programs, schedules, nil
}

func (u scrapingUsecase) Save(studios []entity.Studio, performers []entity.Performer, programs []entity.Program, schedules []entity.Schedule) error {
	if err := u.studioRepo.Save(studios); err != nil {
		return fmt.Errorf("failed to save Studios: %w", err)
	}

	if err := u.performerRepo.Save(performers); err != nil {
		return fmt.Errorf("failed to save Performers: %w", err)
	}

	if err := u.programRepo.Save(programs); err != nil {
		return fmt.Errorf("failed to save Programs: %w", err)
	}

	if err := u.scheduleRepo.Save(schedules); err != nil {
		return fmt.Errorf("failed to save Schedules: %w", err)
	}

	return nil
}
