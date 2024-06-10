// Package service
// Date: 2024/06/10 18:26:22
// Author: Amu
// Description:
package service

import (
	"github.com/spf13/viper"
)

type Config struct {
	Gorm     Gorm
	DB       DB
	Disk     Disk
	Task     Task
	Ethernet Ethernet
}

func NewConfig(configFile string) (*Config, error) {
	config := &Config{}

	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}

	return config, nil
}
