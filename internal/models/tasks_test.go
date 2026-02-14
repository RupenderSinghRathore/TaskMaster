package models

import (
	"reflect"
	"slices"
	"testing"
	"time"
)

func TestInsertionSort(t *testing.T) {
	tests := []struct {
		name     string
		input    Tasks
		expected Tasks
	}{
		{
			name:     "empty",
			input:    Tasks{},
			expected: Tasks{},
		},
		{
			name: "generic",
			input: Tasks{
				{Deadline: time.Now().Add(24 * time.Hour).Round(time.Second)},
				{Deadline: time.Now().Add(14 * time.Hour).Round(time.Second)},
				{Deadline: time.Now().Add(29 * time.Hour).Round(time.Second)},
				{Deadline: time.Now().Add(4 * time.Hour).Round(time.Second)},
				{Deadline: time.Now().Add(time.Hour).Round(time.Second)},
			},
			expected: Tasks{
				{Deadline: time.Now().Add(time.Hour).Round(time.Second)},
				{Deadline: time.Now().Add(4 * time.Hour).Round(time.Second)},
				{Deadline: time.Now().Add(14 * time.Hour).Round(time.Second)},
				{Deadline: time.Now().Add(24 * time.Hour).Round(time.Second)},
				{Deadline: time.Now().Add(29 * time.Hour).Round(time.Second)},
			},
		},
		{
			name: "already sorted",
			input: Tasks{
				{Deadline: time.Now().Add(time.Hour).Round(time.Second)},
				{Deadline: time.Now().Add(4 * time.Hour).Round(time.Second)},
				{Deadline: time.Now().Add(14 * time.Hour).Round(time.Second)},
				{Deadline: time.Now().Add(24 * time.Hour).Round(time.Second)},
				{Deadline: time.Now().Add(29 * time.Hour).Round(time.Second)},
			},
			expected: Tasks{
				{Deadline: time.Now().Add(time.Hour).Round(time.Second)},
				{Deadline: time.Now().Add(4 * time.Hour).Round(time.Second)},
				{Deadline: time.Now().Add(14 * time.Hour).Round(time.Second)},
				{Deadline: time.Now().Add(24 * time.Hour).Round(time.Second)},
				{Deadline: time.Now().Add(29 * time.Hour).Round(time.Second)},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := slices.Clone(tt.input)
			got.InsertionSort()
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf(
					"input: %v, expected: %v, got: %v",
					tt.input,
					tt.expected,
					got,
				)
				return
			}
		})
	}
}

func TestAppend(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected Task
	}{{
		name:  "any string",
		input: "mailtrap",
		expected: Task{
			Title:    "mailtrap",
			Deadline: time.Now().Add(24 * time.Hour).Round(time.Second),
			Status:   Pending,
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tasks := Tasks{}

			got := *tasks.Append(tt.input)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("input: %v, expected: %v, got: %v", tt.input, tt.expected, got)
			}
		})
	}
}
