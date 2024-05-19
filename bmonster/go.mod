module github.com/decorickey/go-apps/bmonster

go 1.22.2

require (
	github.com/decorickey/go-apps/lib/timeutil v0.0.0
	go.uber.org/mock v0.4.0
)

replace github.com/decorickey/go-apps/lib/timeutil => ../lib/timeutil
