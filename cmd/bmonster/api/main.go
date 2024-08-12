package main

import (
	"log"
	"net/http"
	"time"

	"github.com/decorickey/go-apps/internal/bmonster/application/dao"
	"github.com/decorickey/go-apps/internal/bmonster/presentation/handler"
)

func main() {
	// setup
	var (
		studioDao    dao.StudioDAO
		performerDao dao.PerformerDAO
		scheduleDao  dao.ScheduleDAO
	)
	h := handler.NewHandler(studioDao, performerDao, scheduleDao)

	// start
	log.Println("starting ...")
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/bmonster/studios", h.FetchAllStudios)
	mux.HandleFunc("GET /api/bmonster/performers", h.FetchAllPerformers)

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
