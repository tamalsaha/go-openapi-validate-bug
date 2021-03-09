package main

import (
	_ "embed"
	"fmt"

	apiextensionsinternal "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	apiservervalidation "k8s.io/apiextensions-apiserver/pkg/apiserver/validation"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/yaml"
)

var (
	//go:embed cr.yaml
	cr string

	//go:embed crd.v1.yaml
	crd string
)

func main() {
	var crd1 apiextensionsv1.CustomResourceDefinition
	err := yaml.Unmarshal([]byte(crd), &crd1)
	if err != nil {
		panic(err)
	}
	validationSchema := crd1.Spec.Versions[0].Schema

	var internalValidationSchema *apiextensionsinternal.CustomResourceValidation
	if validationSchema != nil {
		internalValidationSchema = &apiextensionsinternal.CustomResourceValidation{}
		if err := apiextensionsv1.Convert_v1_CustomResourceValidation_To_apiextensions_CustomResourceValidation(validationSchema, internalValidationSchema, nil); err != nil {
			panic(fmt.Errorf("failed to convert CRD validation to internal version: %v", err))
		}
	}

	var u unstructured.Unstructured
	err = yaml.Unmarshal([]byte(cr), &u)
	if err != nil {
		panic(err)
	}

	schemaValidator, _, err := apiservervalidation.NewSchemaValidator(internalValidationSchema)
	if err != nil {
		panic(err)
	}

	errlist := apiservervalidation.ValidateCustomResource(nil, u.UnstructuredContent(), schemaValidator)
	fmt.Println(errlist.ToAggregate())
}
