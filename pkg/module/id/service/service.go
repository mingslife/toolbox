package service

import (
	"github.com/mingslife/bone"

	"github.com/mingslife/toolbox/pkg/module/id/repo"
)

type IdService struct {
	Repo *repo.IdRepo `inject:""`
}

func (s *IdService) GenerateSnowflakeId() int64 {
	return s.Repo.GenerateSnowflakeId()
}

var _ bone.Service = (*IdService)(nil)
