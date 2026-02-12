// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log_test

import (
	"context"
	"errors"
	"time"

	"github.com/golang/glog"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/log"
)

var _ = Describe("Log LogLevelSetter", Serial, func() {
	var logLevelSetter log.LogLevelSetter
	var ctx context.Context

	BeforeEach(func() {
		ctx = context.Background()
	})

	Context("NewLogLevelSetter", func() {
		It("is safe for concurrent use from multiple goroutines", func() {
			// This test verifies there are no data races when multiple goroutines
			// call Set concurrently. Run with: go test -race ./...
			logLevelSetter = log.NewLogLevelSetter(glog.Level(1), 50*time.Millisecond)

			const numGoroutines = 20
			done := make(chan bool, numGoroutines)

			for i := 0; i < numGoroutines; i++ {
				go func(level int) {
					defer func() { done <- true }()
					_ = logLevelSetter.Set(ctx, glog.Level(level%5))
				}(i)
			}

			// Wait for all goroutines to complete
			for i := 0; i < numGoroutines; i++ {
				<-done
			}

			// Wait for resets to trigger
			time.Sleep(100 * time.Millisecond)
		})
	})

	Context("LogLevelSetterFunc", func() {
		It("calls the wrapped function", func() {
			called := false
			var receivedLevel glog.Level
			logLevelSetter = log.LogLevelSetterFunc(
				func(ctx context.Context, logLevel glog.Level) error {
					called = true
					receivedLevel = logLevel
					return nil
				},
			)

			err := logLevelSetter.Set(ctx, glog.Level(3))
			Expect(err).ToNot(HaveOccurred())
			Expect(called).To(BeTrue())
			Expect(receivedLevel).To(Equal(glog.Level(3)))
		})

		It("returns error from wrapped function", func() {
			expectedErr := errors.New("test error")
			logLevelSetter = log.LogLevelSetterFunc(
				func(ctx context.Context, logLevel glog.Level) error {
					return expectedErr
				},
			)

			err := logLevelSetter.Set(ctx, glog.Level(1))
			Expect(err).To(Equal(expectedErr))
		})
	})

})
