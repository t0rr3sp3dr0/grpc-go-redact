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
	cases := []struct {
		Name              string
		SecretVal         string
		NonSecretVal      string
		containsSecret    bool
		containsNonSecret bool
	}{
		{
			Name:              "Basic Secret Redaction",
			SecretVal:         secretVal,
			NonSecretVal:      nonSecretVal,
			containsNonSecret: true,
			containsSecret:    false,
		},
		{
			Name:              "Should not redact empty strings",
			SecretVal:         "",
			NonSecretVal:      nonSecretVal,
			containsNonSecret: true,
			containsSecret:    true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			tStruct := &XXX{
				NotSecret: tc.NonSecretVal,
				Secret:    tc.SecretVal,
			}

			strVal := fmt.Sprintln(tStruct)

			assert.Equal(t, tc.containsSecret, strings.Contains(strVal, tc.SecretVal),
				"should contain expected secret valu")

			if tc.containsSecret {
				assert.False(t, strings.Contains(strVal, "REDACTED"), "should not contain redacted string")
			} else {
				assert.True(t, strings.Contains(strVal, "REDACTED"), "should contain redacted string")
			}

			assert.Equal(t, tc.containsNonSecret, strings.Contains(strVal, tc.NonSecretVal),
				"should contain expected non secret valu")
		})
	}
}
