package main

import (
	"reflect"
	"testing"
)

func TestCapitalize(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "single word",
			input:    "helper",
			expected: "Helper",
		},
		{
			name:     "sentence",
			input:    "all women are obstacles",
			expected: "All women are obstacles",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := capitalize(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf(
					"test %v failed: input: %v, expected: %v, got: %v",
					tt.name,
					tt.input,
					tt.expected,
					got,
				)
			}
		})
	}
}
