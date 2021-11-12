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

type AwsCurAttributesDto struct {
	ReportName   string `json:"reportName"`
	S3BucketName string `json:"s3BucketName"`
	Region       string `json:"region,omitempty"`
	S3Prefix     string `json:"s3Prefix,omitempty"`
}
