// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log_test

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/log"
)

var _ = Describe("Log SamplerTime", func() {
	var sampler log.Sampler
	Context("IsSample", func() {
		Context("with 100ms duration", func() {
			BeforeEach(func() {
				sampler = log.NewSampleTime(100 * time.Millisecond)
			})
			It("returns true on first call", func() {
				Expect(sampler.IsSample()).To(BeTrue())
			})
			It("returns false on immediate second call", func() {
				Expect(sampler.IsSample()).To(BeTrue())
				Expect(sampler.IsSample()).To(BeFalse())
			})
			It("returns true after duration has passed", func() {
				Expect(sampler.IsSample()).To(BeTrue())
				time.Sleep(150 * time.Millisecond)
				Expect(sampler.IsSample()).To(BeTrue())
			})
			It("returns false multiple times within duration", func() {
				Expect(sampler.IsSample()).To(BeTrue())
				Expect(sampler.IsSample()).To(BeFalse())
				Expect(sampler.IsSample()).To(BeFalse())
				Expect(sampler.IsSample()).To(BeFalse())
			})
		})
		Context("with zero duration", func() {
			BeforeEach(func() {
				sampler = log.NewSampleTime(0)
			})
			It("returns true on every call", func() {
				Expect(sampler.IsSample()).To(BeTrue())
				Expect(sampler.IsSample()).To(BeTrue())
				Expect(sampler.IsSample()).To(BeTrue())
			})
		})
		Context("with negative duration", func() {
			BeforeEach(func() {
				sampler = log.NewSampleTime(-1 * time.Millisecond)
			})
			It("returns true on every call", func() {
				Expect(sampler.IsSample()).To(BeTrue())
				Expect(sampler.IsSample()).To(BeTrue())
				Expect(sampler.IsSample()).To(BeTrue())
			})
		})
		Context("concurrent access", func() {
			BeforeEach(func() {
				sampler = log.NewSampleTime(50 * time.Millisecond)
			})
			It("handles concurrent calls safely", func() {
				done := make(chan bool, 10)
				results := make(chan bool, 10)

				for i := 0; i < 10; i++ {
					go func() {
						defer GinkgoRecover()
						result := sampler.IsSample()
						results <- result
						done <- true
					}()
				}

				// Wait for all goroutines to complete
				for i := 0; i < 10; i++ {
					<-done
				}
				close(results)

				// At least one should be true (the first one)
				trueCount := 0
				for result := range results {
					if result {
						trueCount++
					}
				}
				Expect(trueCount).To(BeNumerically(">=", 1))
			})
		})
	})
})
