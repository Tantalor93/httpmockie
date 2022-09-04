# Specification
`httpmockie` creates endpoints based on JSON specifications comforming to [JSON schema](../internal/mockspec/specification.json).
Each endpoint specification is placed into separate file (single file can contain only single JSON specification).

single JSON endpoint specification must contain fields:
* `path` = HTTP path to expose the HTTP endpoint on

additionally the endpoint can be further configured using these fields:
* `status` = default is HTTP status OK (200)
* `body` = HTTP response body in text format
* `jsonBody` = HTTP response body in JSON format
* `base64Body` = HTTP response body encoded in Base64 (useful for returning arbitrary non-textual data)
* `headers` = HTTP headers of response
* `delay` = configures delay between responding to client request (useful for setting up advanced scenarios)
  * `durationMs` = static delay in milliseconds
  * `deviationMs` = when used together with `durationMs` it creates non-static delay before responding to client request,
                    the client receives answer after delay in interval `[durationMs-deviationMs,durationMs+deviationMs]`
