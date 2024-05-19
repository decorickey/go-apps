package entity_test

import (
	"testing"

	"github.com/decorickey/go-apps/bmonster/domain/entity"
	"github.com/decorickey/go-apps/lib/timeutil"
)

func TestNewPerformer(t *testing.T) {
	data := []struct {
		name      string
		isSuccess bool
	}{
		{
			name:      "",
			isSuccess: false,
		},
		{
			name:      "Bob",
			isSuccess: true,
		},
	}

	for _, d := range data {
		_, err := entity.NewPerformer(d.name)

		if err != nil && d.isSuccess {
			t.Error(err)
		}
	}
}

func TestNewProgram(t *testing.T) {
	performer, _ := entity.NewPerformer("Bob")
	data := []struct {
		vol       string
		isSuccess bool
	}{
		{
			vol:       "",
			isSuccess: false,
		},
		{
			vol:       "MIX1",
			isSuccess: true,
		},
	}
	for _, d := range data {
		_, err := entity.NewProgram(*performer, d.vol)

		if err != nil && d.isSuccess {
			t.Error(err)
		}
	}
}

func TestNewSchedule(t *testing.T) {
	performer, _ := entity.NewPerformer("Bob")
	program, _ := entity.NewProgram(*performer, "MIX1")
	now := timeutil.NowInJST()
	data := []struct {
		studio    string
		isSuccess bool
	}{
		{
			studio:    "",
			isSuccess: false,
		},
		{
			studio:    "GINZA",
			isSuccess: true,
		},
	}

	for _, d := range data {
		_, err := entity.NewSchedule(*program, d.studio, now)

		if err != nil && d.isSuccess {
			t.Error(err)
		}
	}
}
