package main

import (
	"github.com/enaml-ops/opsmetrics-decorator-plugin/plugin/opsmetrics"
	"github.com/enaml-ops/pluginlib/product"
)

func main() {
	product.Run(new(opsmetrics.Plugin))
}
