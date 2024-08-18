//go:build wireinject
// +build wireinject

package handler

import (
	"github.com/decorickey/go-apps/internal/bmonster/infrastructure/sql"
	"github.com/google/wire"
)

func InitializeHandler() Handler {
	wire.Build(
		NewHandler,
		sql.NewStudioDao,
		sql.NewPerformerDao,
		sql.NewTimetableDao,
		sql.NewDB,
	)
	return handler{}
}
