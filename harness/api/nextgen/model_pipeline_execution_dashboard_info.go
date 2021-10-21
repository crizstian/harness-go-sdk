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

type PipelineExecutionDashboardInfo struct {
	Name string `json:"name,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	OrgIdentifier string `json:"orgIdentifier,omitempty"`
	ProjectIdentifier string `json:"projectIdentifier,omitempty"`
	AccountIdentifier string `json:"accountIdentifier,omitempty"`
	PlanExecutionId string `json:"planExecutionId,omitempty"`
	StartTs int64 `json:"startTs,omitempty"`
}
