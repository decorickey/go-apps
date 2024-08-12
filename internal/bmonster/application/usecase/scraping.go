package usecase

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/decorickey/go-apps/internal/bmonster/domain/entity"
	"github.com/decorickey/go-apps/internal/bmonster/domain/repository"
	"github.com/decorickey/go-apps/pkg/timeutil"
)

const apiBaseUrl = "https://b-monster.hacomono.jp/api/master"

type schedule struct {
	ID          int       `json:"id"`
	StudioID    int       `json:"studio_id"`
	PerformerID int       `json:"instructor_id"`
	ProgramID   int       `json:"program_id"`
	StartAt     time.Time `json:"start_at"`
	EndAt       time.Time `json:"end_at"`
	HashID      string    `json:"id_hash"`
}

func (s schedule) toEntity() *entity.Schedule {
	return &entity.Schedule{
		ID:          s.ID,
		StudioID:    s.StudioID,
		PerformerID: s.PerformerID,
		ProgramID:   s.ProgramID,
		StartAt:     s.StartAt,
		EndAt:       s.EndAt,
		HashID:      s.HashID,
	}
}

type studioBody struct {
	Data struct {
		Studios struct {
			List []studio `json:"list"`
		} `json:"studios"`
	} `json:"data"`
}

type scheduleQuery struct {
	StudioID int    `json:"studio_id"`
	DateFrom string `json:"date_from"`
}

type scheduleBody struct {
	Data struct {
		StudioLessons struct {
			Programs    []program   `json:"programs"`
			Instructors []performer `json:"instructors"`
			Items       []schedule  `json:"items"`
		} `json:"studio_lessons"`
	} `json:"data"`
}

type studio struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (s studio) toEntity() *entity.Studio {
	return &entity.Studio{ID: s.ID, Name: s.Name}
}

type program struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (p program) toEntity() *entity.Program {
	return &entity.Program{ID: p.ID, Name: p.Name}
}

type performer struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ScrapingUsecase interface {
	FetchStudios() ([]entity.Studio, error)
	FetchSchedulesByStudios([]entity.Studio) ([]entity.Performer, []entity.Program, []entity.Schedule, error)
	Save([]entity.Studio, []entity.Performer, []entity.Program, []entity.Schedule) error
}

func NewScrapingUsecase(
	c *http.Client,
	studioRepo repository.StudioRepository,
	performerRepo repository.PerformerRepository,
	programRepo repository.ProgramRepository,
	scheduleRepo repository.ScheduleRepository,
) ScrapingUsecase {
	baseUrl, _ := url.Parse(apiBaseUrl)
	return &scrapingUsecase{
		client:        c,
		baseUrl:       baseUrl,
		studioRepo:    studioRepo,
		performerRepo: performerRepo,
		programRepo:   programRepo,
		scheduleRepo:  scheduleRepo,
	}
}

type scrapingUsecase struct {
	client        *http.Client
	baseUrl       *url.URL
	studioRepo    repository.StudioRepository
	performerRepo repository.PerformerRepository
	programRepo   repository.ProgramRepository
	scheduleRepo  repository.ScheduleRepository
}

func (u scrapingUsecase) FetchStudios() ([]entity.Studio, error) {
	studioUrl := u.baseUrl.JoinPath("studios")

	req, err := http.NewRequest(http.MethodGet, studioUrl.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to generate http request: %w", err)
	}

	res, err := u.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to do http request: %w", err)
	}
	defer res.Body.Close()

	if status := res.StatusCode; status != http.StatusOK {
		return nil, fmt.Errorf("http status code: %d", status)
	}

	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var body studioBody
	if err := json.Unmarshal(buf, &body); err != nil {
		return nil, fmt.Errorf("failed to parse response body: %w", err)
	}

	result := make([]entity.Studio, len(body.Data.Studios.List))
	for i, v := range body.Data.Studios.List {
		result[i] = *v.toEntity()
	}
	return result, nil
}

func (u scrapingUsecase) FetchSchedulesByStudios(studios []entity.Studio) ([]entity.Performer, []entity.Program, []entity.Schedule, error) {
	weeks := []time.Time{
		timeutil.NowInJST(),
		timeutil.NowInJST().AddDate(0, 0, 7),
	}

	var (
		performers []entity.Performer
		programs   []entity.Program
		schedules  []entity.Schedule
	)
	for _, studio := range studios {
		for _, week := range weeks {
			pers, pros, sches, err := u.fetchSchedulesByStudio(studio, week)
			if err != nil {
				return nil, nil, nil, err
			}
			performers = append(performers, pers...)
			programs = append(programs, pros...)
			schedules = append(schedules, sches...)
		}
	}
	return performers, programs, schedules, nil
}

func (u scrapingUsecase) fetchSchedulesByStudio(studio entity.Studio, dateFrom time.Time) ([]entity.Performer, []entity.Program, []entity.Schedule, error) {
	scheduleUrl := u.baseUrl.JoinPath("studio-lessons", "schedule")

	query := scheduleQuery{StudioID: studio.ID, DateFrom: dateFrom.Format(time.DateOnly)}
	q, err := json.Marshal(query)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to generate query params: %w", err)
	}

	v := url.Values{}
	v.Set("query", string(q))

	scheduleUrl.RawQuery = v.Encode()
	req, err := http.NewRequest(http.MethodGet, scheduleUrl.String(), nil)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to generate http request: %w", err)
	}

	res, err := u.client.Do(req)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to do http request: %w", err)
	}
	defer res.Body.Close()

	if status := res.StatusCode; status != http.StatusOK {
		return nil, nil, nil, fmt.Errorf("http status code: %d", status)
	}

	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var body scheduleBody
	if err := json.Unmarshal(buf, &body); err != nil {
		return nil, nil, nil, fmt.Errorf("failed to parse response body: %w", err)
	}

	var (
		performers []entity.Performer
		programs   []entity.Program
		schedules  []entity.Schedule
	)
	for _, v := range body.Data.StudioLessons.Instructors {
		performers = append(performers, *v.toEntity())
	}
	for _, v := range body.Data.StudioLessons.Programs {
		programs = append(programs, *v.toEntity())
	}
	for _, v := range body.Data.StudioLessons.Items {
		schedules = append(schedules, *v.toEntity())
	}
	return performers, programs, schedules, nil
}

func (p performer) toEntity() *entity.Performer {
	return &entity.Performer{ID: p.ID, Name: p.Name}
}

func (u scrapingUsecase) Save(studios []entity.Studio, performers []entity.Performer, programs []entity.Program, schedules []entity.Schedule) error {
	err := u.studioRepo.Save(studios)
	if err != nil {
		return fmt.Errorf("failed to save Studios: %w", err)
	}

	err = u.performerRepo.Save(performers)
	if err != nil {
		return fmt.Errorf("failed to save Performers: %w", err)
	}

	err = u.programRepo.Save(programs)
	if err != nil {
		return fmt.Errorf("failed to save Programs: %w", err)
	}

	err = u.scheduleRepo.Save(schedules)
	if err != nil {
		return fmt.Errorf("failed to save Schedules: %w", err)
	}

	return nil
}
