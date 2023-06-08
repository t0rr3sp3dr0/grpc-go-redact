package test

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/Azure/grpc-go-redact/filehandler"
	"github.com/Azure/grpc-go-redact/generator"
	"github.com/stretchr/testify/assert"
)

const (
	secretVal    = "thisIsASecret"
	nonSecretVal = "thisIsAStandardVal"
)

func TestStringTestEnum(t *testing.T) {
	a := fmt.Sprint(E_A)
	assert.Equal(t, "A", a)

	b := fmt.Sprint(E_B)
	assert.Equal(t, "B", b)
}

func TestStringTestStruct(t *testing.T) {
	t.Run("Basic Secret Redaction", func(t *testing.T) {
		tStruct := &M{
			NonSecret: nonSecretVal,
			Secret:    secretVal,
		}

		strVal := fmt.Sprint(tStruct)

		assert.False(t, strings.Contains(strVal, secretVal), "should not contain secret value")
		assert.True(t, strings.Contains(strVal, "REDACTED"), "should contain redacted string")
		assert.True(t, strings.Contains(strVal, nonSecretVal), "should contain non secret value")
	})

	t.Run("Should still redact empty strings", func(t *testing.T) {
		tStruct := &M{
			NonSecret: nonSecretVal,
			Secret:    "",
		}

		strVal := fmt.Sprint(tStruct)

		assert.True(t, strings.Contains(strVal, "REDACTED"), "should contain redacted string")
		assert.True(t, strings.Contains(strVal, nonSecretVal), "should contain  non secret value")
	})

}

func TestStringTestStructList(t *testing.T) {
	t.Run("Basic Secret Redaction", func(t *testing.T) {
		tStruct := &M{
			NonSecret: nonSecretVal,
			Secret:    secretVal,
		}

		list := &Ms{
			Data: []*M{tStruct},
		}

		strVal := fmt.Sprint(list)

		assert.False(t, strings.Contains(strVal, secretVal), "should not contain secret value")
		assert.True(t, strings.Contains(strVal, "REDACTED"), "should contain redacted string")
		assert.True(t, strings.Contains(strVal, nonSecretVal), "should contain non secret valu")
	})

	t.Run("Should still redact empty strings", func(t *testing.T) {
		tStruct := &M{
			NonSecret: nonSecretVal,
			Secret:    "",
		}

		list := &Ms{
			Data: []*M{tStruct},
		}

		strVal := fmt.Sprint(list)

		assert.True(t, strings.Contains(strVal, "REDACTED"), "should contain redacted string")
		assert.True(t, strings.Contains(strVal, nonSecretVal), "should contain  non secret value")
	})

}

func TestReadAndWrite(t *testing.T) {
	t.Run("Reading And Writing File Keeps Comments", func(t *testing.T) {
		const (
			src = "output.pb.go"
			dst = "output.txt"
		)

		parseInfo, err := filehandler.ParseFile(src)
		assert.Nil(t, err)
		assert.NotNil(t, parseInfo.F.Comments)

		err = generator.GenerateStringFunc(parseInfo)
		assert.Nil(t, err)

		parseInfo.OutputFile = dst

		err = filehandler.WriteASTToFile(parseInfo)
		assert.Nil(t, err)

		baseBytes, err := ioutil.ReadFile(src)
		assert.Nil(t, err)

		writenBytes, err := ioutil.ReadFile(dst)
		assert.Nil(t, err)

		assert.Equal(t, baseBytes, writenBytes)
	})
}

func TestNil(t *testing.T) {
	t.Run("Nil protoreflect.Enum", func(t *testing.T) {
		var e *E

		defer func() {
			r := recover()
			assert.NotNil(t, r)
		}()

		_ = e.String()
		assert.Fail(t, "it should have panicked")
	})

	t.Run("Zero protoreflect.Enum", func(t *testing.T) {
		var e E

		s := e.String()
		assert.NotEqual(t, "{}", s)
	})

	t.Run("Nil protoreflect.ProtoMessage", func(t *testing.T) {
		var m *M

		s := m.String()
		assert.Equal(t, "{}", s)
	})

	t.Run("Zero protoreflect.ProtoMessage", func(t *testing.T) {
		var m M

		s := m.String()
		assert.NotEqual(t, "{}", s)
	})
}
