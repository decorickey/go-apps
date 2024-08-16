package sql

import (
	"time"
)

type Studio struct {
	ID        uint `gorm:"primaryKey;autoIncrement:false"`
	Name      string
	Schedules []Schedule
}

type Performer struct {
	ID        uint `gorm:"primaryKey;autoIncrement:false"`
	Name      string
	Schedules []Schedule
}

type Program struct {
	ID        uint `gorm:"primaryKey;autoIncrement:false"`
	Name      string
	Schedules []Schedule
}

type Schedule struct {
	ID          uint `gorm:"primaryKey;autoIncrement:false"`
	StudioID    uint
	Studio      Studio
	ProgramID   uint
	Program     Program
	PerformerID uint
	Performer   Performer
	StartAt     time.Time
	EndAt       time.Time
	HashID      string
}
