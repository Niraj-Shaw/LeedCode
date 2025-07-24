package main

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {

	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"3 length of substring", "pwwkew", 3},
		{"1 length of substring", "bbbbbb", 1},
		{"0 length of substring", "", 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := lengthOfLongestSubstring(test.input)
			if result != test.want {
				t.Errorf("lengthOfLongestSubstring(%v) = %v, want %d", test.input, result, test.want)
			}
		})
	}
}
