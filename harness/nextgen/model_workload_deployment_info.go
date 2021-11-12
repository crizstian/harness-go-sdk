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

type WorkloadDeploymentInfo struct {
	ServiceName               string                  `json:"serviceName,omitempty"`
	ServiceId                 string                  `json:"serviceId,omitempty"`
	LastExecuted              *LastWorkloadInfo       `json:"lastExecuted,omitempty"`
	DeploymentTypeList        []string                `json:"deploymentTypeList,omitempty"`
	TotalDeployments          int64                   `json:"totalDeployments,omitempty"`
	TotalDeploymentChangeRate float64                 `json:"totalDeploymentChangeRate,omitempty"`
	PercentSuccess            float64                 `json:"percentSuccess,omitempty"`
	RateSuccess               float64                 `json:"rateSuccess,omitempty"`
	FailureRate               float64                 `json:"failureRate,omitempty"`
	FailureRateChangeRate     float64                 `json:"failureRateChangeRate,omitempty"`
	Frequency                 float64                 `json:"frequency,omitempty"`
	FrequencyChangeRate       float64                 `json:"frequencyChangeRate,omitempty"`
	LastPipelineExecutionId   string                  `json:"lastPipelineExecutionId,omitempty"`
	Workload                  []WorkloadDateCountInfo `json:"workload,omitempty"`
}
