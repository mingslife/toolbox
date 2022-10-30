package repo

import (
	"github.com/mingslife/bone"

	"github.com/mingslife/toolbox/pkg/component"
)

type IdRepo struct {
	Snowflake *component.Snowflake `inject:"component.snowflake"`
}

func (r *IdRepo) GenerateSnowflakeId() int64 {
	return r.Snowflake.GenerateID()
}

var _ bone.Repo = (*IdRepo)(nil)
