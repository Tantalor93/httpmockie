[![Go Report Card](https://goreportcard.com/badge/github.com/tantalor93/httpmockie)](https://goreportcard.com/report/github.com/tantalor93/httpmockie)
[![Tantalor93](https://circleci.com/gh/Tantalor93/httpmockie/tree/master.svg?style=svg)](https://circleci.com/gh/Tantalor93/httpmockie?branch=master)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/tantalor93/httpmockie/blob/master/LICENSE)

## httpmockie
simple standalone high-performance mock HTTP server implemented in GO, which is able to provide mocked responses defined by 
mock specifications similar to [Wiremock JSON API](https://wiremock.org/docs/stubbing/).

### Installation
`httpmockie` can be installed using GO tooling

```
go install github.com/tantalor93/httpmockie@latest
```

### Usage
```
Usage:
  httpmockie [flags]

Flags:
  -d, --dir string   directory with mock specifications (default ".")
  -h, --help         help for go-http-mock
  -p, --port int     port to run mock server on (default 8081)
```

### Examples
See [specification](docs/specification.md) and [example](docs/example.md) for examples and mock specification description.
