package internal

import (
	"encoding/json"
	"reflect"

	"github.com/Azure/redact"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

type M struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Secret       string  `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret,omitempty"`
	NonSecret    string  `protobuf:"bytes,2,opt,name=nonSecret,proto3" json:"nonSecret,omitempty" redact:"nonsecret"`
	SecretPtr    *string `protobuf:"bytes,3,opt,name=secretPtr,proto3,oneof" json:"secretPtr,omitempty"`
	NonSecretPtr *string `protobuf:"bytes,4,opt,name=nonSecretPtr,proto3,oneof" json:"nonSecretPtr,omitempty" redact:"nonsecret"`
}

func (x *M) String() string {
	i := interface{}(x)

	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Pointer && v.IsNil() {
		return "{}"
	}

	switch x := i.(type) {
	case protoreflect.Enum:
		return protoimpl.X.EnumStringOf(x.Descriptor(), x.Number())

	case protoreflect.ProtoMessage:
		clone := proto.Clone(x)
		if err := redact.Redact(clone); err != nil {
			return ""
		}
		bytes, err := json.Marshal(clone)
		if err != nil {
			return ""
		}
		return string(bytes)

	default:
		return ""
	}
}
