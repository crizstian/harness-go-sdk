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

// This is the view of a feature restriction metadata object defined in Harness
type FeatureRestrictionMetadata struct {
	Name                string                            `json:"name,omitempty"`
	ModuleType          string                            `json:"moduleType,omitempty"`
	Edition             string                            `json:"edition,omitempty"`
	RestrictionMetadata map[string]RestrictionMetadataDto `json:"restrictionMetadata,omitempty"`
}
