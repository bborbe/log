// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log_test

import (
	"flag"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/log"
)

var _ = Describe("Log SamplerGlogLevel", func() {
	var sampler log.Sampler
	Context("IsSample", func() {
		Context("V(0)", func() {
			BeforeEach(func() {
				_ = flag.Set("v", "0")
				sampler = log.NewSamplerGlogLevel(0)
			})
			It("returns true", func() {
				Expect(sampler.IsSample()).To(BeTrue())
			})
		})
		Context("V(1)", func() {
			BeforeEach(func() {
				_ = flag.Set("v", "0")
				sampler = log.NewSamplerGlogLevel(1)
			})
			It("returns false", func() {
				Expect(sampler.IsSample()).To(BeFalse())
			})
		})
	})
})
