package component

import (
	"github.com/bwmarrin/snowflake"
	"github.com/mingslife/bone"

	"github.com/mingslife/toolbox/pkg/conf"
)

type Snowflake struct {
	node *snowflake.Node
}

func (*Snowflake) Name() string {
	return "component.snowflake"
}

func (c *Snowflake) Init() error {
	cfg := conf.GetConfig()
	var err error
	c.node, err = snowflake.NewNode(cfg.Node)
	return err
}

func (*Snowflake) Register() error {
	return nil
}

func (*Snowflake) Unregister() error {
	return nil
}

func (c *Snowflake) GenerateID() int64 {
	return c.node.Generate().Int64()
}

var _ bone.Component = (*Snowflake)(nil)
