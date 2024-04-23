// Package main
// Date: 2024/4/23 19:33
// Author: Amu
// Description:
package main

import (
	"fmt"
	"github.com/takama/daemon"
	"os"
	"runtime"
)

const (
	name        = "amvector"
	description = "data collection and synchronization services"
)

var dependencies = []string{"dummy.service"}

func main() {
	var kind daemon.Kind
	if runtime.GOOS == "darwin" {
		kind = daemon.GlobalDaemon
	} else if runtime.GOOS == "linux" {
		kind = daemon.SystemDaemon
	}

	src, err := daemon.New(name, description, kind, dependencies...)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	service := &Service{daemon: src}
	status, err := service.manager()
	if err != nil {
		fmt.Println(status, "\nError: ", err)
		os.Exit(1)
	}
	fmt.Println(status)
}
