package common

import (
	"context"

	grpctransport "github.com/go-kit/kit/transport/grpc"
)

func ServeGRPC[T any](s grpctransport.Handler, ctx context.Context, req any, emptyRsp T) (r T, e error) {
	var rsp any
	_, rsp, e = s.ServeGRPC(ctx, req)
	if temp, ok := rsp.(T); ok {
		r = temp
	} else {
		r = emptyRsp
	}
	return
}
