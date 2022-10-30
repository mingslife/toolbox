package id

import (
	"github.com/mingslife/bone"

	"github.com/mingslife/toolbox/pkg/module/id/transport"
)

type Module struct {
	Http *transport.IdHttp `inject:""`
}

func (*Module) Name() string {
	return "module.id"
}

func (*Module) Init() error {
	return nil
}

func (m *Module) Register() error {
	return m.Http.Register()
}

func (*Module) Unregister() error {
	return nil
}

var _ bone.Module = (*Module)(nil)
