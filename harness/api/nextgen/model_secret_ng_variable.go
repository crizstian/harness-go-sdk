/*
 * CD NextGen API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

type SecretNgVariable struct {
	Description string `json:"description,omitempty"`
	CurrentValue *ParameterFieldObject `json:"currentValue,omitempty"`
	Required bool `json:"required,omitempty"`
	Name string `json:"name,omitempty"`
	Type_ string `json:"type,omitempty"`
	Metadata string `json:"metadata,omitempty"`
	Value *ParameterFieldSecretRefData `json:"value"`
	Default_ string `json:"default,omitempty"`
}
