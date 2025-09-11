// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log

import (
	"runtime"
	"sync"
	"time"

	"github.com/golang/glog"
)

//counterfeiter:generate -o mocks/memory-monitor.go --fake-name MemoryMonitor . MemoryMonitor

// MemoryMonitor provides memory usage monitoring functionality
type MemoryMonitor interface {
	LogMemoryUsage(name string)
	LogMemoryUsageOnStart()
	LogMemoryUsageOnEnd()
}

// NewMemoryMonitor creates a new memory monitor that logs memory usage at specified intervals
func NewMemoryMonitor(logInterval time.Duration) MemoryMonitor {
	return &memoryMonitor{
		logInterval: logInterval,
		lastLogTime: time.Time{}, // zero time initially
	}
}

type memoryMonitor struct {
	logInterval time.Duration
	lastLogTime time.Time
	mutex       sync.Mutex
}

// LogMemoryUsage logs current memory usage if enough time has passed since last log (thread-safe)
func (m *memoryMonitor) LogMemoryUsage(name string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	now := time.Now()

	// Check if enough time has passed since last log
	if m.lastLogTime.IsZero() || now.Sub(m.lastLogTime) >= m.logInterval {
		// Update the last log time first
		m.lastLogTime = now

		// Then log the memory usage
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)

		glog.Infof(
			"MEMORY USAGE - %s, Alloc: %.2f MB, Sys: %.2f MB, HeapInUse: %.2f MB, HeapObjects: %d, NumGC: %d",
			name,
			float64(memStats.Alloc)/1024/1024,
			float64(memStats.Sys)/1024/1024,
			float64(memStats.HeapInuse)/1024/1024,
			memStats.HeapObjects,
			memStats.NumGC,
		)
	}
}

// LogMemoryUsageOnStart logs memory usage at the beginning of backtesting
func (m *memoryMonitor) LogMemoryUsageOnStart() {
	glog.Infof("MEMORY MONITOR - Started")

	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	glog.Infof(
		"MEMORY USAGE - START - Alloc: %.2f MB, Sys: %.2f MB, HeapInUse: %.2f MB, HeapObjects: %d, NumGC: %d",
		float64(memStats.Alloc)/1024/1024,
		float64(memStats.Sys)/1024/1024,
		float64(memStats.HeapInuse)/1024/1024,
		memStats.HeapObjects,
		memStats.NumGC,
	)
}

// LogMemoryUsageOnEnd logs memory usage at the end of backtesting
func (m *memoryMonitor) LogMemoryUsageOnEnd() {
	glog.Infof("MEMORY MONITOR - Completed")

	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	glog.Infof(
		"MEMORY USAGE - END - Alloc: %.2f MB, Sys: %.2f MB, HeapInUse: %.2f MB, HeapObjects: %d, NumGC: %d",
		float64(memStats.Alloc)/1024/1024,
		float64(memStats.Sys)/1024/1024,
		float64(memStats.HeapInuse)/1024/1024,
		memStats.HeapObjects,
		memStats.NumGC,
	)

	// Force garbage collection and log again to see the difference
	runtime.GC()
	time.Sleep(100 * time.Millisecond) // Give GC a moment to complete

	runtime.ReadMemStats(&memStats)
	glog.Infof(
		"MEMORY USAGE (after GC) - Alloc: %.2f MB, Sys: %.2f MB, HeapInUse: %.2f MB, HeapObjects: %d, NumGC: %d",
		float64(memStats.Alloc)/1024/1024,
		float64(memStats.Sys)/1024/1024,
		float64(memStats.HeapInuse)/1024/1024,
		memStats.HeapObjects,
		memStats.NumGC,
	)
}
