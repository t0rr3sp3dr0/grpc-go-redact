package internal

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	secretVal    = "thisIsASecret"
	nonSecretVal = "thisIsAStandardVal"
)

var (
	secretPtrVal    = "thisIsAPtrSecret"
	nonSecretPtrVal = "thisIsAPtrStandardVal"
)

func TestString(t *testing.T) {
	t.Run("Basic Secret Redaction", func(t *testing.T) {
		tStruct := &M{
			Secret:       secretVal,
			NonSecret:    nonSecretVal,
			SecretPtr:    &secretPtrVal,
			NonSecretPtr: &nonSecretPtrVal,
		}

		strVal := fmt.Sprintln(tStruct)
		println(strVal)

		assert.False(t, strings.Contains(strVal, secretVal), "should not contain secret value")
		assert.False(t, strings.Contains(strVal, secretPtrVal), "should not contain ptr secret value")
		assert.True(t, strings.Contains(strVal, "REDACTED"), "should contain redacted string")
		assert.True(t, strings.Contains(strVal, nonSecretVal), "should contain non secret value")
		assert.True(t, strings.Contains(strVal, nonSecretPtrVal), "should contain ptr non secret value")
	})

	t.Run("Should still redact empty strings", func(t *testing.T) {
		tStruct := &M{
			NonSecret: nonSecretVal,
			Secret:    "",
		}

		strVal := fmt.Sprintln(tStruct)

		assert.True(t, strings.Contains(strVal, "REDACTED"), "should contain redacted string")
		assert.True(t, strings.Contains(strVal, nonSecretVal), "should contain  non secret value")
	})

}
