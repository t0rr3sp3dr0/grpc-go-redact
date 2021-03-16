package gen

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

func TestString(t *testing.T) {
	t.Run("Basic Secret Redaction", func(t *testing.T) {
		tStruct := &XXX{
			NonSecret: nonSecretVal,
			Secret:    secretVal,
		}

		strVal := fmt.Sprintln(tStruct)

		assert.False(t, strings.Contains(strVal, secretVal), "should not contain secret value")
		assert.True(t, strings.Contains(strVal, "REDACTED"), "should contain redacted string")
		assert.True(t, strings.Contains(strVal, nonSecretVal), "should contain non secret value")
	})

	t.Run("Should still redact empty strings", func(t *testing.T) {
		tStruct := &XXX{
			NonSecret: nonSecretVal,
			Secret:    "",
		}

		strVal := fmt.Sprintln(tStruct)

		assert.True(t, strings.Contains(strVal, "REDACTED"), "should contain redacted string")
		assert.True(t, strings.Contains(strVal, nonSecretVal), "should contain  non secret value")
	})

}
