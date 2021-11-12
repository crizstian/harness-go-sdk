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

type UserInfo struct {
	Uuid                           string                     `json:"uuid,omitempty"`
	Name                           string                     `json:"name,omitempty"`
	Email                          string                     `json:"email,omitempty"`
	Token                          string                     `json:"token,omitempty"`
	DefaultAccountId               string                     `json:"defaultAccountId,omitempty"`
	Intent                         string                     `json:"intent,omitempty"`
	Accounts                       []GatewayAccountRequestDto `json:"accounts,omitempty"`
	Admin                          bool                       `json:"admin,omitempty"`
	TwoFactorAuthenticationEnabled bool                       `json:"twoFactorAuthenticationEnabled,omitempty"`
	EmailVerified                  bool                       `json:"emailVerified,omitempty"`
	Locked                         bool                       `json:"locked,omitempty"`
	SignupAction                   string                     `json:"signupAction,omitempty"`
	Edition                        string                     `json:"edition,omitempty"`
	BillingFrequency               string                     `json:"billingFrequency,omitempty"`
	UtmInfo                        *UtmInfo                   `json:"utmInfo,omitempty"`
}
