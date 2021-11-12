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

type GitSyncFolderConfigDto struct {
	RootFolder string `json:"rootFolder,omitempty"`
	IsDefault  bool   `json:"isDefault,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	Enabled    bool   `json:"enabled,omitempty"`
}
