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

import (
	"encoding/json"
)

type ConnectorInfoDto struct {
	Name              string                      `json:"name"`
	Identifier        string                      `json:"identifier"`
	Description       string                      `json:"description,omitempty"`
	OrgIdentifier     string                      `json:"orgIdentifier,omitempty"`
	ProjectIdentifier string                      `json:"projectIdentifier,omitempty"`
	Tags              map[string]string           `json:"tags,omitempty"`
	Type_             string                      `json:"type"`
	Artifactory       *ArtifactoryConnectorDto    `json:"-"`
	AwsCC             *CeAwsConnectorDto          `json:"-"`
	Aws               *AwsConnectorDto            `json:"-"`
	DockerRegistry    *DockerConnectorDto         `json:"-"`
	Gcp               *GcpConnectorDto            `json:"-"`
	Git               *GitConfigDto               `json:"-"`
	HttpHelm          *HttpHelmConnectorDto       `json:"-"`
	K8sCluster        *KubernetesClusterConfigDto `json:"-"`
	Nexus             *NexusConnectorDto          `json:"-"`
	Spec              json.RawMessage             `json:"spec"`
}
