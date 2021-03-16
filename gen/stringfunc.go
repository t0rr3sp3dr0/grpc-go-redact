package gen

import (
	"encoding/json"

	"github.com/samkreter/redact"
)

type XXX struct {
	Secret    string
	NonSecret string `redact:"nonsecret"`
	SecretPtr *string
}

func (x *XXX) String() string {
	var copy XXX
	jsonBytes, err := json.Marshal(x)
	if err != nil {
		return ""
	}

	err = json.Unmarshal(jsonBytes, &copy)
	if err != nil {
		return ""
	}

	if err := redact.Redact(&copy); err != nil {
		return ""
	}

	jsonBytes, err = json.Marshal(copy)
	if err != nil {
		return ""
	}

	return string(jsonBytes)
}
