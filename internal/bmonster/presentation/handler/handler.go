package handler

import (
	"encoding/json"
	"net/http"

	"github.com/decorickey/go-apps/internal/bmonster/application/usecase"
)

type BmonsterHandler interface {
	AllPerformers(w http.ResponseWriter, r *http.Request)
	AllSchedules(w http.ResponseWriter, r *http.Request)
	SchedulesByPerformer(w http.ResponseWriter, r *http.Request)
}

func NewBmonsterHandler(
	performerQueryUsecase usecase.PerformerQueryUsecase,
	scheduleQueryUsecase usecase.ScheduleQueryUsecase,
) BmonsterHandler {
	return &bmonsterHandler{
		performerQueryUsecase: performerQueryUsecase,
		scheduleQueryUsecase:  scheduleQueryUsecase,
	}
}

type bmonsterHandler struct {
	performerQueryUsecase usecase.PerformerQueryUsecase
	scheduleQueryUsecase  usecase.ScheduleQueryUsecase
}

func (h bmonsterHandler) AllPerformers(w http.ResponseWriter, r *http.Request) {
	performers := h.performerQueryUsecase.All()
	b, err := json.Marshal(performers)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(b)
}

func (h bmonsterHandler) AllSchedules(w http.ResponseWriter, r *http.Request) {
	schedules, err := h.scheduleQueryUsecase.All()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(schedules)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(b)
}

func (h bmonsterHandler) SchedulesByPerformer(w http.ResponseWriter, r *http.Request) {}
