package transport

import (
	"net/http"

	"github.com/mingslife/bone"

	"github.com/mingslife/toolbox/pkg/common"
	"github.com/mingslife/toolbox/pkg/module/id/endpoint"
)

type IdHttp struct {
	Router   *bone.Router         `inject:"application.router"`
	Endpoint *endpoint.IdEndpoint `inject:""`
	Decoder  *IdHttpDecoder       `inject:""`
}

func (t *IdHttp) Register() error {
	s := t.Router.PathPrefix("/api/v1/id").Subrouter()
	s.Methods(http.MethodGet, http.MethodGet).Path("/snowflake").Handler(common.NewServer(t.Endpoint.GenerateSnowflakeId, t.Decoder.GenerateSnowflakeId))
	return nil
}

var _ bone.Transport = (*IdHttp)(nil)
