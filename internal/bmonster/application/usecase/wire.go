//go:build wireinject
// +build wireinject

package usecase

import (
	"github.com/decorickey/go-apps/internal/bmonster/infrastructure/scraping"
	"github.com/decorickey/go-apps/internal/bmonster/infrastructure/sql"
	"github.com/google/wire"
)

func InitializeScrapingUsecase() ScrapingUsecase {
	wire.Build(
		NewScrapingUsecase,
		scraping.NewScrapingRepository,
		sql.RepositorySet,
	)
	return scrapingUsecase{}
}
