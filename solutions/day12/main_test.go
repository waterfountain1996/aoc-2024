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
AAAA
BBCD
BBCC
EEEC
			`,
			want: 140,
		},
		{
			input: `
OOOOO
OXOXO
OOOOO
OXOXO
OOOOO
			`,
			want: 772,
		},
		{
			input: `
RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE
			`,
			want: 1930,
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
AAAA
BBCD
BBCC
EEEC
			`,
			want: 80,
		},
		{
			input: `
OOOOO
OXOXO
OOOOO
OXOXO
OOOOO
			`,
			want: 436,
		},
		{
			input: `
EEEEE
EXXXX
EEEEE
EXXXX
EEEEE
			`,
			want: 236,
		},
		{
			input: `
AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA
			`,
			want: 368,
		},
		{
			input: `
RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE
			`,
			want: 1206,
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
