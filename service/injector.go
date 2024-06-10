// Package service
// Date: 2024/06/10 18:20:47
// Author: Amu
// Description:
package service

import (
	"github.com/google/wire"
)

var InjectorSet = wire.NewSet(NewInjector)

type Injector struct {
	Config *Config
}

func NewInjector(config *Config) (*Injector, error) {
	return &Injector{
		Config: config,
	}, nil
}
