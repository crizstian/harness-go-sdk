/*
 * api/v1/gitops_service.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

type V1ConnectedStatus string

// List of v1ConnectedStatus
const (
	CONNECTED_STATUS_UNSET_V1ConnectedStatus V1ConnectedStatus = "CONNECTED_STATUS_UNSET"
	CONNECTED_V1ConnectedStatus              V1ConnectedStatus = "CONNECTED"
	DISCONNECTED_V1ConnectedStatus           V1ConnectedStatus = "DISCONNECTED"
)
