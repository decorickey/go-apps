package entity

type PerformerRepository interface {
	FindAll() ([]Performer, error)
}

type ScheduleRepository interface {
	FindAll() ([]Schedule, error)
	FindFromToday() ([]Schedule, error)
	FindByPerformer(performer Performer) ([]Schedule, error)
	Save(schedules []Schedule) error
}
