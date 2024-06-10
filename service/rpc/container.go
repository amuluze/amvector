// Package rpc
// Date: 2022/11/9 10:18
// Author: Amu
// Description:
package rpc

import "github.com/amuluze/amvector/service/model"

type QueryArgs struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

func (s *Service) ContainerList(args *QueryArgs, reply model.Containers) error {
	return nil
}

type QueryCountArgs struct{}
type QueryCountReply struct {
	Count int
}

func (s *Service) ContainerCount(args *QueryCountArgs, reply *QueryCountReply) error {
	return nil
}
