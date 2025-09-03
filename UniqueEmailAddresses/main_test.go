package main

import "testing"

func TestMain(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
		output int
	}{
		{
			name: "test with 2 distinct emails",
			input: []string{"test.email+alex@leetcode.com",
				"test.e.mail+bob.cathy@leetcode.com",
				"testemail+david@lee.tcode.com"},
			output: 2,
		},
		{
			name: "test with all distinct emails",
			input: []string{"a@leetcode.com",
				"b@leetcode.com",
				"c@leetcode.com"},
			output: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := numUniqueEmails(test.input)
			if result != test.output {
				t.Fatalf("wanted %d, found %d", test.output, result)
			}
		})
	}
}
