package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {

	tests := []struct {
		name  string
		input []string
		want  []string
	}{
		{
			name:  "Mixed letter and digit logs",
			input: []string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero"},
			want:  []string{"let1 art can", "let3 art zero", "let2 own kit dig", "dig1 8 1 5 1", "dig2 3 6"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := reorderLogFiles(tc.input)
			assert.Equal(t, tc.want, got, "The output order is incorrect")
		})
	}

}
