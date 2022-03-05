[![Go Report Card](https://goreportcard.com/badge/github.com/tantalor93/go-http-mock)](https://goreportcard.com/report/github.com/tantalor93/go-http-mock)
[![Tantalor93](https://circleci.com/gh/Tantalor93/go-http-mock/tree/master.svg?style=svg)](https://circleci.com/gh/Tantalor93/go-http-mock?branch=master)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/tantalor93/go-http-mock/blob/master/LICENSE)

## GO HTTP MOCK
simple standalone high-performance mock HTTP server implemented in GO, which is able to provide mocked responses defined by 
mock specifications similar to [Wiremock JSON API](https://wiremock.org/docs/stubbing/).

### Installation
`go-http-mock` can be installed using GO tooling

```
go install github.com/tantalor93/go-http-mock@latest
```

### Usage
```
Usage:
  go-http-mock [flags]

Flags:
  -d, --dir string   directory with mock specifications (default ".")
  -h, --help         help for go-http-mock
  -p, --port int     port to run mock server on (default 8081)
```

### Examples
See [specification](docs/specification.md) and [example](docs/example.md) for examples and mock specification description.
