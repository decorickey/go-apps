package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/decorickey/go-apps/internal/bmonster/application/usecase"
	"github.com/decorickey/go-apps/internal/bmonster/infrastructure/sql"
	"github.com/decorickey/go-apps/internal/bmonster/presentation/handler"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	db, err := gorm.Open(
		sqlite.Open("file::memory:?cache=shared"),
		&gorm.Config{
			Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
				SlowThreshold:             200 * time.Millisecond,
				LogLevel:                  logger.Info,
				IgnoreRecordNotFoundError: false,
				Colorful:                  true,
			})},
	)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to connect database: %w", err))
	}
	db.AutoMigrate(&sql.Studio{}, &sql.Performer{}, &sql.Program{}, &sql.Schedule{})

	studioRepo := sql.NewStudioRepository(db)
	performerRepo := sql.NewPerformerRepository(db)
	programRepo := sql.NewProgramRepository(db)
	scheduleRepo := sql.NewScheduleRepository(db)
	u := usecase.NewScrapingUsecase(studioRepo, performerRepo, programRepo, scheduleRepo)

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

	studioDao := sql.NewStudioDao(db)
	performerDao := sql.NewPerformerDao(db)
	timetableDao := sql.NewTimetableDao(db)
	h := handler.NewHandler(studioDao, performerDao, timetableDao)

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
