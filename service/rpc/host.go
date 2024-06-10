// Package rpc
// Date: 2022/11/9 10:18
// Author: Amu
// Description:
package rpc

import (
	"time"

	"github.com/amuluze/amvector/service/model"
	"github.com/amuluze/amvector/service/schema"
)

func (s *Service) HostInfo(args *schema.HostArgs, reply model.Host) error {
	if err := s.db.Order("timestamp desc").Take(&reply).Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) CPUInfo(args *schema.CPUArgs, reply model.CPU) error {
	if err := s.db.Order("timestamp desc").Take(&reply).Error; err != nil {
		return err
	}
	return nil
}
func (s *Service) CPUUsage(args *schema.CPUUsageArgs, reply []model.CPU) error {
	if err := s.db.Model(&model.CPU{}).Where("timestamp > ? and timestamp < ?", time.Unix(args.StartTime, 0), time.Unix(args.EndTime, 0)).Order("timestamp asc").Find(&reply).Error; err != nil {
		return nil
	}
	return nil
}
func (s *Service) MemInfo(args *schema.MemoryArgs, reply model.Memory) error {
	if err := s.db.Order("timestamp desc").Take(&reply).Error; err != nil {
		return err
	}
	return nil
}
func (s *Service) MemUsage(args *schema.MemoryUsageArgs, reply []model.Memory) error {
	if err := s.db.Model(&model.Memory{}).Where("timestamp > ?", time.Unix(args.StartTime, 0)).Order("timestamp asc").Find(&reply).Error; err != nil {
		return nil
	}
	return nil
}
func (s *Service) DiskInfo(args *schema.DiskArgs, reply []model.Disk) error {
	if err := s.db.Model(&model.Disk{}).Group("device").Order("timestamp desc").Find(&reply).Error; err != nil {
		return err
	}
	return nil
}
func (s *Service) DiskUsage(args schema.DiskUsageArgs, reply []model.Disk) error {
	if err := s.db.Model(&model.Disk{}).Where("timestamp > ?", time.Unix(args.StartTime, 0)).Order("timestamp asc").Find(&reply).Error; err != nil {
		return nil
	}
	return nil
}
func (s *Service) NetUsage(args schema.NetworkUsageArgs, reply []model.Net) error {
	if err := s.db.Model(&model.Net{}).Where("timestamp > ?", time.Unix(args.StartTime, 0)).Order("timestamp asc").Find(&reply).Error; err != nil {
		return nil
	}
	return nil
}
