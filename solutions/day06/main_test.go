package main

import (
	"strconv"
	"testing"
)

var testInput = `
....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...
`

func TestSolve_Part1(t *testing.T) {
	testcases := []struct {
		input string
		want  int
	}{
		{
			input: testInput,
			want:  41,
		},
	}

	for i, tt := range testcases {
		t.Run("#"+strconv.Itoa(i), func(t *testing.T) {
			lines := parseInput([]byte(tt.input))
			if res := solve_part1(lines); res != tt.want {
				t.Fatalf("unexpected result:\nwant: %d\nhave: %d", tt.want, res)
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
			input: testInput,
			want:  6,
		},
	}

	for i, tt := range testcases {
		t.Run("#"+strconv.Itoa(i), func(t *testing.T) {
			lines := parseInput([]byte(tt.input))
			if res := solve_part2(lines); res != tt.want {
				t.Fatalf("unexpected result:\nwant: %d\nhave: %d", tt.want, res)
			}
		})
	}
}
