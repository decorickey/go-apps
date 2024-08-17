package command

import (
	"time"

	"github.com/decorickey/go-apps/internal/bmonster/domain/entity"
)

type StudioBody struct {
	Data struct {
		Studios struct {
			List []Studio `json:"list"`
		} `json:"studios"`
	} `json:"data"`
}

type ScheduleBody struct {
	Data struct {
		StudioLessons struct {
			Programs    []Program   `json:"programs"`
			Instructors []Performer `json:"instructors"`
			Items       []Schedule  `json:"items"`
		} `json:"studio_lessons"`
	} `json:"data"`
}

type Studio struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (s Studio) ToEntity() *entity.Studio {
	return &entity.Studio{ID: s.ID, Name: s.Name}
}

type Program struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (p Program) ToEntity() *entity.Program {
	return &entity.Program{ID: p.ID, Name: p.Name}
}

type Performer struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (p Performer) ToEntity() *entity.Performer {
	return &entity.Performer{ID: p.ID, Name: p.Name}
}

type Schedule struct {
	ID          uint      `json:"id"`
	StudioID    uint      `json:"studio_id"`
	PerformerID uint      `json:"instructor_id"`
	ProgramID   uint      `json:"program_id"`
	StartAt     time.Time `json:"start_at"`
	EndAt       time.Time `json:"end_at"`
	HashID      string    `json:"id_hash"`
}

func (s Schedule) ToEntity() *entity.Schedule {
	return &entity.Schedule{
		ID:          s.ID,
		StudioID:    s.StudioID,
		PerformerID: s.PerformerID,
		ProgramID:   s.ProgramID,
		StartAt:     s.StartAt,
		EndAt:       s.EndAt,
		HashID:      s.HashID,
	}
}
