package test

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/samkreter/grpc-go-redact/filehandler"
	"github.com/samkreter/grpc-go-redact/generator"
	"github.com/stretchr/testify/assert"
)

const (
	secretVal    = "thisIsASecret"
	nonSecretVal = "thisIsAStandardVal"
)

func TestStringTestStruct(t *testing.T) {
	t.Run("Basic Secret Redaction", func(t *testing.T) {
		tStruct := &X{
			NonSecret: nonSecretVal,
			Secret:    secretVal,
		}

		strVal := fmt.Sprintln(tStruct)

		assert.False(t, strings.Contains(strVal, secretVal), "should not contain secret value")
		assert.True(t, strings.Contains(strVal, "REDACTED"), "should contain redacted string")
		assert.True(t, strings.Contains(strVal, nonSecretVal), "should contain non secret value")
	})

	t.Run("Should still redact empty strings", func(t *testing.T) {
		tStruct := &X{
			NonSecret: nonSecretVal,
			Secret:    "",
		}

		strVal := fmt.Sprintln(tStruct)

		assert.True(t, strings.Contains(strVal, "REDACTED"), "should contain redacted string")
		assert.True(t, strings.Contains(strVal, nonSecretVal), "should contain  non secret value")
	})

}

func TestStringTestStructList(t *testing.T) {
	t.Run("Basic Secret Redaction", func(t *testing.T) {
		tStruct := &X{
			NonSecret: nonSecretVal,
			Secret:    secretVal,
		}

		list := &Xs{
			Data: []*X{tStruct},
		}

		strVal := fmt.Sprintln(list)

		assert.False(t, strings.Contains(strVal, secretVal), "should not contain secret value")
		assert.True(t, strings.Contains(strVal, "REDACTED"), "should contain redacted string")
		assert.True(t, strings.Contains(strVal, nonSecretVal), "should contain non secret valu")
	})

	t.Run("Should still redact empty strings", func(t *testing.T) {
		tStruct := &X{
			NonSecret: nonSecretVal,
			Secret:    "",
		}

		list := &Xs{
			Data: []*X{tStruct},
		}

		strVal := fmt.Sprintln(list)

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
