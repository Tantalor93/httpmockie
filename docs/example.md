# Examples
For running the tool you have to specify mock specifications (format described in [specification](specification.md))
and point the tool to a directory containing those specs (`--dir` flag). `httpmockie` will then expose endpoints on configurable
port (flag `--port`)

Running this example after cloning `httpmockie` repository from the project root:
```
httpmockie --port 8081 --dir "docs/examples"
```
Will expose four endpoints based on [example specifications](examples/):
```
/example.bytes
/example.json
/example-delayed.json
/example
```

These endpoints can be then accessed for example using `curl`
```
curl localhost:8081/example.json -i
```
