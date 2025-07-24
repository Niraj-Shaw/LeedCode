package main

import "testing"

func TestFindLongestPrefix(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected string
	}{
		{"Common prefix 'fl'", []string{"flower", "flow", "flight"}, "fl"},
		{"Common prefix 'do'", []string{"dog", "doctor", "dot"}, "do"},
		{"No common prefix", []string{"apple", "banana", "cherry"}, ""},
		{"Single word", []string{"single"}, "single"},
		{"All same words", []string{"same", "same", "same"}, "same"},
		{"Common prefix 'ab'", []string{"abc", "abd", "ab"}, "ab"},
		{"Empty list", []string{}, ""},
		{"One empty string", []string{""}, ""},
		{"Mixed case", []string{"Test", "testing", "tester"}, "Test"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findLongestPrefix(tt.input)
			if result != tt.expected {
				t.Errorf("findLongestPrefix(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}
