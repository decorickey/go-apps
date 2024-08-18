//go:build wireinject
// +build wireinject

package usecase

import (
	"github.com/decorickey/go-apps/internal/bmonster/infrastructure/sql"
	"github.com/google/wire"
)

func InitializeScrapingUsecase() ScrapingUsecase {
	wire.Build(
		NewScrapingUsecase,
		sql.RepositorySet,
	)
	return scrapingUsecase{}
}
