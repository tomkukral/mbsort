package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPosition(t *testing.T) {

	ts := []string{"a", "b", "c", "abc", "123"}

	td := []struct {
		value    string
		expected int
	}{
		{"b", 1},
		{"abc", 3},
		{"lemur", -1},
	}

	for _, n := range td {
		assert.Equal(t, n.expected, getPosition(ts, n.value), "Wrong position returned")
	}
}
