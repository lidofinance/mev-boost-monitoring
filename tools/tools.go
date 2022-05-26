//go:build tools
// +build tools

//go:generate bash -c "go build -ldflags \"-X 'main.version=$(go list -m -f '{{.Version}}' github.com/golangci/golangci-lint)' -X 'main.commit=test' -X 'main.date=test'\" -o ../bin/golangci-lint github.com/golangci/golangci-lint/cmd/golangci-lint"
//go:generate go build -o ../bin/mockery github.com/vektra/mockery/v2
//go:generate go build -o ../bin/goimports golang.org/x/tools/cmd/goimports
//go:generate go build -tags 'postgres' -o ../bin/migrate github.com/golang-migrate/migrate/v4/cmd/migrate

//go:generate go build -o ../bin/ifacemaker github.com/vburenin/ifacemaker

// Package tools contains go:generate commands for all project tools with versions stored in local go.mod file
// See https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module
package tools

import (
	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/cmd/migrate"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/vburenin/ifacemaker"
	_ "github.com/vektra/mockery/v2"

	_ "golang.org/x/tools/cmd/goimports"
)