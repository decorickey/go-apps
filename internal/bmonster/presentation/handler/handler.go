package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/decorickey/go-apps/internal/bmonster/application/dao"
	"github.com/decorickey/go-apps/pkg/timeutil"
)

type Handler interface {
	FetchAllStudios(http.ResponseWriter, *http.Request)
	FetchAllPerformers(http.ResponseWriter, *http.Request)
	FetchTimetableByStudioID(http.ResponseWriter, *http.Request)
	FetchTimetableByPerformerID(http.ResponseWriter, *http.Request)
}

func NewHandler(
	studioDao dao.StudioDAO,
	performerDao dao.PerformerDAO,
	timetableDao dao.TimetableDAO,
) Handler {
	return &handler{
		studioDao:    studioDao,
		performerDao: performerDao,
		timetableDao: timetableDao,
	}
}

type handler struct {
	studioDao    dao.StudioDAO
	performerDao dao.PerformerDAO
	timetableDao dao.TimetableDAO
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

func (h handler) FetchTimetableByStudioID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	studioID, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	date := timeutil.NowInJST()
	if q := r.URL.Query().Get("date"); q != "" {
		if v, err := time.Parse(time.DateOnly, q); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		} else {
			date = v
		}
	}

	timetable, err := h.timetableDao.FetchByStudioIDAndDate(uint(studioID), date)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(timetable)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(b)
}

func (h handler) FetchTimetableByPerformerID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	performerID, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	date := timeutil.NowInJST()
	if q := r.URL.Query().Get("date"); q != "" {
		if v, err := time.Parse(time.DateOnly, q); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		} else {
			date = v
		}
	}

	timetable, err := h.timetableDao.FetchByPerformerIDAndDate(uint(performerID), date)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(timetable)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(b)
}
