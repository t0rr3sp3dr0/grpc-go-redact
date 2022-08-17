package gen

import (
	"encoding/json"

	"github.com/samkreter/redact"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

type XXX struct {
	Secret    string
	NonSecret string `redact:"nonsecret"`
	SecretPtr *string
}

func (x *XXX) String() string {
	type EnumType interface {
		Descriptor() protoreflect.EnumDescriptor
		Number() protoreflect.EnumNumber
	}

	enumType, ok := interface{}(x).(EnumType)
	if ok {
		return protoimpl.X.EnumStringOf(enumType.Descriptor(), enumType.Number())
	}

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
