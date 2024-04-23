// Package main
// Date: 2024/4/23 20:12
// Author: Amu
// Description:
package main

import (
	"github.com/takama/daemon"
	"os"
)

type Service struct {
	daemon daemon.Daemon
}

func (s *Service) Start() {}

func (s *Service) Run() {}

func (s *Service) Stop() {}

func (s *Service) manager() (string, error) {
	usage := "Usage: amvector install | remove | start | stop | status"

	if len(os.Args) > 0 {
		cmd := os.Args[0]
		switch cmd {
		case "install":
			return s.daemon.Install()
		case "remove":
			return s.daemon.Remove()
		case "start":
			return s.daemon.Start()
		case "stop":
			return s.daemon.Stop()
		case "status":
			return s.daemon.Status()
		default:
			return usage, nil
		}
	}

	return s.daemon.Run(s)
}
