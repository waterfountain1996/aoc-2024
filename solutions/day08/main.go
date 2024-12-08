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

type point [2]int

func (p point) String() string {
	return fmt.Sprintf("(%d; %d)", p[0], p[1])
}

func solve_part1(lines [][]byte) int {
	antennas := make(map[byte][]point)
	for row, line := range lines {
		for col, node := range line {
			if node == '.' {
				continue
			}

			antennas[node] = append(antennas[node], point{row, col})
		}
	}

	isOOB := func(p point) bool {
		return p[0] < 0 || p[0] >= len(lines) || p[1] < 0 || p[1] >= len(lines[p[0]])
	}

	antinodes := make(map[point]struct{})
	for _, pts := range antennas {
		for i := 0; i < len(pts); i++ {
			for j := i + 1; j < len(pts); j++ {
				a, b := pts[i], pts[j]

				for _, reverse := range [...]bool{false, true} {
					walkLine(a, b, isOOB, reverse, func(p point) bool {
						antinodes[p] = struct{}{}
						return false
					})
				}
			}
		}
	}
	return len(antinodes)
}

func solve_part2(lines [][]byte) int {
	antennas := make(map[byte][]point)
	for row, line := range lines {
		for col, node := range line {
			if node == '.' {
				continue
			}
			antennas[node] = append(antennas[node], point{row, col})
		}
	}

	isOOB := func(p point) bool {
		return p[0] < 0 || p[0] >= len(lines) || p[1] < 0 || p[1] >= len(lines[p[0]])
	}

	antinodes := make(map[point]struct{})

	for _, pts := range antennas {
		for i := 0; i < len(pts); i++ {
			for j := i + 1; j < len(pts); j++ {
				a, b := pts[i], pts[j]
				for _, reverse := range [...]bool{false, true} {
					walkLine(a, b, isOOB, reverse, func(p point) bool {
						antinodes[p] = struct{}{}
						return true
					})
				}
			}
		}
	}

	return len(antinodes)
}

func walkLine(a, b point, isOOB func(point) bool, reverse bool, f func(point) bool) {
	node := a
	dy, dx := b[0]-a[0], b[1]-a[1]
	if reverse {
		node = b
		dy, dx = a[0]-b[0], a[1]-b[1]
	}

	for {
		next := point{node[0] + dy, node[1] + dx}
		if isOOB(next) {
			return
		}

		if ok := f(next); !ok {
			return
		}

		node = next
	}
}
