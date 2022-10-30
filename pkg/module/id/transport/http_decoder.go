package transport

import (
	"context"
	"net/http"
)

type IdHttpDecoder struct{}

func (*IdHttpDecoder) GenerateSnowflakeId(_ context.Context, r *http.Request) (any, error) {
	return nil, nil
}
