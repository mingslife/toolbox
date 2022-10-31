package transport

import (
	"context"

	"github.com/mingslife/toolbox/pkg/module/id/entity"
	"github.com/mingslife/toolbox/pkg/module/id/pb"
)

type IdGrpcEncoder struct{}

func (*IdGrpcEncoder) GenerateSnowflakeId(_ context.Context, r any) (any, error) {
	rsp, ps := r.(*entity.GenerateSnowflakeIdRsp), &pb.GenerateSnowflakeIdRsp{}
	ps.Code = rsp.Code
	ps.Message = rsp.Message
	ps.Data = rsp.Data
	return ps, nil
}
