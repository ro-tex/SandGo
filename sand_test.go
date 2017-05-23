package main

import (
	"runtime"
	"testing"
)

func Test_VertexString(t *testing.T) {

	// This test is safe for parallel execution with other tests:
	t.Parallel()

	var tests = []struct {
		v        Vertex
		expected string
	}{
		{Vertex{1, 2}, "{X: 1, Y: 2}"}, // happy case
		{Vertex{}, "{X: 0, Y: 0}"},     // empty
	}

	for _, test := range tests {
		actual := test.v.String()
		if actual != test.expected {
			t.Errorf("Expected '%s', got '%s'", test.expected, actual)
		}
	}
}

func Test_Windows(t *testing.T) {

	// We can skip this test in certain conditions (logs "Skipped"):
	if runtime.GOOS != "windows" {
		t.Skip("We need Windows for this test.")
	}

	// Test Windows-dependent functionality...
}
