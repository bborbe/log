// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/log"
)

var _ = Describe("Log SamplerMod", func() {
	var sampler log.Sampler
	Context("IsSample", func() {
		Context("mod 1", func() {
			BeforeEach(func() {
				sampler = log.NewSampleMod(1)
			})
			It("returns true", func() {
				Expect(sampler.IsSample()).To(BeTrue())
				Expect(sampler.IsSample()).To(BeTrue())
				Expect(sampler.IsSample()).To(BeTrue())
			})
		})
		Context("mod 2", func() {
			BeforeEach(func() {
				sampler = log.NewSampleMod(2)
			})
			It("returns false,true,false", func() {
				Expect(sampler.IsSample()).To(BeFalse())
				Expect(sampler.IsSample()).To(BeTrue())
				Expect(sampler.IsSample()).To(BeFalse())
			})
		})
	})
})
