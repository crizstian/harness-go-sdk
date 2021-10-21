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

type CiModuleLicenseDto struct {
	Id string `json:"id,omitempty"`
	AccountIdentifier string `json:"accountIdentifier,omitempty"`
	ModuleType string `json:"moduleType,omitempty"`
	Edition string `json:"edition,omitempty"`
	LicenseType string `json:"licenseType,omitempty"`
	Status string `json:"status,omitempty"`
	StartTime int64 `json:"startTime,omitempty"`
	ExpiryTime int64 `json:"expiryTime,omitempty"`
	CreatedAt int64 `json:"createdAt,omitempty"`
	LastModifiedAt int64 `json:"lastModifiedAt,omitempty"`
	NumberOfCommitters int32 `json:"numberOfCommitters,omitempty"`
}
