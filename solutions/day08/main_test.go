package main

import (
	"bytes"
	"strconv"
	"testing"
)

var testInput = `
............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............
`

var testInputT = `
T.........
...T......
.T........
..........
..........
..........
..........
..........
..........
..........
`

func TestSolve_Part1(t *testing.T) {
	testcases := []struct {
		input string
		want  int
	}{
		{testInput, 14},
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
		{testInputT, 9},
		{testInput, 34},
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
