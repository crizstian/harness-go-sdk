package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *ConnectorInfoDto) UnmarshalJSON(data []byte) error {

	type Alias ConnectorInfoDto

	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(a),
	}

	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}

	switch a.Type_ {
	case ConnectorTypes.Artifactory.String():
		err = json.Unmarshal(aux.Spec, &a.Artifactory)
	case ConnectorTypes.Aws.String():
		err = json.Unmarshal(aux.Spec, &a.Aws)
	case ConnectorTypes.CEAws.String():
		err = json.Unmarshal(aux.Spec, &a.AwsCC)
	case ConnectorTypes.DockerRegistry.String():
		err = json.Unmarshal(aux.Spec, &a.DockerRegistry)
	case ConnectorTypes.Gcp.String():
		err = json.Unmarshal(aux.Spec, &a.Gcp)
	case ConnectorTypes.Git.String():
		err = json.Unmarshal(aux.Spec, &a.Git)
	case ConnectorTypes.HttpHelmRepo.String():
		err = json.Unmarshal(aux.Spec, &a.HttpHelm)
	case ConnectorTypes.K8sCluster.String():
		err = json.Unmarshal(aux.Spec, &a.K8sCluster)
	case ConnectorTypes.Nexus.String():
		err = json.Unmarshal(aux.Spec, &a.Nexus)
	default:
		panic(fmt.Sprintf("unknown connector type %s", a.Type_))
	}

	return err
}

func (a *ConnectorInfoDto) MarshalJSON() ([]byte, error) {
	type Alias ConnectorInfoDto

	var spec []byte
	var err error

	switch a.Type_ {
	case ConnectorTypes.Artifactory.String():
		spec, err = json.Marshal(a.Artifactory)
	case ConnectorTypes.Aws.String():
		spec, err = json.Marshal(a.Aws)
	case ConnectorTypes.CEAws.String():
		spec, err = json.Marshal(a.AwsCC)
	case ConnectorTypes.DockerRegistry.String():
		spec, err = json.Marshal(a.DockerRegistry)
	case ConnectorTypes.Gcp.String():
		spec, err = json.Marshal(a.Gcp)
	case ConnectorTypes.Git.String():
		spec, err = json.Marshal(a.Git)
	case ConnectorTypes.HttpHelmRepo.String():
		spec, err = json.Marshal(a.HttpHelm)
	case ConnectorTypes.K8sCluster.String():
		spec, err = json.Marshal(a.K8sCluster)
	case ConnectorTypes.Nexus.String():
		spec, err = json.Marshal(a.Nexus)
	default:
		panic(fmt.Sprintf("unknown connector type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
