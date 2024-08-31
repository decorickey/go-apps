//go:build tools
// +build tools

package tools

import (
	_ "github.com/google/wire/cmd/wire"
	_ "github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen"
	_ "go.uber.org/mock/mockgen"
)
