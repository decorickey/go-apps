package entity

type PerformerRepository interface {
	All() ([]Performer, error)
}

type ScheduleRepository interface {
	All() ([]Schedule, error)
	FilterByPerformer(performer Performer) ([]Schedule, error)
	Save(schedules []Schedule) error
}
