package nextgen

import (
	"encoding/json"
	"fmt"
)

func (a *DockerAuthenticationDto) UnmarshalJSON(data []byte) error {

	type Alias DockerAuthenticationDto

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
	case DockerAuthTypes.UsernamePassword.String():
		err = json.Unmarshal(aux.Spec, &a.UsernamePassword)
	case DockerAuthTypes.Anonymous.String():
		// nothing to do
	default:
		panic(fmt.Sprintf("unknown docker auth method type %s", a.Type_))
	}

	return err
}

func (a *DockerAuthenticationDto) MarshalJSON() ([]byte, error) {
	type Alias DockerAuthenticationDto

	var spec []byte
	var err error

	switch a.Type_ {
	case DockerAuthTypes.UsernamePassword.String():
		spec, err = json.Marshal(a.UsernamePassword)
	case DockerAuthTypes.Anonymous.String():
		// nothing to do
	default:
		panic(fmt.Sprintf("unknown docker auth method type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
