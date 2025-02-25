// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package usecase

import (
	"github.com/decorickey/go-apps/internal/bmonster/infrastructure/scraping"
	"github.com/decorickey/go-apps/internal/bmonster/infrastructure/sql"
)

// Injectors from wire.go:

func InitializeScrapingUsecase() ScrapingUsecase {
	scrapingRepository := scraping.NewScrapingRepository()
	db := sql.NewDB()
	studioRepository := sql.NewStudioRepository(db)
	performerRepository := sql.NewPerformerRepository(db)
	programRepository := sql.NewProgramRepository(db)
	scheduleRepository := sql.NewScheduleRepository(db)
	usecaseScrapingUsecase := NewScrapingUsecase(scrapingRepository, studioRepository, performerRepository, programRepository, scheduleRepository)
	return usecaseScrapingUsecase
}
