/*
 * CD NextGen API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

// Contains list of Entities based on Entity Type
type GitSyncEntityList struct {
	EntityType      string          `json:"entityType,omitempty"`
	Count           int64           `json:"count,omitempty"`
	GitSyncEntities []GitSyncEntity `json:"gitSyncEntities,omitempty"`
}