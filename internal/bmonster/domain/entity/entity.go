package entity

import (
	"time"
)

type Studio struct {
	ID   uint
	Name string
}

type Performer struct {
	ID   uint
	Name string
}

type Program struct {
	ID   uint
	Name string
}

type Schedule struct {
	ID          uint
	StudioID    uint
	ProgramID   uint
	PerformerID uint
	StartAt     time.Time
	EndAt       time.Time
	HashID      string
}
