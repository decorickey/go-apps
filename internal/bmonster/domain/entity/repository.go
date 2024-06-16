package entity

type PerformerRepository interface {
	List() ([]Performer, error)
}

type ScheduleRepository interface {
	List() ([]Schedule, error)
	FindByPerformer(performer Performer) ([]Schedule, error)
	Save(schedules []Schedule) error
}
