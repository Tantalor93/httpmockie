# Examples
`httpmockie` is a simple CLI tool for fast setup of mock HTTP servers

For running the tool you have to specify mock specifications (format described in [specification section](specification.md))
and point the tool to a directory containing those specs (`--dir` flag). `httpmockie` will then expose endpoints on configurable
port (flag `--port`)

Running this example:
```
httpmockie --port 8081 --dir "docs/examples"
```
Will expose three endpoints
```
/example.bytes
/example.json
/example
```

Which can be then accessed for example using `curl`
```
curl localhost:8081/example.json -i
```
