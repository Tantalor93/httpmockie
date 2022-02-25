# Specification structure
Specification for `go-http-mock` is given as a JSON file described by schema:
```
{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "type": "object",
  "properties": {
    "path": {
      "type": "string"
    },
    "status": {
      "type": "integer"
    },
    "body": {
      "type": "string"
    },
    "jsonBody": {
      "type": "object"
    },
    "base64Body":{
      "type": "string"
    },
    "headers": {
      "type": "object" 
    }
  },
  "required": [
    "path"
  ]
}
```

