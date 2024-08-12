package dto

import (
	"time"
)

var (
	timeTableList []string = []string{
		"06:50:00",
		"08:10:00",
		"08:15:00",
		"09:30:00",
		"09:50:00",
		"10:50:00",
		"11:25:00",
		"12:10:00",
		"12:50:00",
		"13:30:00",
		"14:15:00",
		"14:50:00",
		"15:40:00",
		"16:10:00",
		"17:05:00",
		"17:30:00",
		"18:30:00",
		"18:50:00",
		"19:55:00",
		"20:05:00",
		"21:20:00",
	}
	// {"hh:mm:ss": idx}
	timeTableMap map[string]int
)

func init() {
	timeTableMap = make(map[string]int, len(timeTableList))
	for i, v := range timeTableList {
		timeTableMap[v] = i
	}
}

type Studio struct {
	ID   int
	Name string
}

type Performer struct {
	ID   int
	Name string
}

type TimeTable []*Schedule

type Schedule struct {
	StudioName    string    `json:"studio_name"`
	ProgramName   string    `json:"program_name"`
	PerformerName string    `json:"performer_name"`
	StartAt       time.Time `json:"start_at"`
	EndAt         time.Time `json:"end_at"`
	HashID        string    `json:"id_hash"`
}

func (s Schedule) DateOnlyString() string {
	return s.StartAt.Format(time.DateOnly)
}

func (s Schedule) TimeOnlyString() string {
	return s.StartAt.Format(time.TimeOnly)
}

func (s Schedule) TimeTableIndex() int {
	return timeTableMap[s.TimeOnlyString()]
}

type Schedules []Schedule

func (s Schedules) ToTimeTable() TimeTable {
	if len(s) > len(timeTableList) {
		return nil
	}

	ss := make(TimeTable, len(timeTableList))
	for _, v := range s {
		ss[v.TimeTableIndex()] = &v
	}
	return ss
}
