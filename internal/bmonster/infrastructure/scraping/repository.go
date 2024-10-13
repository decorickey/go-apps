package scraping

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/decorickey/go-apps/internal/bmonster/application/command"
	"github.com/decorickey/go-apps/internal/bmonster/domain/entity"
	"github.com/decorickey/go-apps/internal/bmonster/domain/repository"
	"github.com/decorickey/go-apps/pkg/timeutil"
)

const apiBaseUrl = "https://b-monster.hacomono.jp/api/master"

type scheduleQuery struct {
	StudioID uint   `json:"studio_id"`
	DateFrom string `json:"date_from"`
}

func NewScrapingRepository() repository.ScrapingRepository {
	c := &http.Client{Timeout: 5 * time.Second}
	baseUrl, _ := url.Parse(apiBaseUrl)
	return scrapingRepository{
		client:  c,
		baseUrl: baseUrl,
	}
}

type scrapingRepository struct {
	client  *http.Client
	baseUrl *url.URL
}

func (repo scrapingRepository) FetchStudios() ([]entity.Studio, error) {
	studioUrl := repo.baseUrl.JoinPath("studios")

	req, err := http.NewRequest(http.MethodGet, studioUrl.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("generate http request: %w", err)
	}

	res, err := repo.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do http request: %w", err)
	}
	defer res.Body.Close()

	if status := res.StatusCode; status != http.StatusOK {
		return nil, fmt.Errorf("http status code: %d", status)
	}

	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body: %w", err)
	}

	var body command.StudioBody
	if err := json.Unmarshal(buf, &body); err != nil {
		return nil, fmt.Errorf("parse response body: %w", err)
	}

	result := make([]entity.Studio, len(body.Data.Studios.List))
	for i, v := range body.Data.Studios.List {
		result[i] = *v.ToEntity()
	}
	return result, nil
}

func (repo scrapingRepository) FetchSchedulesByStudios(studios []entity.Studio) ([]entity.Performer, []entity.Program, []entity.Schedule, error) {
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
			pers, pros, sches, err := repo.fetchSchedulesByStudio(studio, week)
			if err != nil {
				return nil, nil, nil, err
			}
			performers = append(performers, pers...)
			programs = append(programs, pros...)
			schedules = append(schedules, sches...)
		}
	}

	slices.SortFunc(performers, func(a, b entity.Performer) int { return int(a.ID) - int(b.ID) })
	performers = slices.CompactFunc(performers, func(a, b entity.Performer) bool { return a == b })
	slices.SortFunc(programs, func(a, b entity.Program) int { return int(a.ID) - int(b.ID) })
	programs = slices.CompactFunc(programs, func(a, b entity.Program) bool { return a == b })
	return performers, programs, schedules, nil
}

func (repo scrapingRepository) fetchSchedulesByStudio(studio entity.Studio, dateFrom time.Time) ([]entity.Performer, []entity.Program, []entity.Schedule, error) {
	scheduleUrl := repo.baseUrl.JoinPath("studio-lessons", "schedule")

	query := scheduleQuery{StudioID: studio.ID, DateFrom: dateFrom.Format(time.DateOnly)}
	q, err := json.Marshal(query)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("generate query params: %w", err)
	}

	v := url.Values{}
	v.Set("query", string(q))

	scheduleUrl.RawQuery = v.Encode()
	req, err := http.NewRequest(http.MethodGet, scheduleUrl.String(), nil)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("generate http request: %w", err)
	}

	res, err := repo.client.Do(req)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("do http request: %w", err)
	}
	defer res.Body.Close()

	if status := res.StatusCode; status != http.StatusOK {
		return nil, nil, nil, fmt.Errorf("http status code: %d", status)
	}

	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("read response body: %w", err)
	}

	var body command.ScheduleBody
	if err := json.Unmarshal(buf, &body); err != nil {
		return nil, nil, nil, fmt.Errorf("parse response body: %w", err)
	}

	var (
		performers []entity.Performer
		programs   []entity.Program
		schedules  []entity.Schedule
	)
	for _, v := range body.Data.StudioLessons.Instructors {
		performers = append(performers, *v.ToEntity())
	}
	for _, v := range body.Data.StudioLessons.Programs {
		programs = append(programs, *v.ToEntity())
	}
	for _, v := range body.Data.StudioLessons.Items {
		schedules = append(schedules, *v.ToEntity())
	}
	return performers, programs, schedules, nil
}
