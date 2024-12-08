package main

import (
	"strconv"
	"testing"
)

var testInput = `
190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20
`

func TestSolve_Part1(t *testing.T) {
	testcases := []struct {
		input string
		want  int64
	}{
		{testInput, 3749},
	}

	for i, tt := range testcases {
		t.Run("#"+strconv.Itoa(i), func(t *testing.T) {
			eqs := parseInput(tt.input)
			if res := solve_part1(eqs); res != tt.want {
				t.Fatalf("unexpected result:\nwant: %d\nhave: %d", tt.want, res)
			}
		})
	}
}

func TestSolve_Part2(t *testing.T) {
	testcases := []struct {
		input string
		want  int64
	}{
		{"156: 15 6", 156},
		{"7290: 6 8 6 15", 7290},
		{testInput, 11387},
	}

	for i, tt := range testcases {
		t.Run("#"+strconv.Itoa(i), func(t *testing.T) {
			eqs := parseInput(tt.input)
			if res := solve_part2(eqs); res != tt.want {
				t.Fatalf("unexpected result:\nwant: %d\nhave: %d", tt.want, res)
			}
		})
	}
}
