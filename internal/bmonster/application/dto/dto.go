package dto

import (
	"time"
)

var (
	// {"hh:mm:ss": idx}
	TimeTable map[string]int
	timeTable []string = []string{
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
)

func init() {
	TimeTable = make(map[string]int, len(timeTable))
	for i, v := range timeTable {
		TimeTable[v] = i
	}
}

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
	return TimeTable[s.TimeOnlyString()]
}

type SchedulesWithTimeTable []*Schedule
type Schedules []Schedule

func (s Schedules) WithTimeTable() SchedulesWithTimeTable {
	ss := make(SchedulesWithTimeTable, len(timeTable))
	for _, v := range s {
		ss[v.TimeTableIndex()] = &v
	}
	return ss
}

func (s Schedules) PerDate() []SchedulesPerDate {
	m := map[string]Schedules{}
	for _, v := range s {
		m[v.DateOnlyString()] = append(m[v.DateOnlyString()], v)
	}

	results := []SchedulesPerDate{}
	for k, v := range m {
		result := SchedulesPerDate{
			Date:      k,
			Schedules: v,
		}
		results = append(results, result)
	}

	return results
}

func (s Schedules) PerDateAndStudio() []SchedulesPerDateAndStudio {
	results := []SchedulesPerDateAndStudio{}
	for _, v := range s.PerDate() {
		results = append(results, v.AndStudio()...)
	}
	return results
}

func (s Schedules) PerDateAndPerformer() []SchedulesPerDateAndPerformer {
	results := []SchedulesPerDateAndPerformer{}
	for _, v := range s.PerDate() {
		results = append(results, v.AndPerformer()...)
	}
	return results
}

type SchedulesPerDate struct {
	Date      string    `json:"date"`
	Schedules Schedules `json:"schedules"`
}

func (s SchedulesPerDate) AndStudio() []SchedulesPerDateAndStudio {
	m := map[string]Schedules{}
	for _, v := range s.Schedules {
		m[v.StudioName] = append(m[v.StudioName], v)
	}

	results := []SchedulesPerDateAndStudio{}
	for k, v := range m {
		result := SchedulesPerDateAndStudio{
			Date:       s.Date,
			StudioName: k,
			Schedules:  v.WithTimeTable(),
		}
		results = append(results, result)
	}
	return results
}

func (s SchedulesPerDate) AndPerformer() []SchedulesPerDateAndPerformer {
	m := map[string]Schedules{}
	for _, v := range s.Schedules {
		m[v.PerformerName] = append(m[v.PerformerName], v)
	}

	results := []SchedulesPerDateAndPerformer{}
	for k, v := range m {
		result := SchedulesPerDateAndPerformer{
			Date:          s.Date,
			PerformerName: k,
			Schedules:     v.WithTimeTable(),
		}
		results = append(results, result)
	}
	return results
}

type SchedulesPerDateAndStudio struct {
	Date       string                 `json:"date"`
	StudioName string                 `json:"studio_name"`
	Schedules  SchedulesWithTimeTable `json:"schedules"`
}

type SchedulesPerDateAndPerformer struct {
	Date          string                 `json:"date"`
	PerformerName string                 `json:"performer_name"`
	Schedules     SchedulesWithTimeTable `json:"schedules"`
}
