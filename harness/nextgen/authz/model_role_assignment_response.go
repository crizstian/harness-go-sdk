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

type RoleAssignmentResponse struct {
	RoleAssignment *RoleAssignment `json:"roleAssignment,omitempty"`
	Scope          *Scope          `json:"scope,omitempty"`
	CreatedAt      int64           `json:"createdAt,omitempty"`
	LastModifiedAt int64           `json:"lastModifiedAt,omitempty"`
	HarnessManaged bool            `json:"harnessManaged,omitempty"`
}
