package main

import (
	"github.com/go-openapi/spec"
	"github.com/go-openapi/validate"
	"github.com/go-openapi/strfmt"
)


func main() {
	openapiSchema := &spec.Schema{
		SchemaProps: spec.SchemaProps{
			ID: "",
			Type: spec.StringOrArray{
				"object",
			},
			Properties: map[string]spec.Schema{
				"astr": spec.Schema{
					SchemaProps: spec.SchemaProps{
						Type: spec.StringOrArray{
							"string",
						},
					},
				},
			},
		},
	}
	validator := validate.NewSchemaValidator(openapiSchema, nil, "", strfmt.Default)
	result := validator.Validate(map[string]interface{}{
		"astr": "",
	})
	if len(result.Errors) > 0 {
		panic(result.Errors)
	}
}
