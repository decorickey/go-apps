package main

import (
	"log"
	"net/http"
	"time"

	"github.com/decorickey/go-apps/internal/bmonster/application/usecase"
	"github.com/decorickey/go-apps/internal/bmonster/domain/entity"
	"github.com/decorickey/go-apps/internal/bmonster/presentation/handler"
)

type mockPerformerRepository struct {
	performers []entity.Performer
}

func (repo mockPerformerRepository) All() ([]entity.Performer, error) {
	return repo.performers, nil
}

type mockScheduleRepository struct {
	schedules []entity.Schedule
}

func (repo mockScheduleRepository) All() ([]entity.Schedule, error) {
	return repo.schedules, nil
}

func (repo mockScheduleRepository) FilterByPerformer(performer entity.Performer) ([]entity.Schedule, error) {
	return nil, nil
}

func (repo mockScheduleRepository) Save(schedules []entity.Schedule) error { return nil }

func main() {
	log.Println("starting ...")
	c := &http.Client{Timeout: 5 * time.Second}
	su := usecase.NewScrapingUsecase(c)

	performerQueries, _ := su.Performers()
	performers := make([]entity.Performer, 0)
	for _, query := range performerQueries {
		performer, _ := query.ToEntity()
		performers = append(performers, *performer)
	}
	performerRepository := &mockPerformerRepository{performers: performers}

	schedules := make([]entity.Schedule, 0)
	for _, performer := range performers {
		gots, err := su.SchedulesByPerformer(performer.ID, performer.Name)
		if err != nil {
			continue
		}
		for _, got := range gots {
			program, _ := entity.NewProgram(performer, got.Vol)
			schedule, _ := entity.NewSchedule(*program, got.Studio, got.StartAt())
			schedules = append(schedules, *schedule)
		}
	}
	scheduleRepository := &mockScheduleRepository{schedules: schedules}

	pqu := usecase.NewPerformerQueryUsecase(performerRepository)
	squ := usecase.NewScheduleQueryUsecase(scheduleRepository)
	h := handler.NewBmonsterHandler(pqu, squ)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/bmonster/performers", h.AllPerformers)
	mux.HandleFunc("GET /api/bmonster/performers/{ID}/schedules", h.AllSchedules)
	mux.HandleFunc("GET /api/bmonster/schedules", h.AllSchedules)

	port := ":8080"
	s := &http.Server{
		Addr:         port,
		Handler:      mux,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  30 * time.Second,
	}
	log.Printf("listen at http://localhost%s\n", port)
	log.Fatal(s.ListenAndServe())
}
