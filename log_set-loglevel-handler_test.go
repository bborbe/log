// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/log"
	"github.com/bborbe/log/mocks"
)

var _ = Describe("Log SetLoglevelHandler", func() {
	var handler http.Handler
	var mockLogLevelSetter *mocks.LogLevelSetter
	var ctx context.Context

	BeforeEach(func() {
		ctx = context.Background()
		mockLogLevelSetter = &mocks.LogLevelSetter{}
		handler = log.NewSetLoglevelHandler(ctx, mockLogLevelSetter)
	})

	Context("with valid log level", func() {
		It("sets log level and returns success message", func() {
			mockLogLevelSetter.SetReturns(nil)

			req := httptest.NewRequest("POST", "/loglevel/3", nil)
			req = mux.SetURLVars(req, map[string]string{"level": "3"})
			resp := httptest.NewRecorder()

			handler.ServeHTTP(resp, req)

			Expect(resp.Code).To(Equal(http.StatusOK))
			Expect(resp.Body.String()).To(Equal("set loglevel to 3 completed\n"))
			Expect(mockLogLevelSetter.SetCallCount()).To(Equal(1))

			actualCtx, actualLevel := mockLogLevelSetter.SetArgsForCall(0)
			Expect(actualCtx).To(Equal(ctx))
			Expect(actualLevel).To(Equal(glog.Level(3)))
		})

		It("handles level 0", func() {
			mockLogLevelSetter.SetReturns(nil)

			req := httptest.NewRequest("POST", "/loglevel/0", nil)
			req = mux.SetURLVars(req, map[string]string{"level": "0"})
			resp := httptest.NewRecorder()

			handler.ServeHTTP(resp, req)

			Expect(resp.Code).To(Equal(http.StatusOK))
			Expect(resp.Body.String()).To(Equal("set loglevel to 0 completed\n"))

			_, actualLevel := mockLogLevelSetter.SetArgsForCall(0)
			Expect(actualLevel).To(Equal(glog.Level(0)))
		})

		It("handles high log levels", func() {
			mockLogLevelSetter.SetReturns(nil)

			req := httptest.NewRequest("POST", "/loglevel/10", nil)
			req = mux.SetURLVars(req, map[string]string{"level": "10"})
			resp := httptest.NewRecorder()

			handler.ServeHTTP(resp, req)

			Expect(resp.Code).To(Equal(http.StatusOK))
			Expect(resp.Body.String()).To(Equal("set loglevel to 10 completed\n"))

			_, actualLevel := mockLogLevelSetter.SetArgsForCall(0)
			Expect(actualLevel).To(Equal(glog.Level(10)))
		})
	})

	Context("with invalid log level", func() {
		It("returns error for non-numeric level", func() {
			req := httptest.NewRequest("POST", "/loglevel/invalid", nil)
			req = mux.SetURLVars(req, map[string]string{"level": "invalid"})
			resp := httptest.NewRecorder()

			handler.ServeHTTP(resp, req)

			Expect(resp.Code).To(Equal(http.StatusOK)) // Handler doesn't set error status codes
			Expect(resp.Body.String()).To(ContainSubstring("parse loglevel failed:"))
			Expect(mockLogLevelSetter.SetCallCount()).To(Equal(0))
		})

		It("returns error for empty level", func() {
			req := httptest.NewRequest("POST", "/loglevel/", nil)
			req = mux.SetURLVars(req, map[string]string{"level": ""})
			resp := httptest.NewRecorder()

			handler.ServeHTTP(resp, req)

			Expect(resp.Code).To(Equal(http.StatusOK))
			Expect(resp.Body.String()).To(ContainSubstring("parse loglevel failed:"))
			Expect(mockLogLevelSetter.SetCallCount()).To(Equal(0))
		})

		It("returns error for fractional level", func() {
			req := httptest.NewRequest("POST", "/loglevel/3.5", nil)
			req = mux.SetURLVars(req, map[string]string{"level": "3.5"})
			resp := httptest.NewRecorder()

			handler.ServeHTTP(resp, req)

			Expect(resp.Code).To(Equal(http.StatusOK))
			Expect(resp.Body.String()).To(ContainSubstring("parse loglevel failed:"))
			Expect(mockLogLevelSetter.SetCallCount()).To(Equal(0))
		})

		It("handles negative levels", func() {
			mockLogLevelSetter.SetReturns(nil)

			req := httptest.NewRequest("POST", "/loglevel/-1", nil)
			req = mux.SetURLVars(req, map[string]string{"level": "-1"})
			resp := httptest.NewRecorder()

			handler.ServeHTTP(resp, req)

			Expect(resp.Code).To(Equal(http.StatusOK))
			Expect(resp.Body.String()).To(Equal("set loglevel to -1 completed\n"))

			_, actualLevel := mockLogLevelSetter.SetArgsForCall(0)
			Expect(actualLevel).To(Equal(glog.Level(-1)))
		})
	})

	Context("when LogLevelSetter.Set fails", func() {
		It("returns error message", func() {
			expectedErr := errors.New("setter failed")
			mockLogLevelSetter.SetReturns(expectedErr)

			req := httptest.NewRequest("POST", "/loglevel/5", nil)
			req = mux.SetURLVars(req, map[string]string{"level": "5"})
			resp := httptest.NewRecorder()

			handler.ServeHTTP(resp, req)

			Expect(resp.Code).To(Equal(http.StatusOK))
			Expect(resp.Body.String()).To(Equal("set loglevel failed: setter failed\n"))
			Expect(mockLogLevelSetter.SetCallCount()).To(Equal(1))
		})
	})

	Context("without mux vars", func() {
		It("handles missing level parameter", func() {
			req := httptest.NewRequest("POST", "/loglevel", nil)
			// Don't set mux vars to simulate missing parameter
			resp := httptest.NewRecorder()

			handler.ServeHTTP(resp, req)

			Expect(resp.Code).To(Equal(http.StatusOK))
			Expect(resp.Body.String()).To(ContainSubstring("parse loglevel failed:"))
			Expect(mockLogLevelSetter.SetCallCount()).To(Equal(0))
		})
	})

	Context("different HTTP methods", func() {
		It("handles GET request", func() {
			mockLogLevelSetter.SetReturns(nil)

			req := httptest.NewRequest("GET", "/loglevel/2", nil)
			req = mux.SetURLVars(req, map[string]string{"level": "2"})
			resp := httptest.NewRecorder()

			handler.ServeHTTP(resp, req)

			Expect(resp.Code).To(Equal(http.StatusOK))
			Expect(resp.Body.String()).To(Equal("set loglevel to 2 completed\n"))
		})

		It("handles PUT request", func() {
			mockLogLevelSetter.SetReturns(nil)

			req := httptest.NewRequest("PUT", "/loglevel/4", nil)
			req = mux.SetURLVars(req, map[string]string{"level": "4"})
			resp := httptest.NewRecorder()

			handler.ServeHTTP(resp, req)

			Expect(resp.Code).To(Equal(http.StatusOK))
			Expect(resp.Body.String()).To(Equal("set loglevel to 4 completed\n"))
		})
	})

	Context("with request body", func() {
		It("ignores request body and uses URL parameter", func() {
			mockLogLevelSetter.SetReturns(nil)

			req := httptest.NewRequest("POST", "/loglevel/6", strings.NewReader("ignored body"))
			req = mux.SetURLVars(req, map[string]string{"level": "6"})
			resp := httptest.NewRecorder()

			handler.ServeHTTP(resp, req)

			Expect(resp.Code).To(Equal(http.StatusOK))
			Expect(resp.Body.String()).To(Equal("set loglevel to 6 completed\n"))

			_, actualLevel := mockLogLevelSetter.SetArgsForCall(0)
			Expect(actualLevel).To(Equal(glog.Level(6)))
		})
	})
})
