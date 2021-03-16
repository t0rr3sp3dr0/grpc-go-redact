package test

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

func TestStringTestStruct(t *testing.T) {
	t.Run("Basic Secret Redaction", func(t *testing.T) {
		tStruct := &TestStruct{
			NonSecret: nonSecretVal,
			Secret:    secretVal,
		}

		strVal := fmt.Sprintln(tStruct)

		assert.False(t, strings.Contains(strVal, secretVal), "should not contain secret value")
		assert.True(t, strings.Contains(strVal, "REDACTED"), "should contain redacted string")
		assert.True(t, strings.Contains(strVal, nonSecretVal), "should contain non secret value")
	})

	t.Run("Should still redact empty strings", func(t *testing.T) {
		tStruct := &TestStruct{
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
		tStruct := &TestStruct{
			NonSecret: nonSecretVal,
			Secret:    secretVal,
		}

		list := &TestStructList{
			Data: []*TestStruct{tStruct},
		}

		strVal := fmt.Sprintln(list)

		assert.False(t, strings.Contains(strVal, secretVal), "should not contain secret value")
		assert.True(t, strings.Contains(strVal, "REDACTED"), "should contain redacted string")
		assert.True(t, strings.Contains(strVal, nonSecretVal), "should contain non secret valu")
	})

	t.Run("Should still redact empty strings", func(t *testing.T) {
		tStruct := &TestStruct{
			NonSecret: nonSecretVal,
			Secret:    "",
		}

		list := &TestStructList{
			Data: []*TestStruct{tStruct},
		}

		strVal := fmt.Sprintln(list)

		assert.True(t, strings.Contains(strVal, "REDACTED"), "should contain redacted string")
		assert.True(t, strings.Contains(strVal, nonSecretVal), "should contain  non secret value")
	})

}
