package main

import (
	"strconv"
	"testing"
)

var testInput = `2333133121414131402`

func TestSolve_Part1(t *testing.T) {
	testcases := []struct {
		input string
		want  int
	}{
		{testInput, 1928},
	}

	for i, tt := range testcases {
		t.Run("#"+strconv.Itoa(i), func(t *testing.T) {
			disk := parseInput(tt.input)
			if have := solve_part1(disk); have != tt.want {
				t.Fatalf("unexpected result:\nwant: %d\nhave: %d", tt.want, have)
			}
		})
	}
}

func TestSolve_Part2(t *testing.T) {
	testcases := []struct {
		input string
		want  int64
	}{
		{testInput, 2858},
	}

	for i, tt := range testcases {
		t.Run("#"+strconv.Itoa(i), func(t *testing.T) {
			disk := parseInput(tt.input)
			if have := solve_part2(disk); have != tt.want {
				t.Fatalf("unexpected result:\nwant: %d\nhave: %d", tt.want, have)
			}
		})
	}
}
