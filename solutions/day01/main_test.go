package main

import (
	"strings"
	"testing"
)

var testInput = strings.TrimSpace(`
3   4
4   3
2   5
1   3
3   9
3   3
`)

func TestSolve_Part1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("panicked: %v", r)
		}
	}()

	if res := solve_part1(parseInput(testInput)); res != 11 {
		t.Fatalf("unexpected result:\nwant: %v\nhave: %v", 11, res)
	}
}

func TestSolve_Part2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("panicked: %v", r)
		}
	}()

	if res := solve_part2(parseInput(testInput)); res != 31 {
		t.Fatalf("unexpected result:\nwant: %v\nhave: %v", 31, res)
	}
}
