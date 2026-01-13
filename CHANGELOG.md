# Changelog

All notable changes to this project will be documented in this file.

Please choose versions by [Semantic Versioning](http://semver.org/).

* MAJOR version when you make incompatible API changes,
* MINOR version when you add functionality in a backwards-compatible manner, and
* PATCH version when you make backwards-compatible bug fixes.

## v1.6.1

- Update Go to 1.25.5
- Update golang.org/x/crypto to v0.47.0
- Update dependencies

## v1.6.0

- update go and deps

## v1.5.0
- Add comprehensive package-level documentation (doc.go) with examples and usage guidance
- Add Ginkgo v2 CLI to development tools for better test execution
- Enhance README with comprehensive Testing section showing multiple testing approaches
- Add Full Example section to README demonstrating production-like usage
- Update CI badge to new GitHub Actions syntax
- Remove deprecated golint tool (replaced by golangci-lint)
- Add horizontal rules to README for better visual section separation

## v1.4.4
- Update dependencies (github.com/bborbe/time v1.20.0, github.com/securego/gosec/v2 v2.22.10, and 19 indirect dependencies)
- Add exclusion for golang.org/x/tools v0.38.0 due to counterfeiter compatibility

## v1.4.3
- Update Go version from 1.25.2 to 1.25.3

## v1.4.2
- Fix integer overflow vulnerability in log level handler (G115)
- Add golangci-lint configuration and security scanning tools
- Update Go version from 1.24.5 to 1.25.2
- Add osv-scanner, gosec, and trivy security checks to Makefile
- Update CI workflow with Trivy installation
- Update dependencies and tooling

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
