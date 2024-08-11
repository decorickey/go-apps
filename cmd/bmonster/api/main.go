package main

import (
	"log"
	"net/http"
	"time"

	"github.com/decorickey/go-apps/internal/bmonster/application/dto"
	"github.com/decorickey/go-apps/internal/bmonster/application/usecase"
	"github.com/decorickey/go-apps/internal/bmonster/domain/entity"
	"github.com/decorickey/go-apps/internal/bmonster/presentation/handler"
)

func main() {
	// setup
	studioRepo, performerRepo, programRepo, scheduleRepo := setupLocalRepositories()
	studioQU := usecase.NewStudioQueryUsecase(studioRepo)
	performerQU := usecase.NewPerformerQueryUsecase(performerRepo)
	scheduleQU := newScheduleQueryUsecase(studioRepo, performerRepo, programRepo, scheduleRepo)
	h := handler.NewHandler(studioQU, performerQU, scheduleQU)

	// start
	log.Println("starting ...")
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/bmonster/schedules", h.FetchAllSchedules)
	mux.HandleFunc("GET /api/bmonster/performers", h.FetchAllPerformers)
	mux.HandleFunc("GET /api/bmonster/performers/schedules", h.FetchAllSchedulesPerPerformer)
	mux.HandleFunc("GET /api/bmonster/studios", h.FetchAllStudios)
	mux.HandleFunc("GET /api/bmonster/studios/schedules", h.FetchAllSchedulesPerStudio)

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

func newLocalRepository[T entity.Studio | entity.Performer | entity.Program | entity.Schedule]() *localRepository[T] {
	records := make([]T, 0)
	return &localRepository[T]{records: records}
}

type localRepository[T entity.Studio | entity.Performer | entity.Program | entity.Schedule] struct {
	records []T
}

func (repo *localRepository[T]) Save(entities []T) error {
	repo.records = append(repo.records, entities...)
	return nil
}

func (repo *localRepository[T]) FindAll() ([]T, error) {
	return repo.records, nil
}

func setupLocalRepositories() (
	*localRepository[entity.Studio],
	*localRepository[entity.Performer],
	*localRepository[entity.Program],
	*localRepository[entity.Schedule],
) {
	c := &http.Client{Timeout: 5 * time.Second}
	u := usecase.NewScrapingUsecase(c)
	studios, err := u.FetchStudios()
	if err != nil {
		log.Fatal(err)
	}
	lessons, err := u.FetchSchedulesByStudios(studios)
	if err != nil {
		log.Fatal(err)
	}
	studioRepo := newLocalRepository[entity.Studio]()
	studioRepo.Save(studios)
	performerRepo := newLocalRepository[entity.Performer]()
	performerRepo.Save(lessons.Instructors)
	programRepo := newLocalRepository[entity.Program]()
	programRepo.Save(lessons.Programs)
	scheduleRepo := newLocalRepository[entity.Schedule]()
	scheduleRepo.Save(lessons.Items)
	return studioRepo, performerRepo, programRepo, scheduleRepo
}

func newScheduleQueryUsecase(
	studioRepo *localRepository[entity.Studio],
	performerRepo *localRepository[entity.Performer],
	programRepo *localRepository[entity.Program],
	scheduleRepo *localRepository[entity.Schedule],
) usecase.ScheduleQueryUsecase {
	return scheduleQueryUsecase{
		studioRepo:    studioRepo,
		performerRepo: performerRepo,
		programRepo:   programRepo,
		scheduleRepo:  scheduleRepo,
	}
}

type scheduleQueryUsecase struct {
	studioRepo    *localRepository[entity.Studio]
	performerRepo *localRepository[entity.Performer]
	programRepo   *localRepository[entity.Program]
	scheduleRepo  *localRepository[entity.Schedule]
}

func (qu scheduleQueryUsecase) FetchAll() (dto.Schedules, error) {
	studios := qu.studioRepo.records
	studioMap := make(map[int]entity.Studio, len(studios))
	for _, v := range studios {
		studioMap[v.ID] = v
	}

	performers := qu.performerRepo.records
	performerMap := make(map[int]entity.Performer, len(performers))
	for _, v := range performers {
		performerMap[v.ID] = v
	}

	programs := qu.programRepo.records
	programMap := make(map[int]entity.Program, len(programs))
	for _, v := range programs {
		programMap[v.ID] = v
	}

	schedules := qu.scheduleRepo.records
	results := make(dto.Schedules, len(schedules))
	for i, v := range schedules {
		results[i] = dto.Schedule{
			StartAt:       v.StartAt,
			EndAt:         v.EndAt,
			HashID:        v.HashID,
			StudioName:    studioMap[v.StudioID].Name,
			PerformerName: performerMap[v.PerformerID].Name,
			ProgramName:   programMap[v.ProgramID].Name,
		}
	}

	return results, nil
}
