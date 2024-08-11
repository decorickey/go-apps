package handler

import (
	"encoding/json"
	"net/http"

	"github.com/decorickey/go-apps/internal/bmonster/application/usecase"
)

type Handler interface {
	FetchAllStudios(w http.ResponseWriter, r *http.Request)
	FetchAllPerformers(w http.ResponseWriter, r *http.Request)
	FetchAllSchedules(w http.ResponseWriter, r *http.Request)
	FetchAllSchedulesPerStudio(w http.ResponseWriter, r *http.Request)
	FetchAllSchedulesPerPerformer(w http.ResponseWriter, r *http.Request)
}

func NewHandler(
	studioQueryUsecase usecase.StudioQueryUsecase,
	performerQueryUsecase usecase.PerformerQueryUsecase,
	scheduleQueryUsecase usecase.ScheduleQueryUsecase,
) Handler {
	return &handler{
		studioQueryUsecase:    studioQueryUsecase,
		performerQueryUsecase: performerQueryUsecase,
		scheduleQueryUsecase:  scheduleQueryUsecase,
	}
}

type handler struct {
	studioQueryUsecase    usecase.StudioQueryUsecase
	performerQueryUsecase usecase.PerformerQueryUsecase
	scheduleQueryUsecase  usecase.ScheduleQueryUsecase
}

func (h handler) FetchAllStudios(w http.ResponseWriter, r *http.Request) {
	studios, err := h.studioQueryUsecase.FetchAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(studios)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(b)
}

func (h handler) FetchAllPerformers(w http.ResponseWriter, r *http.Request) {
	performers, err := h.performerQueryUsecase.FetchAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(performers)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(b)
}

func (h handler) FetchAllSchedules(w http.ResponseWriter, r *http.Request) {
	schedules, err := h.scheduleQueryUsecase.FetchAll()
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

func (h handler) FetchAllSchedulesPerStudio(w http.ResponseWriter, r *http.Request) {
	schedules, err := h.scheduleQueryUsecase.FetchAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(schedules.PerDateAndStudio())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(b)
}

func (h handler) FetchAllSchedulesPerPerformer(w http.ResponseWriter, r *http.Request) {
	schedules, err := h.scheduleQueryUsecase.FetchAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(schedules.PerDateAndPerformer())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(b)
}
