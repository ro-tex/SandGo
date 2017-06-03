package hello

import "testing"

func TestHello(t *testing.T) {

	tests := []struct {
		input    string // we won't use this one this time but it's good to know how to do it
		expected string
	}{
		{"", "Hello Go!"},
	}

	for _, test := range tests {
		actual := Hello()
		if actual != test.expected {
			t.Errorf("Failed! Expected %s, got %s.", test.expected, actual)
		}
	}
}
