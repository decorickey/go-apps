package main

import (
	"log"
	"net/http"
	"time"

	"github.com/decorickey/go-apps/internal/bmonster/application/usecase"
	"github.com/decorickey/go-apps/internal/bmonster/presentation/handler"
)

func main() {
	c := &http.Client{Timeout: 5 * time.Second}
	su := usecase.NewScrapingUsecase(c)
	pqu := usecase.NewMockPerformerQueryUsecase(su)
	squ := usecase.NewMockScheduleQueryUsecase(su)
	h := handler.NewBmonsterHandler(pqu, squ)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/bmonster/performers", h.AllPerformers)
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
