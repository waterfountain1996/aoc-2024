package main

import (
	"strconv"
	"testing"
)

func TestSolve_Part1(t *testing.T) {
	cases := []struct {
		input      string
		start, end point
		n          int
		want       int
	}{
		{
			input: `
5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0
`,
			start: point{0, 0},
			end:   point{6, 6},
			n:     12,
			want:  22,
		},
	}

	for i, tt := range cases {
		t.Run("#"+strconv.Itoa(i), func(t *testing.T) {
			bytes := parseInput(tt.input)
			if have := solve_part1(bytes, tt.start, tt.end, tt.n); have != tt.want {
				t.Fatalf("unexpected result:\nwant: %d\nhave: %d", tt.want, have)
			}
		})
	}
}

func TestSolve_Part2(t *testing.T) {
	cases := []struct {
		input      string
		start, end point
		want       string
	}{
		{
			input: `
5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0
`,
			start: point{0, 0},
			end:   point{6, 6},
			want:  "6,1",
		},
	}

	for i, tt := range cases {
		t.Run("#"+strconv.Itoa(i), func(t *testing.T) {
			bytes := parseInput(tt.input)
			if have := solve_part2(bytes, tt.start, tt.end); have != tt.want {
				t.Fatalf("unexpected result:\nwant: %s\nhave: %s", tt.want, have)
			}
		})
	}
}
