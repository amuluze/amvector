// Package rpc
// Date: 2022/11/9 10:18
// Author: Amu
// Description:
package rpc

import (
	"github.com/amuluze/amutool/database"
	"github.com/amuluze/amutool/docker"
	"log/slog"
)

type Service struct {
	db      *database.DB
	manager *docker.Manager
}

func NewService(db *database.DB, manager *docker.Manager) *Service {
	return &Service{db: db, manager: manager}
}

func (s *Service) Ping() error {
	slog.Info("rpc service ping")
	return nil
}
