// Package rpc
// Date: 2022/11/9 10:18
// Author: Amu
// Description:
package rpc

import (
	"context"

	"github.com/amuluze/amvector/service/model"

	"github.com/amuluze/amvector/service/schema"
)

var ctx = context.TODO()

func (s *Service) DockerVersion(args *schema.VersionArgs, reply model.Docker) error {
	if err := s.db.Model(&model.Docker{}).First(&reply).Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) ContainerList(args *schema.ContainerQueryArgs, reply model.Containers) error {
	if err := s.db.Model(&model.Container{}).Order("created_at desc").Offset((args.Page - 1) * args.Size).Limit(args.Size).Find(&reply).Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) ContainerCount(args *schema.QueryCountArgs, reply *schema.QueryCountReply) error {
	var count int64
	if err := s.db.Model(&model.Container{}).Count(&count).Error; err != nil {
		return err
	}
	reply.Count = int(count)
	return nil
}

func (s *Service) ContainerCreate(args *schema.ContainerCreateArgs, reply *schema.ContainerCreateReply) error {
	return nil
}

func (s *Service) ContainerDelete(args *schema.ContainerDeleteArgs, reply *schema.ContainerDeleteReply) error {
	if err := s.manager.DeleteContainer(ctx, args.ContainerID); err != nil {
		return err
	}
	if err := s.db.Model(&model.Container{}).Delete(&model.Container{ContainerID: args.ContainerID}).Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) ContainerStart(args *schema.ContainerStartArgs, reply *schema.ContainerStartReply) error {
	if err := s.manager.StartContainer(ctx, args.ContainerID); err != nil {
		return err
	}
	if err := s.db.Model(&model.Container{}).Where("container_id = ?", args.ContainerID).Update("state", "running").Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) ContainerStop(args *schema.ContainerStopArgs, reply *schema.ContainerStopReply) error {
	if err := s.manager.StopContainer(ctx, args.ContainerID); err != nil {
		return err
	}
	if err := s.db.Model(&model.Container{}).Where("container_id = ?", args.ContainerID).Update("state", "stopped").Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) ContainerRestart(args *schema.ContainerRestartArgs, reply *schema.ContainerRestartReply) error {
	if err := s.manager.RestartContainer(ctx, args.ContainerID); err != nil {
		return err
	}
	if err := s.db.Model(&model.Container{}).Where("container_id = ?", args.ContainerID).Update("state", "running").Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) ImageList(args *schema.ImageQueryArgs, reply model.Images) error {
	if err := s.db.Model(&model.Image{}).Order("created_at desc").Offset((args.Page - 1) * args.Size).Limit(args.Size).Find(&reply).Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) ImagePull(args *schema.ImagePullArgs, reply *schema.ImagePullReply) error {
	if err := s.manager.PullImage(ctx, args.ImageName); err != nil {
		return err
	}
	return nil
}

func (s *Service) ImageTag(args *schema.ImageTagArgs, reply *schema.ImageTagReply) error {
	if err := s.manager.TagImage(ctx, args.OldTag, args.NewTag); err != nil {
		return err
	}
	return nil
}

func (s *Service) ImageCount(args *schema.ImageCountArgs, reply *schema.ImageCountReply) error {
	var total int64
	if err := s.db.Model(&model.Images{}).Order("created_at desc").Count(&total).Error; err != nil {
		return err
	}
	reply.Count = int(total)
	return nil
}

func (s *Service) ImageDelete(args *schema.ImageDeleteArgs, reply *schema.ImageDeleteReply) error {
	if err := s.manager.RemoveImage(ctx, args.ImageID); err != nil {
		return err
	}
	if err := s.db.Where("image_id = ?", args.ImageID).Delete(&model.Image{}).Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) ImagesPrune() error {
	return s.manager.PruneImages(ctx)
}
