package entity

import (
	"time"
)

type Studio struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Program struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Performer struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Schedule struct {
	ID          int       `json:"id"`
	StudioID    int       `json:"studio_id"`
	ProgramID   int       `json:"program_id"`
	PerformerID int       `json:"instructor_id"`
	StartAt     time.Time `json:"start_at"`
	EndAt       time.Time `json:"end_at"`
	HashID      string    `json:"id_hash"`
}
