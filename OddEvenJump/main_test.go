package main

import "testing"

func TestMain(t *testing.T) {
	tests := []struct {
		name   string
		input  [][]int
		output int
	}{
		{
			name:   "test with output 2",
			input:  [][]int{{10, 13, 12, 14, 15}},
			output: 2,
		},
		{
			name:   "test with output 3",
			input:  [][]int{{2, 3, 1, 1, 4}, {5, 1, 3, 4, 2}},
			output: 3,
		},
		{
			name:   "test with output 6",
			input:  [][]int{{1, 2, 3, 2, 1, 4, 4, 5}},
			output: 6,
		},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, ti := range tt.input {
				print(i)
				result := oddEvenJumps(ti)

				if result != tt.output {
					t.Fatalf("wanted %d, found %d", tt.output, result)
				}
			}
		})
	}
}
