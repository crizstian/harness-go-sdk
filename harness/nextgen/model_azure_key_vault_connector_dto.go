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

type AzureKeyVaultConnectorDto struct {
	ClientId             string   `json:"clientId"`
	SecretKey            string   `json:"secretKey"`
	TenantId             string   `json:"tenantId"`
	VaultName            string   `json:"vaultName"`
	Subscription         string   `json:"subscription"`
	IsDefault            bool     `json:"isDefault,omitempty"`
	AzureEnvironmentType string   `json:"azureEnvironmentType,omitempty"`
	DelegateSelectors    []string `json:"delegateSelectors,omitempty"`
	Default_             bool     `json:"default,omitempty"`
}
