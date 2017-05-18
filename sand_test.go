package main

import "testing"

func Test_VertexString(t *testing.T) {

	t.Parallel() // this test is safe for parallel execution with other tests

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
