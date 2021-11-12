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

// Connetor Resource along with metadata
type ConnectorResponse struct {
	Connector             *ConnectorInfoDto             `json:"connector,omitempty"`
	CreatedAt             int64                         `json:"createdAt,omitempty"`
	LastModifiedAt        int64                         `json:"lastModifiedAt,omitempty"`
	Status                *ConnectorConnectivityDetails `json:"status,omitempty"`
	ActivityDetails       *ConnectorActivityDetails     `json:"activityDetails,omitempty"`
	HarnessManaged        bool                          `json:"harnessManaged,omitempty"`
	GitDetails            *EntityGitDetails             `json:"gitDetails,omitempty"`
	EntityValidityDetails *EntityValidityDetails        `json:"entityValidityDetails,omitempty"`
}
