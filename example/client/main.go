// Package main
// Date: 2024/06/11 16:29:19
// Author: Amu
// Description:
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/smallnest/rpcx/client"
	"gorm.io/gorm"
)

type ContainerQueryArgs struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

type Container struct {
	gorm.Model
	Timestamp   time.Time
	ContainerID string
	Name        string
	Image       string
	IP          string
	State       string
	Uptime      string
	CPUPercent  float64
	MemPercent  float64
	MemUsage    float64
	MemLimit    float64
}

func main() {
	addr := "/data/amvector/amvector.socket"
	d, _ := client.NewPeer2PeerDiscovery("unix@"+addr, "")
	xclient := client.NewXClient("Service", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	args := &ContainerQueryArgs{Page: 1, Size: 20}
	var reply []Container
	err := xclient.Call(context.Background(), "ContainerList", args, &reply)
	if err != nil {
		fmt.Println("failed to call:", err)
	}
	fmt.Println("reply:", reply)
}
