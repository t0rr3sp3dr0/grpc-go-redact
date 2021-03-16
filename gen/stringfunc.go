package gen

import (
	"encoding/json"

	"github.com/samkreter/redact"
)

type XXX struct {
	Secret    string
	NonSecret string `redact:"nonsecret"`
}

func (x *XXX) String() string {
	var copy XXX
	jsonBytes, err := json.Marshal(x)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(jsonBytes, &copy)
	if err != nil {
		panic(err)
	}

	if err := redact.Redact(&copy); err != nil {
		panic(err)
	}

	jsonBytes, err = json.Marshal(copy)
	if err != nil {
		panic(err)
	}

	return string(jsonBytes)
}
