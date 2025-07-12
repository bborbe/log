// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/log"
)

var _ = Describe("Log SamplerList", func() {
	var sampler log.Sampler
	Context("IsSample", func() {
		Context("empty", func() {
			BeforeEach(func() {
				sampler = log.SamplerList{}
			})
			It("returns false", func() {
				Expect(sampler.IsSample()).To(BeFalse())
			})
		})
		Context("one true", func() {
			BeforeEach(func() {
				sampler = log.SamplerList{
					log.SamplerFunc(func() bool {
						return true
					}),
				}
			})
			It("returns true", func() {
				Expect(sampler.IsSample()).To(BeTrue())
			})
		})
		Context("one false", func() {
			BeforeEach(func() {
				sampler = log.SamplerList{
					log.SamplerFunc(func() bool {
						return false
					}),
				}
			})
			It("returns false", func() {
				Expect(sampler.IsSample()).To(BeFalse())
			})
		})
		Context("true and false", func() {
			BeforeEach(func() {
				sampler = log.SamplerList{
					log.SamplerFunc(func() bool {
						return true
					}),
					log.SamplerFunc(func() bool {
						return false
					}),
				}
			})
			It("returns true", func() {
				Expect(sampler.IsSample()).To(BeTrue())
			})
		})
		Context("false and true", func() {
			BeforeEach(func() {
				sampler = log.SamplerList{
					log.SamplerFunc(func() bool {
						return false
					}),
					log.SamplerFunc(func() bool {
						return true
					}),
				}
			})
			It("returns true", func() {
				Expect(sampler.IsSample()).To(BeTrue())
			})
		})
	})
})
