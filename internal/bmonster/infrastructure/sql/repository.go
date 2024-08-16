package sql

import (
	"github.com/decorickey/go-apps/internal/bmonster/domain/entity"
	"github.com/decorickey/go-apps/internal/bmonster/domain/repository"
	"gorm.io/gorm"
)

func NewStudioRepository(db *gorm.DB) repository.StudioRepository {
	return studioRepository{db: db}
}

type studioRepository struct {
	db *gorm.DB
}

func (repo studioRepository) Save(entities []entity.Studio) error {
	records := make([]*Studio, len(entities))
	for i, v := range entities {
		records[i] = &Studio{ID: v.ID, Name: v.Name}
	}

	if err := repo.db.Create(records).Error; err != nil {
		return err
	}
	return nil
}

func NewPerformerRepository(db *gorm.DB) repository.PerformerRepository {
	return performerRepository{db: db}
}

type performerRepository struct {
	db *gorm.DB
}

func (repo performerRepository) Save(entities []entity.Performer) error {
	records := make([]*Performer, len(entities))
	for i, v := range entities {
		records[i] = &Performer{ID: v.ID, Name: v.Name}
	}
	if err := repo.db.Create(records).Error; err != nil {
		return err
	}
	return nil
}

func NewProgramRepository(db *gorm.DB) repository.ProgramRepository {
	return programRepository{db: db}
}

type programRepository struct {
	db *gorm.DB
}

func (repo programRepository) Save(entities []entity.Program) error {
	records := make([]*Program, len(entities))
	for i, v := range entities {
		records[i] = &Program{ID: v.ID, Name: v.Name}
	}

	if err := repo.db.Create(records).Error; err != nil {
		return err
	}
	return nil
}

func NewScheduleRepository(db *gorm.DB) repository.ScheduleRepository {
	return scheduleRepository{db: db}
}

type scheduleRepository struct {
	db *gorm.DB
}

func (repo scheduleRepository) Save(entities []entity.Schedule) error {
	records := make([]*Schedule, len(entities))
	for i, v := range entities {
		records[i] = &Schedule{
			ID:          v.ID,
			StudioID:    v.StudioID,
			ProgramID:   v.ProgramID,
			PerformerID: v.PerformerID,
			StartAt:     v.StartAt,
			EndAt:       v.EndAt,
			HashID:      v.HashID,
		}
	}

	if err := repo.db.Create(records).Error; err != nil {
		return err
	}
	return nil
}
