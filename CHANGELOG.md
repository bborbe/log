# Changelog

All notable changes to this project will be documented in this file.

Please choose versions by [Semantic Versioning](http://semver.org/).

* MAJOR version when you make incompatible API changes,
* MINOR version when you add functionality in a backwards-compatible manner, and
* PATCH version when you make backwards-compatible bug fixes.

## v1.4.1

- Add LogMemoryUsagef method for formatted memory logging
- Refactor memory stats logging to reduce code duplication
- Extract MemoryStats utility function to separate file
- Remove trading-specific comments from memory monitor

## v1.4.0

- Add MemoryMonitor interface for runtime memory usage monitoring
- Implement memory monitoring with configurable logging intervals
- Add memory usage logging at start/end with garbage collection metrics
- Generate MemoryMonitor mock for testing

## v1.3.0

- Add comprehensive GoDoc documentation to all exported functions and types
- Enhance README.md with detailed usage examples, API documentation, and status badges
- Add GitHub Actions CI/CD workflow with automated testing and coverage reporting
- Improve .gitignore with comprehensive Go development patterns
- Add status badges for CI, Go Report Card, pkg.go.dev reference, and code coverage
- Enhance project structure following Go library best practices

## v1.2.1

- go mod update
- add tests

## v1.2.0

- generate mocks
- go mod update
- add tests

## v1.1.0

- remove vendor
- go mod update

## v1.0.1

- add license
- go mod update

## v1.0.0

- Initial Version
