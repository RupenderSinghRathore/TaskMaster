package main

import (
	"RupenderSinghRathore/TaskMaster/internal/models"
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
			expected := time.Now().Add(tt.expectedShift).Round(time.Second)
			got, err := getDeadline(tt.input)
			got = got.Round(time.Second)

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

			if !reflect.DeepEqual(expected, got) {
				diff := got.Sub(expected)
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

func TestGetTimeperiod(t *testing.T) {
	tests := []struct {
		name       string
		inputShift time.Duration
		expected   string
	}{
		{
			name:       "negative duration",
			inputShift: -24 * time.Hour,
			expected:   "💀",
		},
		{
			name:       "integer duration",
			inputShift: 2 * 24 * time.Hour,
			expected:   "2.0d",
		},
		{
			name:       "border time",
			inputShift: 24 * time.Hour,
			expected:   "24.0h",
		},
		{
			name:       "float duration",
			inputShift: 45 * 24 * time.Hour,
			expected:   "1.5m",
		},
		{
			name:       "hours",
			inputShift: 4 * time.Hour,
			expected:   "4.0h",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := time.Now().Add(tt.inputShift)
			got := getTimeperiod(input)

			if !reflect.DeepEqual(tt.expected, got) {
				t.Errorf("input: %v, expected: %v, got: %v", tt.inputShift, tt.expected, got)
			}
		})
	}
}

func TestGetTaskId(t *testing.T) {
	mockTasks := models.Tasks{
		{Title: "sleep"},
		{Title: "go fuck yourself"},
		{Title: "sleep"},
		{Title: "go fuck yourself"},
	}

	tests := []struct {
		name          string
		input         string
		expected      int
		expectedError bool
	}{
		{
			name:     "valid int",
			input:    "3",
			expected: 2,
		},
		{
			name:          "zero",
			input:         "0",
			expectedError: true,
		},
		{
			name:          "zero",
			input:         "-1",
			expectedError: true,
		},
		{
			name:          "invalid int",
			input:         "20fuck",
			expectedError: true,
		},
		{
			name:          "grater then len",
			input:         "20",
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := application{tasks: mockTasks}
			got, err := app.getTaskId(tt.input)

			if tt.expectedError {
				if err == nil {
					t.Errorf(
						"expected error got none; input: %v, expected: %v, got: %v",
						tt.input,
						tt.expected,
						got,
					)
				}
				return
			}
			if err != nil {
				t.Errorf(
					"unexpected error: %v, for input: input: %v, expected: %v",
					err,
					tt.input,
					tt.expected,
				)
			}
			if !reflect.DeepEqual(tt.expected, got) {
				t.Errorf(
					"input: input: %v, expected: %v, got: %v",
					tt.input,
					tt.expected,
					got,
				)
			}
		})
	}
}
