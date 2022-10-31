package entity

import (
	"github.com/mingslife/toolbox/pkg/common"
)

type GenerateSnowflakeIdRsp struct {
	common.BaseDTO
	Data int64 `json:"data"`
}
