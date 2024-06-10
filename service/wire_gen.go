// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package service

import (
	"github.com/amuluze/amvector/service/model"
)

// Injectors from wire.go:

func BuildInjector(configFile string) (*Injector, func(), error) {
	config, err := NewConfig(configFile)
	if err != nil {
		return nil, nil, err
	}
	models := model.NewModels()
	db, err := NewDB(config, models)
	if err != nil {
		return nil, nil, err
	}
	timedTask := NewTimedTask(config, db)
	logger := NewLogger(config)
	injector, err := NewInjector(config, timedTask, logger)
	if err != nil {
		return nil, nil, err
	}
	return injector, func() {
	}, nil
}
