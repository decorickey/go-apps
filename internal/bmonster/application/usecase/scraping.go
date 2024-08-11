package usecase

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/decorickey/go-apps/internal/bmonster/domain/entity"
	"github.com/decorickey/go-apps/pkg/timeutil"
)

const apiBaseUrl = "https://b-monster.hacomono.jp/api/master"

type ScrapingUsecase interface {
	FetchStudios() ([]entity.Studio, error)
	FetchSchedulesByStudios(studios []entity.Studio) (*studioLessons, error)
}

func NewScrapingUsecase(c *http.Client) ScrapingUsecase {
	baseUrl, _ := url.Parse(apiBaseUrl)
	return &scrapingUsecase{
		client:  c,
		baseUrl: baseUrl,
	}
}

type scrapingUsecase struct {
	client  *http.Client
	baseUrl *url.URL
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

	return body.Data.Studios.List, nil
}

func (u scrapingUsecase) FetchSchedulesByStudios(studios []entity.Studio) (*studioLessons, error) {
	weeks := []time.Time{
		timeutil.NowInJST(),
		timeutil.NowInJST().AddDate(0, 0, 7),
	}

	var results studioLessons
	for _, studio := range studios {
		for _, week := range weeks {
			lessons, err := u.fetchSchedulesByStudio(studio, week)
			if err != nil {
				return nil, err
			}
			results.Programs = append(results.Programs, lessons.Programs...)
			results.Instructors = append(results.Instructors, lessons.Instructors...)
			results.Items = append(results.Items, lessons.Items...)
		}
	}
	return &results, nil
}

func (u scrapingUsecase) fetchSchedulesByStudio(studio entity.Studio, dateFrom time.Time) (*studioLessons, error) {
	scheduleUrl := u.baseUrl.JoinPath("studio-lessons", "schedule")

	query := scheduleQuery{StudioID: studio.ID, DateFrom: dateFrom.Format(time.DateOnly)}
	q, err := json.Marshal(query)
	if err != nil {
		return nil, fmt.Errorf("failed to generate query params: %w", err)
	}

	v := url.Values{}
	v.Set("query", string(q))

	scheduleUrl.RawQuery = v.Encode()
	req, err := http.NewRequest(http.MethodGet, scheduleUrl.String(), nil)
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

	var body scheduleBody
	if err := json.Unmarshal(buf, &body); err != nil {
		return nil, fmt.Errorf("failed to parse response body: %w", err)
	}

	return &body.Data.StudioLessons, nil
}

type studioBody struct {
	Data struct {
		Studios struct {
			List []entity.Studio `json:"list"`
		} `json:"studios"`
	} `json:"data"`
}

type scheduleQuery struct {
	StudioID int    `json:"studio_id"`
	DateFrom string `json:"date_from"`
}

type scheduleBody struct {
	Data struct {
		StudioLessons studioLessons `json:"studio_lessons"`
	} `json:"data"`
}

type studioLessons struct {
	Programs    []entity.Program   `json:"programs"`
	Instructors []entity.Performer `json:"instructors"`
	Items       []entity.Schedule  `json:"items"`
}
