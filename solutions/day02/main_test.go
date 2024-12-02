package main

import (
	"strings"
	"testing"
)

var testInput = strings.TrimSpace(`
7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
`)

func TestSolve_Part1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("panicked: %v", r)
		}
	}()

	if res := solve_part1(parseInput(testInput)); res != 2 {
		t.Fatalf("unexpected result:\nwant: %d\nhave: %d", 2, res)
	}
}

func TestSolve_Part2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("panicked: %v", r)
		}
	}()

	if res := solve_part2(parseInput(testInput)); res != 4 {
		t.Fatalf("unexpected result:\nwant: %d\nhave: %d", 4, res)
	}
}
