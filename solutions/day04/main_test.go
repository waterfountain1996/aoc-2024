package main

import (
	"runtime/debug"
	"strconv"
	"testing"
)

func TestSolve_Part1(t *testing.T) {
	testcases := []struct {
		input string
		want  int
	}{
		{

			input: `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`,
			want: 18,
		},
		{
			input: `S  S  S
 A A A
  MMM
SAMXMAS
  MMM
 A A A
S  S  S`,
			want: 8,
		},
	}

	for i, tt := range testcases {
		t.Run("#"+strconv.Itoa(i), func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Fatalf("panicked: %v\n%s", r, debug.Stack())
				}
			}()

			if res := solve_part1(normalizeInput(tt.input)); res != tt.want {
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

			input: `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`,
			want: 9,
		},
	}

	for i, tt := range testcases {
		t.Run("#"+strconv.Itoa(i), func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Fatalf("panicked: %v\n%s", r, debug.Stack())
				}
			}()

			if res := solve_part2(normalizeInput(tt.input)); res != tt.want {
				t.Fatalf("unexpected result:\nwant: %d\nhave: %d", tt.want, res)
			}
		})
	}
}
