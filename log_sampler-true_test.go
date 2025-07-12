// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/log"
)

var _ = Describe("Log SamplerTrue", func() {
	var sampler log.Sampler
	Context("IsSample", func() {
		BeforeEach(func() {
			sampler = log.NewSamplerTrue()
		})
		It("returns true", func() {
			Expect(sampler.IsSample()).To(BeTrue())
		})
	})
})
