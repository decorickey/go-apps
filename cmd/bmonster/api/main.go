package main

import (
	"log"
	"net/http"
	"time"
)

type BmonsterHandler struct{}

func (h *BmonsterHandler) Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ping"))
}

func main() {
	h := &BmonsterHandler{}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", h.Ping)

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
