package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/decorickey/go-apps/internal/bmonster/application/usecase"
)

func main() {
	c := &http.Client{Timeout: 5 * time.Second}
	u := usecase.NewScrapingUsecase(c)
	studios, err := u.FetchStudios()
	if err != nil {
		log.Fatal(fmt.Errorf("failed to fetch studios: %w", err))
	}

	lessons, err := u.FetchSchedulesByStudios(studios)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to fetch schedules: %w", err))
	}

	// TODO repository.Save()
	fmt.Println(lessons)
}
