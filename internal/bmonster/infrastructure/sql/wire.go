//go:build wireinject
// +build wireinject

package sql

import "github.com/google/wire"

var RepositorySet = wire.NewSet(
	NewDB,
	NewStudioRepository,
	NewProgramRepository,
	NewPerformerRepository,
	NewScheduleRepository,
)

var DaoSet = wire.NewSet(
	NewDB,
	NewStudioDao,
	NewPerformerDao,
	NewTimetableDao,
)
