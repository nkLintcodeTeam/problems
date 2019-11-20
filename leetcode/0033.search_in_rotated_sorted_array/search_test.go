package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Search(t *testing.T) {
	scenarios := []struct {
		argS      []int
		argTarget int
		expected  int
	}{
		{
			argS:      []int{4, 5, 6, 0, 1, 2, 3},
			argTarget: 5,
			expected:  1,
		},
		{
			argS:      []int{4, 5, 6, 7, 0, 1, 2},
			argTarget: 4,
			expected:  0,
		},
		{
			argS:      []int{3, 1},
			argTarget: 1,
			expected:  1,
		},
	}

	for _, v := range scenarios {
		assert.Equal(t, v.expected, search(v.argS, v.argTarget))
	}
}
