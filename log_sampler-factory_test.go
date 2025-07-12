// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/log"
)

var _ = Describe("Log SamplerFactory", func() {
	var samplerFactory log.SamplerFactory
	Context("DefaultSamplerFactory", func() {
		BeforeEach(func() {
			samplerFactory = log.DefaultSamplerFactory
		})
		It("returns a sampler", func() {
			Expect(samplerFactory.Sampler()).NotTo(BeNil())
		})
	})
	Context("SamplerFactoryFunc", func() {
		var sampler log.Sampler
		BeforeEach(func() {
			sampler = log.SamplerFunc(func() bool {
				return true
			})
			samplerFactory = log.SamplerFactoryFunc(func() log.Sampler {
				return sampler
			})
		})
		It("returns a sampler", func() {
			Expect(samplerFactory.Sampler().IsSample()).To(Equal(sampler.IsSample()))
		})
	})
})
