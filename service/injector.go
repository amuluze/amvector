// Package service
// Date: 2024/06/10 18:20:47
// Author: Amu
// Description:
package service

import (
	"github.com/amuluze/amutool/logger"
	"github.com/google/wire"
)

var InjectorSet = wire.NewSet(NewInjector)

type Injector struct {
	Config *Config
	Logger *logger.Logger
	Task   *TimedTask
}

func NewInjector(config *Config, task *TimedTask, logx *logger.Logger) (*Injector, error) {
	return &Injector{
		Config: config,
		Task:   task,
		Logger: logx,
	}, nil
}
