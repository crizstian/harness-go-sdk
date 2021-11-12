/*
 * Access Control API Reference
 *
 * This is the Open Api Spec 3 for the Access Control Service. This is under active development. Beware of the breaking change with respect to the generated code stub
 *
 * API version: 1.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package authz

type ResponseDtoFeatureRestrictionUsageDto struct {
	Status        string                      `json:"status,omitempty"`
	Data          *FeatureRestrictionUsageDto `json:"data,omitempty"`
	MetaData      *interface{}                `json:"metaData,omitempty"`
	CorrelationId string                      `json:"correlationId,omitempty"`
}
