package main

import (
	"fmt"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/validate"
	"github.com/go-openapi/strfmt"
)

func main2() {
	openapiSchema := &spec.Schema{
		SchemaProps: spec.SchemaProps{
			ID: "",
			Type: spec.StringOrArray{
				"string",
			},
		},
	}
	validator := validate.NewSchemaValidator(openapiSchema, nil, "", strfmt.Default)
	result := validator.Validate("")
	fmt.Println(result.AsError())
}

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
	fmt.Println(result.AsError())
}
