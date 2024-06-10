//go:build wireinject
// +build wireinject

package service

import (
	"github.com/google/wire"
)

func BuildInjector(configFile string) (*Injector, func(), error) {
	wire.Build(
		NewConfig,
		NewLogger,
		NewDB,
	)
	return new(Injector), nil, nil
}
