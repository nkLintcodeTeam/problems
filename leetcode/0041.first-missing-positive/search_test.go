package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_seach(t *testing.T) {
	scenarios := []struct {
		arg      []int
		expected int
	}{
		{
			arg:      []int{-1, 1, 100, 200},
			expected: 2,
		},
		{
			arg:      []int{-1, 300, 100, 200},
			expected: 1,
		},
		{
			arg:      []int{-1, -2, -3},
			expected: 1,
		},
	}
	for _, v := range scenarios {
		assert.Equal(t, v.expected, firstMissingPositive(v.arg))
	}
}
