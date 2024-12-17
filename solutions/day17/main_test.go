package main

import (
	"strconv"
	"testing"
)

func TestSolve_Part1(t *testing.T) {
	testcases := []struct {
		input string
		want  string
	}{
		{
			input: `
Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0
`,
			want: "4,6,3,5,6,3,5,2,1,0",
		},
	}

	for i, tt := range testcases {
		t.Run("#"+strconv.Itoa(i), func(t *testing.T) {
			comp := parseInput(tt.input)
			if have := solve_part1(comp); have != tt.want {
				t.Fatalf("unexpected result:\nwant: %s\nhave: %s", tt.want, have)
			}
		})
	}
}

func TestSolve_Part2(t *testing.T) {
	testcases := []struct {
		input string
		want  uint64
	}{
		{
			input: `
Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0
`,
			want: 117440,
		},
	}

	for i, tt := range testcases {
		t.Run("#"+strconv.Itoa(i), func(t *testing.T) {
			comp := parseInput(tt.input)
			if have := solve_part2(comp); have != tt.want {
				t.Fatalf("unexpected result:\nwant: %d\nhave: %d", tt.want, have)
			}
		})
	}
}
