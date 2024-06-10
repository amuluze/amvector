// Package main
// Date: 2024/4/23 20:12
// Author: Amu
// Description:
package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/takama/daemon"
)

type Service struct {
	daemon daemon.Daemon
}

// Start start amvector bootstrap service non-blocking
func (s *Service) Start() {
	fmt.Printf("Starting amvector bootstrap service...\n")
}

// Run start amvector bootstrap service blocking and wait for exit signal
func (s *Service) Run() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

	go func() {
		select {
		case sig := <-interrupt:
			log.Println("Got signal:", sig)
			bootstrap.Stop()
			os.Exit(0)
		}
	}()

	for {
		bootstrap.Run()
		time.Sleep(10 * time.Second)
	}
}

// Stop stop amvector bootstrap service
func (s *Service) Stop() {
	fmt.Println("Stopping amvector bootstrap service...")
}

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
