package handler

import (
	"encoding/json"
	"net/http"

	"github.com/decorickey/go-apps/internal/bmonster/application/dao"
)

type Handler interface {
	FetchAllStudios(w http.ResponseWriter, r *http.Request)
	FetchAllPerformers(w http.ResponseWriter, r *http.Request)
}

func NewHandler(
	studioDao dao.StudioDAO,
	performerDao dao.PerformerDAO,
	scheduleDao dao.ScheduleDAO,
) Handler {
	return &handler{
		studioDao:    studioDao,
		performerDao: performerDao,
		scheduleDao:  scheduleDao,
	}
}

type handler struct {
	studioDao    dao.StudioDAO
	performerDao dao.PerformerDAO
	scheduleDao  dao.ScheduleDAO
}

func (h handler) FetchAllStudios(w http.ResponseWriter, r *http.Request) {
	studios, err := h.studioDao.FetchAll()
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
	performers, err := h.performerDao.FetchAll()
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
