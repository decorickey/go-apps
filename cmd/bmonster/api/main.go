package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/decorickey/go-apps/internal/bmonster/application/usecase"
	"github.com/decorickey/go-apps/internal/bmonster/presentation/handler"
	"github.com/decorickey/go-apps/internal/bmonster/presentation/openapi"
)

func main() {
	u := usecase.InitializeScrapingUsecase()

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

	h := handler.InitializeHandler()
	mux := http.NewServeMux()
	log.Println("starting ...")
	hh := openapi.HandlerFromMux(h, mux)

	port := ":8080"
	s := &http.Server{
		Addr:         port,
		Handler:      hh,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  30 * time.Second,
	}
	log.Printf("listen at http://localhost%s\n", port)
	log.Fatal(s.ListenAndServe())
}
