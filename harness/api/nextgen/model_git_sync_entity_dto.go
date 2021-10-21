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

type GitSyncEntityDto struct {
	EntityName string `json:"entityName,omitempty"`
	EntityType string `json:"entityType,omitempty"`
	EntityIdentifier string `json:"entityIdentifier,omitempty"`
	GitConnectorId string `json:"gitConnectorId,omitempty"`
	RepoUrl string `json:"repoUrl,omitempty"`
	Branch string `json:"branch,omitempty"`
	FolderPath string `json:"folderPath,omitempty"`
	EntityGitPath string `json:"entityGitPath,omitempty"`
	RepoProviderType string `json:"repoProviderType,omitempty"`
	EntityReference *EntityReference `json:"entityReference,omitempty"`
	AccountId string `json:"accountId,omitempty"`
}
