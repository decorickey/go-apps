package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/decorickey/go-apps/internal/bmonster/application/usecase"
	"github.com/decorickey/go-apps/internal/bmonster/presentation/handler"
)

func main() {
	u := usecase.InitializeScrapingUsecase()
	h := handler.InitializeHandler()

	studios, err := u.FetchStudios()
	if err != nil {
		log.Fatal(fmt.Errorf("failed to fetch studios: %w", err))
	}

	performers, programs, schedules, err := u.FetchSchedulesByStudios(studios)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to fetch schedules: %w", err))
	}

	if err := u.Save(studios, performers, programs, schedules); err != nil {
		log.Fatal(fmt.Errorf("failed to save Entities: %w", err))
	}

	log.Println("starting ...")
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/bmonster/studios", h.FetchAllStudios)
	mux.HandleFunc("GET /api/bmonster/performers", h.FetchAllPerformers)
	mux.HandleFunc("GET /api/bmonster/studios/{id}/timetable", h.FetchTimetableByStudioID)
	mux.HandleFunc("GET /api/bmonster/performers/{id}/timetable", h.FetchTimetableByPerformerID)

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
