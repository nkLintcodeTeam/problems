package match

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Match(t *testing.T) {
	scenarios := []struct {
		argS     string
		argP     string
		expected bool
	}{
		{
			argS:     "aa",
			argP:     "a",
			expected: false,
		},
		{
			argS:     "aa",
			argP:     "a*",
			expected: true,
		},
	}

	for _, v := range scenarios {
		assert.Equal(t, v.expected, isMatch(v.argS, v.argP))
	}
}
