package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{name: "two parenthesis", input: "()", want: true},
		{name: "negative scenario", input: "(]", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValid(tt.input)

			assert.Equal(t, tt.want, got, "invalid input")
		})
	}
}
