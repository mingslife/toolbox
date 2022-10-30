package endpoint

import (
	"context"

	"github.com/mingslife/bone"

	"github.com/mingslife/toolbox/pkg/module/id/entity"
	"github.com/mingslife/toolbox/pkg/module/id/service"
)

type IdEndpoint struct {
	Service   *service.IdService `inject:""`
	Validator *IdValidator       `inject:""`
}

func (e *IdEndpoint) GenerateSnowflakeId(context.Context, any) (any, error) {
	var (
		rsp = new(entity.GenerateSnowflakeIdRsp)
		err error
	)
	rsp.Data = e.Service.GenerateSnowflakeId()
	return rsp, err
}

var _ bone.Endpoint = (*IdEndpoint)(nil)
