package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var inputs = []struct {
	n      int
	nums   []string
	result string
}{
	{3, []string{"123", "-129", "221"}, "221 123 -129"},
	{4, []string{"21", "3", "81", "27"}, "3 21 27 81"},
}

func TestSolution(t *testing.T) {
	for _, in := range inputs {
		assert.Equal(t, in.result, solution(in.nums))
	}
}
