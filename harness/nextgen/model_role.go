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

type Role struct {
	Identifier         string            `json:"identifier,omitempty"`
	Name               string            `json:"name,omitempty"`
	Permissions        []string          `json:"permissions,omitempty"`
	AllowedScopeLevels []string          `json:"allowedScopeLevels,omitempty"`
	Description        string            `json:"description,omitempty"`
	Tags               map[string]string `json:"tags,omitempty"`
}
