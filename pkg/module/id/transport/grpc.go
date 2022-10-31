package transport

import (
	"context"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/mingslife/bone"
	"google.golang.org/grpc"

	"github.com/mingslife/toolbox/pkg/common"
	"github.com/mingslife/toolbox/pkg/module/id/endpoint"
	"github.com/mingslife/toolbox/pkg/module/id/pb"
)

const (
	HandlerGenerateSnowflakeId = iota
)

type IdGrpc struct {
	pb.UnimplementedIdServiceServer

	hm map[int]grpctransport.Handler

	Grpc     *grpc.Server         `inject:"application.grpc"`
	Endpoint *endpoint.IdEndpoint `inject:""`
	Decoder  *IdGrpcDecoder       `inject:""`
	Encoder  *IdGrpcEncoder       `inject:""`
}

func (t *IdGrpc) Register() error {
	t.hm = map[int]grpctransport.Handler{
		HandlerGenerateSnowflakeId: grpctransport.NewServer(t.Endpoint.GenerateSnowflakeId, t.Decoder.GenerateSnowflakeId, t.Encoder.GenerateSnowflakeId),
	}
	return nil
}

func (t *IdGrpc) GenerateSnowflakeId(ctx context.Context, req *pb.GenerateSnowflakeIdReq) (*pb.GenerateSnowflakeIdRsp, error) {
	return common.ServeGRPC(t.hm[HandlerGenerateSnowflakeId], ctx, req, &pb.GenerateSnowflakeIdRsp{})
}

var _ bone.Transport = (*IdGrpc)(nil)
var _ pb.IdServiceServer = (*IdGrpc)(nil)
