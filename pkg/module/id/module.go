package id

import (
	"github.com/mingslife/bone"

	"github.com/mingslife/toolbox/pkg/module/id/transport"
)

type Module struct {
	Http *transport.IdHttp `inject:""`
	Grpc *transport.IdGrpc `inject:""`
}

func (*Module) Name() string {
	return "module.id"
}

func (*Module) Init() error {
	return nil
}

func (m *Module) Register() error {
	if err := m.Http.Register(); err != nil {
		return err
	}
	if err := m.Grpc.Register(); err != nil {
		return err
	}
	return nil
}

func (*Module) Unregister() error {
	return nil
}

var _ bone.Module = (*Module)(nil)
