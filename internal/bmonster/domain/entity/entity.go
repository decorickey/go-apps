package entity

import (
	"errors"
	"time"
)

var ErrValidation = errors.New("validation error")

type Performer struct {
	ID   int
	Name string
}

func NewPerformer(id int, name string) (*Performer, error) {
	if name == "" {
		return nil, ErrValidation
	}
	return &Performer{ID: id, Name: name}, nil
}

type Program struct {
	Performer Performer
	Vol       string
}

func NewProgram(performer Performer, vol string) (*Program, error) {
	if vol == "" {
		return nil, ErrValidation
	}
	return &Program{Performer: performer, Vol: vol}, nil
}

type Schedule struct {
	Program
	Studio  string
	StartAt time.Time
}

func NewSchedule(program Program, studio string, startAt time.Time) (*Schedule, error) {
	if studio == "" {
		return nil, ErrValidation
	}

	return &Schedule{
		Program: program,
		Studio:  studio,
		StartAt: startAt,
	}, nil
}
