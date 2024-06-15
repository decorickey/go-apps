package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/decorickey/go-apps/internal/bmonster/application/usecase"
)

func main() {
	client := &http.Client{Timeout: 5 * time.Second}
	su := usecase.NewScrapingUsecase(client)

	performers, err := su.Performers()
	if err != nil {
		slog.Error("failed to scraping performers", err)
	}

	var schedules []usecase.ScheduleQuery
	for _, performer := range performers {
		gots, err := su.SchedulesByPerformer(performer.ID, performer.Name)
		if err != nil {
			slog.Error("unexpected error", err)
			continue
		}
		for _, got := range gots {
			schedule := got.ToQuery()
			schedules = append(schedules, *schedule)
		}
	}

	fmt.Println(schedules)
}
