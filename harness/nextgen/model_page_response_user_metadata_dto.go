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

type PageResponseUserMetadataDto struct {
	TotalPages    int64             `json:"totalPages,omitempty"`
	TotalItems    int64             `json:"totalItems,omitempty"`
	PageItemCount int64             `json:"pageItemCount,omitempty"`
	PageSize      int64             `json:"pageSize,omitempty"`
	Content       []UserMetadataDto `json:"content,omitempty"`
	PageIndex     int64             `json:"pageIndex,omitempty"`
	Empty         bool              `json:"empty,omitempty"`
}
