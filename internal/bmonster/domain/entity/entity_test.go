package entity_test

import (
	"testing"

	"github.com/decorickey/go-apps/internal/bmonster/domain/entity"
	"github.com/decorickey/go-apps/pkg/timeutil"
)

func TestNewPerformer(t *testing.T) {
	tests := []struct {
		id        int
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

	for _, tt := range tests {
		_, err := entity.NewPerformer(tt.id, tt.name)

		if err != nil && tt.isSuccess {
			t.Error(err)
		}
	}
}

func TestNewProgram(t *testing.T) {
	performer, _ := entity.NewPerformer(0, "Bob")
	tests := []struct {
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
	for _, tt := range tests {
		_, err := entity.NewProgram(*performer, tt.vol)

		if err != nil && tt.isSuccess {
			t.Error(err)
		}
	}
}

func TestNewSchedule(t *testing.T) {
	performer, _ := entity.NewPerformer(0, "Bob")
	program, _ := entity.NewProgram(*performer, "MIX1")
	now := timeutil.NowInJST()
	tests := []struct {
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

	for _, tt := range tests {
		_, err := entity.NewSchedule(*program, tt.studio, now)

		if err != nil && tt.isSuccess {
			t.Error(err)
		}
	}
}
