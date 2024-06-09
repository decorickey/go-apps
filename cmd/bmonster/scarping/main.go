package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/decorickey/go-apps/internal/bmonster/application/usecase"
)

func main() {
	client := &http.Client{Timeout: 5 * time.Second}
	u := usecase.NewScrapingUsecase(client)

	performers := u.ScrapingPerformers()

	var schedules []usecase.ScheduleQuery
	for _, performer := range performers {
		gots := u.ScrapingSchedulesByPerformer(performer.ID, performer.Name)
		if gots != nil {
			for _, got := range gots {
				schedules = append(schedules, *got.ToQuery())
			}
		}
	}

	fmt.Println(schedules)
}
