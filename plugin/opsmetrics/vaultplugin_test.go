package opsmetrics_test

import (
	. "github.com/enaml-ops/opsmetrics-decorator-plugin/plugin/opsmetrics"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("bla", func() {
	var plgn *Plugin

	BeforeEach(func() {
		plgn = new(Plugin)
	})

	Context("bleh", func() {

		It("blu", func() {
			Ω(true).Should(BeTrue())
		})
	})
})
