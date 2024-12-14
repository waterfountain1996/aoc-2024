package main

import (
	"strconv"
	"testing"
)

func TestSolve_Part1(t *testing.T) {
	testcases := []struct {
		input         string
		seconds       int
		width, height int
		want          int
	}{
		{
			input: `
p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3
			`,
			seconds: 100,
			width:   11,
			height:  7,
			want:    12,
		},
	}

	for i, tt := range testcases {
		t.Run("#"+strconv.Itoa(i), func(t *testing.T) {
			rules := parseInput(tt.input)
			if have := solve_part1(rules, tt.seconds, tt.width, tt.height); have != tt.want {
				t.Fatalf("unexpected result:\nwant: %d\nhave: %d", tt.want, have)
			}
		})
	}
}
