package entity

import (
	"time"
)

type Studio struct {
	ID   int
	Name string
}

type Performer struct {
	ID   int
	Name string
}

type Program struct {
	ID   int
	Name string
}

type Schedule struct {
	ID          int
	StudioID    int
	ProgramID   int
	PerformerID int
	StartAt     time.Time
	EndAt       time.Time
	HashID      string
}
