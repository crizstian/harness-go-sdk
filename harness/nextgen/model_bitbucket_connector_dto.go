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

type BitbucketConnectorDto struct {
	Url               string                      `json:"url"`
	ValidationRepo    string                      `json:"validationRepo,omitempty"`
	Authentication    *BitbucketAuthenticationDto `json:"authentication"`
	ApiAccess         *BitbucketApiAccessDto      `json:"apiAccess,omitempty"`
	DelegateSelectors []string                    `json:"delegateSelectors,omitempty"`
	Type_             string                      `json:"type"`
}
