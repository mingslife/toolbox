package main

import (
	"github.com/mingslife/bone"

	"github.com/mingslife/toolbox/pkg/component"
	"github.com/mingslife/toolbox/pkg/conf"
	"github.com/mingslife/toolbox/pkg/module/id"
)

var cfg *conf.Config

func main() {
	app := bone.NewApplication(&bone.ApplicationOptions{
		Debug:      cfg.Debug,
		HttpEnable: true,
		HttpHost:   cfg.Http.Host,
		HttpPort:   cfg.Http.Port,
		GrpcEnable: true,
		GrpcHost:   cfg.Grpc.Host,
		GrpcPort:   cfg.Grpc.Port,
	})

	app.Use(
		// component
		new(component.Snowflake),

		// module
		new(id.Module),
	)

	app.Run()
}

func init() {
	cfg = conf.ParseConfig()
}
