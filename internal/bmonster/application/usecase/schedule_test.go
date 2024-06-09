package usecase_test

import (
	"errors"
	"testing"
	"time"

	"github.com/decorickey/go-apps/internal/bmonster/application/usecase"
	"github.com/decorickey/go-apps/internal/bmonster/domain/entity"
	"go.uber.org/mock/gomock"
)

func TestScheduleQueryUsecase_FetchAll(t *testing.T) {
	ctrl := gomock.NewController(t)

	Performer1 := entity.Performer{
		Name: "Performer1",
	}
	Program1 := entity.Program{
		Performer: Performer1,
		Vol:       "MIX1",
	}
	Performer2 := entity.Performer{
		Name: "Performer2",
	}
	Program2 := entity.Program{
		Performer: Performer2,
		Vol:       "MIX2",
	}
	studio := "GINZA"
	startAt := time.Now()
	schedule1 := entity.Schedule{
		Studio:  studio,
		StartAt: startAt,
		Program: Program1,
	}
	schedule2 := entity.Schedule{
		Studio:  studio,
		StartAt: startAt,
		Program: Program2,
	}
	tests := []struct {
		name             string
		mockScheduleRepo func(scheduleRepo *entity.MockScheduleRepository)
		want             []usecase.ScheduleQuery
		wantErr          bool
	}{
		{
			name: "failed to find schedules",
			mockScheduleRepo: func(scheduleRepo *entity.MockScheduleRepository) {
				scheduleRepo.EXPECT().FindAll().Return([]entity.Schedule{}, errors.New("Error"))
			},
			wantErr: true,
		},
		{
			name: "success",
			mockScheduleRepo: func(scheduleRepo *entity.MockScheduleRepository) {
				scheduleRepo.EXPECT().FindAll().Return([]entity.Schedule{schedule1, schedule2}, nil)
			},
			want: []usecase.ScheduleQuery{
				{
					Studio:    schedule1.Studio,
					StartAt:   schedule1.StartAt,
					Performer: schedule1.Performer.Name,
					Vol:       schedule1.Vol,
				},
				{
					Studio:    schedule2.Studio,
					StartAt:   schedule2.StartAt,
					Performer: schedule2.Performer.Name,
					Vol:       schedule2.Vol,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scheduleRepo := entity.NewMockScheduleRepository(ctrl)
			tt.mockScheduleRepo(scheduleRepo)
			uc := usecase.NewScheduleQueryUsecase(scheduleRepo)
			gots, err := uc.FetchAll()

			if tt.wantErr && err == nil {
				t.Error("want error")

			}
			for i, got := range gots {
				if got != tt.want[i] {
					t.Error(got, tt.want[i])
				}
			}
		})
	}
}

func TestScheduleCommandUsecase_Save(t *testing.T) {
	ctrl := gomock.NewController(t)

	tests := []struct {
		name             string
		mockScheduleRepo func(scheduleRepo *entity.MockScheduleRepository)
		params           []usecase.ScheduleCommand
		want             []usecase.ScheduleCommand
		wantErr          bool
	}{
		{
			name:             "validation error",
			mockScheduleRepo: func(scheduleRepo *entity.MockScheduleRepository) {},
			params: []usecase.ScheduleCommand{
				{},
				{Performer: "aaa"},
				{Performer: "aaa", Vol: "bbb"},
			},
			want: []usecase.ScheduleCommand{
				{Err: usecase.ErrValidation},
				{Performer: "aaa", Err: usecase.ErrValidation},
				{Performer: "aaa", Vol: "bbb", Err: usecase.ErrValidation},
			},
			wantErr: true,
		},
		{
			name: "repository error",
			mockScheduleRepo: func(scheduleRepo *entity.MockScheduleRepository) {
				scheduleRepo.EXPECT().Save(gomock.Any()).Return(errors.New(""))
			},
			params: []usecase.ScheduleCommand{
				{Performer: "aaa", Vol: "bbb", Studio: "ccc"},
			},
			want: []usecase.ScheduleCommand{
				{Performer: "aaa", Vol: "bbb", Studio: "ccc", Err: usecase.ErrScheduleRepository},
			},
			wantErr: true,
		},
		{
			name: "success",
			mockScheduleRepo: func(scheduleRepo *entity.MockScheduleRepository) {
				scheduleRepo.EXPECT().Save(gomock.Any()).Return(nil)
			},
			params: []usecase.ScheduleCommand{
				{Performer: "aaa", Vol: "bbb", Studio: "ccc"},
				{Performer: "ddd", Vol: "eee", Studio: "fff"},
			},
			want: []usecase.ScheduleCommand{
				{Performer: "aaa", Vol: "bbb", Studio: "ccc"},
				{Performer: "ddd", Vol: "eee", Studio: "fff"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scheduleRepo := entity.NewMockScheduleRepository(ctrl)
			tt.mockScheduleRepo(scheduleRepo)
			uc := usecase.NewScheduleCommandUsecase(scheduleRepo)
			gots, err := uc.BulkUpsert(tt.params)

			if tt.wantErr && err == nil {
				t.Error("want error")
			}
			for i, got := range gots {
				if got != tt.want[i] {
					t.Error(got, tt.want[i])
				}
			}
		})
	}
}
