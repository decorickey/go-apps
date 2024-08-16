package sql

import (
	"time"

	"github.com/decorickey/go-apps/internal/bmonster/application/dao"
	"github.com/decorickey/go-apps/internal/bmonster/application/dto"
	"gorm.io/gorm"
)

func NewStudioDao(db *gorm.DB) dao.StudioDAO {
	return studioDao{db: db}
}

type studioDao struct {
	db *gorm.DB
}

func (dao studioDao) FetchAll() ([]dto.Studio, error) {
	var records []Studio
	if err := dao.db.Find(&records).Error; err != nil {
		return nil, err
	}

	results := make([]dto.Studio, len(records))
	for i, v := range records {
		results[i] = dto.Studio{ID: v.ID, Name: v.Name}
	}
	return results, nil
}

func NewPerformerDao(db *gorm.DB) dao.PerformerDAO {
	return performerDao{db: db}
}

type performerDao struct {
	db *gorm.DB
}

func (dao performerDao) FetchAll() ([]dto.Performer, error) {
	var records []Performer
	if err := dao.db.Find(&records).Error; err != nil {
		return nil, err
	}

	results := make([]dto.Performer, len(records))
	for i, v := range records {
		results[i] = dto.Performer{ID: v.ID, Name: v.Name}
	}
	return results, nil
}

func NewTimetableDao(db *gorm.DB) dao.TimetableDAO {
	return timetableDao{db: db}
}

type timetableDao struct {
	db *gorm.DB
}

func (dao timetableDao) FetchByStudioIDAndDate(studioID uint, date time.Time) (dto.Timetable, error) {
	var records []Schedule
	if err := dao.db.Preload("Studio").Preload("Program").Preload("Performer").Where("strftime('%Y-%m-%d', start_at, 'localtime') = ? AND studio_id = ?", date.Format(time.DateOnly), studioID).Find(&records).Error; err != nil {
		return nil, err
	}

	results := make(dto.Schedules, len(records))
	for i, v := range records {
		results[i] = dto.Schedule{
			StudioName:    v.Studio.Name,
			ProgramName:   v.Program.Name,
			PerformerName: v.Performer.Name,
			StartAt:       v.StartAt,
			EndAt:         v.EndAt,
			HashID:        v.HashID,
		}
	}
	return results.ToTimeTable(), nil
}

func (dao timetableDao) FetchByPerformerIDAndDate(performerID uint, date time.Time) (dto.Timetable, error) {
	var records []Schedule
	if err := dao.db.Preload("Studio").Preload("Program").Preload("Performer").Where("strftime('%Y-%m-%d', start_at, 'localtime') = ? AND performer_id = ?", date.Format(time.DateOnly), performerID).Find(&records).Error; err != nil {
		return nil, err
	}

	results := make(dto.Schedules, len(records))
	for i, v := range records {
		results[i] = dto.Schedule{
			StudioName:    v.Studio.Name,
			ProgramName:   v.Program.Name,
			PerformerName: v.Performer.Name,
			StartAt:       v.StartAt,
			EndAt:         v.EndAt,
			HashID:        v.HashID,
		}
	}
	return results.ToTimeTable(), nil
}
