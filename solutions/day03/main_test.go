package main

import "testing"

func TestSolve(t *testing.T) {
	testcases := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "part1",
			input: `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`,
			want:  161,
		},
		{
			name:  "part2",
			input: `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`,
			want:  48,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Fatalf("panicked: %v", r)
				}
			}()

			if res := doSolve(tt.input, tt.name == "part1"); res != tt.want {
				t.Fatalf("unexpected result:\nwant: %d\nhave: %d", tt.want, res)
			}
		})
	}
}
