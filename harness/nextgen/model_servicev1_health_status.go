/*
 * api/v1/gitops_service.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

type Servicev1HealthStatus string

// List of servicev1HealthStatus
const (
	HEALTH_STATUS_UNSET_Servicev1HealthStatus Servicev1HealthStatus = "HEALTH_STATUS_UNSET"
	HEALTHY_Servicev1HealthStatus             Servicev1HealthStatus = "HEALTHY"
	UNHEALTHY_Servicev1HealthStatus           Servicev1HealthStatus = "UNHEALTHY"
)
