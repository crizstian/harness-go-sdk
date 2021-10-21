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

type KubernetesOpenIdConnectDto struct {
	OidcIssuerUrl   string `json:"oidcIssuerUrl,omitempty"`
	OidcUsername    string `json:"oidcUsername,omitempty"`
	OidcUsernameRef string `json:"oidcUsernameRef,omitempty"`
	OidcClientIdRef string `json:"oidcClientIdRef"`
	OidcPasswordRef string `json:"oidcPasswordRef"`
	OidcSecretRef   string `json:"oidcSecretRef,omitempty"`
	OidcScopes      string `json:"oidcScopes,omitempty"`
}
