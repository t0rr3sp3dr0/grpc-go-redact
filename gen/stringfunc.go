package gen

import (
	"encoding/json"
	"reflect"
)

type XXX struct {
	Secret    string
	NotSecret string `nonsecret:"true"`
}

func (x *XXX) String() string {
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
