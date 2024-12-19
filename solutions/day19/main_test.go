package main

import (
	"strconv"
	"testing"
)

func TestSolve_Part1(t *testing.T) {
	cases := []struct {
		input string
		want  int
	}{
		{
			input: `
r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb
`,
			want: 6,
		},
	}

	for i, tt := range cases {
		t.Run("#"+strconv.Itoa(i), func(t *testing.T) {
			patterns, designs := parseInput(tt.input)
			if have := solve_part1(patterns, designs); have != tt.want {
				t.Fatalf("unexpected result:\nwant: %d\nhave: %d", tt.want, have)
			}
		})
	}
}

func TestSolve_Part2(t *testing.T) {
	cases := []struct {
		input string
		want  int
	}{
		{
			input: `
r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb
`,
			want: 16,
		},
	}

	for i, tt := range cases {
		t.Run("#"+strconv.Itoa(i), func(t *testing.T) {
			patterns, designs := parseInput(tt.input)
			if have := solve_part2(patterns, designs); have != tt.want {
				t.Fatalf("unexpected result:\nwant: %d\nhave: %d", tt.want, have)
			}
		})
	}
}
