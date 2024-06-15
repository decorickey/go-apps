package usecase

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/decorickey/go-apps/pkg/timeutil"
)

type ScrapingUsecase struct {
	client *http.Client
}

func NewScrapingUsecase(client *http.Client) *ScrapingUsecase {
	return &ScrapingUsecase{
		client: client,
	}
}

func (u ScrapingUsecase) Performers() ([]PerformerQueryCommand, error) {
	res, err := u.client.Get("https://www.b-monster.jp/_inc_/instructors/inst_tpl?mode=filtering")
	if err != nil {
		return nil, fmt.Errorf("failed to http request: %w", err)
	}
	defer res.Body.Close()

	if status := res.StatusCode; status != http.StatusOK {
		return nil, fmt.Errorf("http status code: %d", status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse html: %w", err)
	}

	performers := make([]PerformerQueryCommand, 0)
	doc.Find("div.panel").Each(func(i int, s *goquery.Selection) {
		name := s.Find("h3.ofonts").Text()

		href, ok := s.Find("a.ofonts").Attr("href")
		if !ok {
			return
		}

		u, err := url.ParseRequestURI(href)
		if err != nil {
			return
		}

		ids, ok := u.Query()["instructor_id"]
		if !ok || len(ids) == 0 {
			return
		}

		id, err := strconv.Atoi(ids[0])
		if err != nil {
			return
		}

		performers = append(performers, PerformerQueryCommand{ID: id, Name: name})
	})
	return performers, nil
}

func (u ScrapingUsecase) SchedulesByPerformer(perfomerID int, performerName string) ([]ScheduleCommand, error) {
	date := time.Now().In(timeutil.JST)

	baseUrl, err := url.ParseRequestURI("https://www.b-monster.jp/reserve/")
	if err != nil {
		return nil, err
	}
	q := baseUrl.Query()
	q.Set("instructor_id", strconv.Itoa(perfomerID))
	baseUrl.RawQuery = q.Encode()
	q.Set("date", date.Format(time.DateOnly))

	req, err := http.NewRequest(http.MethodGet, baseUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := u.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to http request: %w", err)
	}
	defer res.Body.Close()

	if status := res.StatusCode; status != http.StatusOK {
		return nil, fmt.Errorf("http status code: %d", status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse html: %w", err)
	}

	commands := make([]ScheduleCommand, 0)
	days := doc.Find("#scroll-box .flex-no-wrap")
	days.Each(func(i int, s *goquery.Selection) {
		ttd := date.AddDate(0, 0, i)
		day := s.Find(".daily-panel li")
		day.Each(func(ii int, ss *goquery.Selection) {
			content := ss.Find(".panel-content")

			ttt := content.Find(".tt-time").Text()
			if ttt == "" {
				return
			}
			ttt = strings.Split(ttt, " ")[0]
			hour, err := strconv.Atoi(ttt[0:2])
			min, err := strconv.Atoi(ttt[3:5])
			if err != nil {
				return
			}

			tti := content.Find(".tt-instructor").Text()
			if tti == "" {
				return
			}

			ttm := content.Find(".tt-mode").Text()
			if ttm == "" {
				return
			}
			ttm = strings.ReplaceAll(strings.ReplaceAll(ttm, "\n", ""), " ", "")

			command := ScheduleCommand{
				Studio:     tti,
				StartYear:  ttd.Year(),
				StartMonth: ttd.Month(),
				StartDay:   ttd.Day(),
				StartHour:  hour,
				StartMin:   min,
				Performer:  performerName,
				Vol:        ttm,
			}
			commands = append(commands, command)
		})
	})
	return commands, nil
}
