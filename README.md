# Go Log Utilities

[![CI](https://github.com/bborbe/log/workflows/CI/badge.svg)](https://github.com/bborbe/log/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/bborbe/log)](https://goreportcard.com/report/github.com/bborbe/log)
[![Go Reference](https://pkg.go.dev/badge/github.com/bborbe/log.svg)](https://pkg.go.dev/github.com/bborbe/log)
[![License](https://img.shields.io/badge/License-BSD%203--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)
[![codecov](https://codecov.io/gh/bborbe/log/branch/master/graph/badge.svg)](https://codecov.io/gh/bborbe/log)

A Go library providing advanced logging utilities focused on log sampling and dynamic log level management, designed to integrate seamlessly with Google's `glog` library.

## Features

- **Log Sampling**: Reduce log volume with intelligent sampling mechanisms
- **Dynamic Log Level Management**: Change log levels at runtime via HTTP endpoints
- **Multiple Sampler Types**: Counter-based, time-based, glog-level-based, and custom samplers
- **Thread-Safe**: All components are designed for concurrent use
- **Extensible**: Factory pattern and interface-based design for easy customization

## Installation

```bash
go get github.com/bborbe/log
```

## Quick Start

### Basic Log Sampling

```go
package main

import (
    "time"
    "github.com/bborbe/log"
    "github.com/golang/glog"
)

func main() {
    // Sample every 10th log entry
    modSampler := log.NewSampleMod(10)
    
    // Sample once every 10 seconds
    timeSampler := log.NewSampleTime(10 * time.Second)
    
    // Use in your logging code
    if modSampler.IsSample() {
        glog.V(2).Infof("This will be logged every 10th time")
    }
    
    if timeSampler.IsSample() {
        glog.V(2).Infof("This will be logged at most once every 10 seconds")
    }
}
```

### Dynamic Log Level Management

```go
package main

import (
    "context"
    "net/http"
    "time"
    
    "github.com/bborbe/log"
    "github.com/golang/glog"
    "github.com/gorilla/mux"
)

func main() {
    ctx := context.Background()
    
    // Create log level setter that auto-resets after 5 minutes
    logLevelSetter := log.NewLogLevelSetter(
        glog.Level(1), // default level
        5*time.Minute, // auto-reset duration
    )
    
    // Set up HTTP handler for dynamic log level changes
    router := mux.NewRouter()
    router.Handle("/debug/loglevel/{level}", 
        log.NewSetLoglevelHandler(ctx, logLevelSetter))
    
    http.ListenAndServe(":8080", router)
}
```

Now you can change log levels at runtime:
```bash
curl http://localhost:8080/debug/loglevel/4
```

## Sampler Types

### ModSampler
Samples every Nth log entry based on a counter:
```go
sampler := log.NewSampleMod(100) // Sample every 100th log
```

### TimeSampler
Samples based on time intervals:
```go
sampler := log.NewSampleTime(30 * time.Second) // Sample at most once per 30 seconds
```

### GlogLevelSampler
Samples based on glog verbosity levels:
```go
sampler := log.NewSamplerGlogLevel(3) // Sample when glog level >= 3
```

### ListSampler
Combines multiple samplers with OR logic:
```go
sampler := log.SamplerList{
    log.NewSampleTime(10 * time.Second),
    log.NewSamplerGlogLevel(4),
}
```

### FuncSampler
Create custom sampling logic:
```go
sampler := log.SamplerFunc(func() bool {
    // Your custom sampling logic
    return shouldSample()
})
```

### TrueSampler
Always samples (useful for testing or special cases):
```go
sampler := log.SamplerTrue{}
```

## Factory Pattern

Use the factory pattern for dependency injection:

```go
// Use the default factory
factory := log.DefaultSamplerFactory
sampler := factory.Sampler()

// Or create a custom factory
customFactory := log.SamplerFactoryFunc(func() log.Sampler {
    return log.NewSampleMod(50)
})
```

## HTTP Log Level Management

The library provides built-in HTTP handlers for runtime log level changes:

- **Endpoint**: `GET/POST /debug/loglevel/{level}`
- **Auto-reset**: Automatically reverts to default level after specified duration
- **Thread-safe**: Safe for concurrent access

Example integration with gorilla/mux:
```go
router := mux.NewRouter()
logLevelSetter := log.NewLogLevelSetter(glog.Level(1), 5*time.Minute)
router.Handle("/debug/loglevel/{level}", 
    log.NewSetLoglevelHandler(context.Background(), logLevelSetter))
```

## Development

### Running Tests
```bash
make test
```

### Code Generation (Mocks)
```bash
make generate
```

### Full Development Workflow
```bash
make precommit  # Format, test, lint, and check
```

## Testing Framework

The library uses:
- **Ginkgo v2** for BDD-style testing
- **Gomega** for assertions
- **Counterfeiter** for mock generation

## License

This project is licensed under the BSD 3-Clause License - see the [LICENSE](LICENSE) file for details.

## Contributing

1. Fork the repository
2. Create your feature branch
3. Add tests for your changes
4. Run `make precommit` to ensure code quality
5. Submit a pull request

## Dependencies

- [glog](https://github.com/golang/glog) - Core logging functionality
- [gorilla/mux](https://github.com/gorilla/mux) - HTTP routing for log level endpoints
- [github.com/bborbe/time](https://github.com/bborbe/time) - Time utilities
