## GO HTTP MOCK
simple standalone high-performance mock HTTP server implemented in GO, which is able to provide mocked responses defined by 
mock specifications similar to [Wiremock JSON API](https://wiremock.org/docs/stubbing/).

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
