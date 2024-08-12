package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/decorickey/go-apps/internal/bmonster/application/usecase"
	"github.com/decorickey/go-apps/internal/bmonster/domain/repository"
)

func main() {
	c := &http.Client{Timeout: 5 * time.Second}

	var (
		studioRepo    repository.StudioRepository
		performerRepo repository.PerformerRepository
		programRepo   repository.ProgramRepository
		scheduleRepo  repository.ScheduleRepository
	)
	u := usecase.NewScrapingUsecase(c, studioRepo, performerRepo, programRepo, scheduleRepo)

	studios, err := u.FetchStudios()
	if err != nil {
		log.Fatal(fmt.Errorf("failed to fetch studios: %w", err))
	}

	performers, programs, schedules, err := u.FetchSchedulesByStudios(studios)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to fetch schedules: %w", err))
	}

	fmt.Println(studios, performers, programs, schedules)
}
