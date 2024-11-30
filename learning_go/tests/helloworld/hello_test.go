package main

import "testing"

// Example of table driven testing.
func TestHello(t *testing.T) {
	testcases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "test success hello world",
			input:    "hello world",
			expected: "hello world",
		},
		{
			name:     "test failing hello world",
			input:    "hello world",
			expected: "failed world",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got := Hello()
			if got != tc.expected {
				t.Errorf("got %q want %q", got, tc.expected)
			}
		})
	}
}
