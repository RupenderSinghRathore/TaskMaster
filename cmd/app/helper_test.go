package main

import (
	"reflect"
	"testing"
	"time"
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
					"input: %v, expected: %v, got: %v",
					tt.input,
					tt.expected,
					got,
				)
			}
		})
	}
}

func TestGetDeadline(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedShift time.Duration
		expectedError bool
	}{
		{
			name:          "int duration",
			input:         "2d",
			expectedShift: 2 * 24 * time.Hour,
			expectedError: false,
		},
		{
			name:          "float duration",
			input:         "1.5d",
			expectedShift: 36 * time.Hour,
			expectedError: false,
		},
		{
			name:          "invalid duration",
			input:         "2hf",
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expected := time.Now().Add(tt.expectedShift)
			got, err := getDeadline(tt.input)
			if tt.expectedError {
				if err == nil {
					t.Errorf(
						"expected error got none; input: %v, expected: %v, got: %v",
						tt.input,
						expected.String(),
						got,
					)
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			diff := got.Sub(expected)
			if diff < 0 {
				diff = -diff
			}

			if diff > time.Second {
				t.Errorf(
					"input: %v, expected: %v, got: %v, diff: %v",
					tt.input,
					expected.Format(time.DateTime),
					got.Format(time.DateTime),
					time.Duration(diff),
				)
			}
		})
	}
}
