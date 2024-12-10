package main

import (
	"bytes"
	"strconv"
	"testing"
)

func TestSolve_Part1(t *testing.T) {
	testcases := []struct {
		input string
		want  int
	}{
		{
			input: `
..90..9
...1.98
...2..7
6543456
765.987
876....
987....
			`,
			want: 4,
		},
		{
			input: `
10..9..
2...8..
3...7..
4567654
...8..3
...9..2
.....01
			`,
			want: 3,
		},
		{
			input: `
89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732
			`,
			want: 36,
		},
	}

	for i, tt := range testcases {
		t.Run("#"+strconv.Itoa(i), func(t *testing.T) {
			lines := bytes.Split(bytes.TrimSpace([]byte(tt.input)), []byte{'\n'})
			if have := solve_part1(lines); have != tt.want {
				t.Fatalf("unexpected result:\nwant: %d\nhave: %d", tt.want, have)
			}
		})
	}
}

func TestSolve_Part2(t *testing.T) {
	testcases := []struct {
		input string
		want  int
	}{
		{
			input: `
.....0.
..4321.
..5..2.
..6543.
..7..4.
..8765.
..9....
			`,
			want: 3,
		},
		{
			input: `
..90..9
...1.98
...2..7
6543456
765.987
876....
987....
			`,
			want: 13,
		},
		{
			input: `
012345
123456
234567
345678
4.6789
56789.
			`,
			want: 227,
		},
		{
			input: `
89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732
			`,
			want: 81,
		},
	}

	for i, tt := range testcases {
		t.Run("#"+strconv.Itoa(i), func(t *testing.T) {
			lines := bytes.Split(bytes.TrimSpace([]byte(tt.input)), []byte{'\n'})
			if have := solve_part2(lines); have != tt.want {
				t.Fatalf("unexpected result:\nwant: %d\nhave: %d", tt.want, have)
			}
		})
	}
}
