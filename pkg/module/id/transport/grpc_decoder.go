package transport

import "context"

type IdGrpcDecoder struct{}

func (*IdGrpcDecoder) GenerateSnowflakeId(_ context.Context, r any) (any, error) {
	return nil, nil
}
