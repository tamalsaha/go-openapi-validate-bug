# go-openapi-validate-bug

Bug; https://github.com/go-openapi/validate/issues/137

## How to produce the bug

```
$ go run main.go 
spec.target.rules.sourceHost: Invalid value: "null": spec.target.rules.sourceHost in body must be of type string: "null"
```
