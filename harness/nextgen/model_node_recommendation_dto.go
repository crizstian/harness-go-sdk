/*
 * CD NextGen API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

type NodeRecommendationDto struct {
	Id                  string                   `json:"id,omitempty"`
	NodePoolId          *NodePoolId              `json:"nodePoolId,omitempty"`
	ResourceRequirement *RecommendClusterRequest `json:"resourceRequirement,omitempty"`
	Current             *RecommendationResponse  `json:"current,omitempty"`
	Recommended         *RecommendationResponse  `json:"recommended,omitempty"`
}