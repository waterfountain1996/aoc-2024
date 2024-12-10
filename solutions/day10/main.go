package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
		os.Exit(1)
	}
	p1, p2 := solve(input)
	fmt.Printf("%s\n%s\n", p1, p2)
}

func solve(input []byte) (string, string) {
	lines := bytes.Split(bytes.TrimSpace(input), []byte{'\n'})
	p1 := solve_part1(lines)
	p2 := solve_part2(lines)
	return strconv.Itoa(p1), strconv.Itoa(p2)
}

type matrix [][]byte

func (m matrix) IsOOB(p point) bool {
	y, x := p[0], p[1]
	return y < 0 || y >= len(m) || x < 0 || x >= len(m[y])
}

func (m matrix) TileAt(p point) byte {
	return m[p[0]][p[1]]
}

type point [2]int

func (p point) String() string {
	return fmt.Sprintf("(%d; %d)", p[0], p[1])
}

func solve_part1(m matrix) int {
	var score int
	for y, line := range m {
		for x, tile := range line {
			if tile == '0' {
				score += calculateScore(m, point{y, x}, nil)
			}
		}
	}
	return score
}

func calculateScore(m matrix, start point, cache map[point]struct{}) int {
	if cache == nil {
		cache = map[point]struct{}{}
	}

	if m.IsOOB(start) {
		return 0
	}

	if m.TileAt(start) == '9' {
		cache[start] = struct{}{}
		return 0
	}

	y, x := start[0], start[1]
	pts := [4]point{
		{y - 1, x}, // Up
		{y, x + 1}, // Right
		{y + 1, x}, // Down
		{y, x - 1}, // Left
	}
	for _, p := range pts {
		if m.IsOOB(p) {
			continue
		}

		if m.TileAt(p)-1 == m.TileAt(start) {
			calculateScore(m, p, cache)
		}
	}
	return len(cache)
}

func solve_part2(m matrix) int {
	var (
		cache = make(map[point]int)
		total int
	)
	for y, line := range m {
		for x, tile := range line {
			if tile == '0' {
				total += findUniquePaths(m, point{y, x}, cache)
			}
		}
	}
	return total
}

func findUniquePaths(m matrix, start point, cache map[point]int) int {
	if v, ok := cache[start]; ok {
		return v
	}
	cache[start] = 0

	y, x := start[0], start[1]
	pts := [4]point{
		{y - 1, x}, // Up
		{y, x + 1}, // Right
		{y + 1, x}, // Down
		{y, x - 1}, // Left
	}
	for _, p := range pts {
		if m.IsOOB(p) {
			continue
		}
		if tile := m.TileAt(p); tile-1 == m.TileAt(start) {
			if tile == '9' {
				cache[start]++
			} else {
				cache[start] += findUniquePaths(m, p, cache)
			}
		}
	}
	return cache[start]
}
