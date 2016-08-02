package opsmetrics

import (
	"github.com/enaml-ops/enaml"
	"github.com/enaml-ops/pluginlib/pcli"
	"github.com/enaml-ops/pluginlib/product"
)

type Plugin struct {
}

func (s *Plugin) GetFlags() (flags []pcli.Flag) {
	return []pcli.Flag{}
}

func (s *Plugin) GetMeta() product.Meta {
	return product.Meta{
		Name: "opsmetrics",
	}
}

func (s *Plugin) GetProduct(args []string, cloudConfig []byte) (b []byte) {
	var dm = new(enaml.DeploymentManifest)
	return dm.Bytes()
}
