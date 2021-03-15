package gen

import (
	"encoding/json"
	"reflect"

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

func (x *XXX) Test() string {
	ifv := reflect.ValueOf(x)

	// ensure we use a pointer
	if ifv.Kind() != reflect.Ptr {
		ifv = reflect.ValueOf(&x)
	}

	ift := reflect.Indirect(ifv).Type()
	if ift.Kind() != reflect.Struct {
		return ""
	}

	for i := 0; i < ift.NumField(); i++ {
		v := ift.Field(i)
		el := reflect.Indirect(ifv.Elem().FieldByName(v.Name))

		// Currently only support redacting strings
		switch el.Kind() {
		case reflect.String:
			if el.CanSet() {
				tag := v.Tag.Get("nonsecret")
				if tag != "true" && el.String() != "" {
					el.SetString("REDACTED")
				}
			}
		}
	}

	b, _ := json.Marshal(x)
	return string(b)
}
