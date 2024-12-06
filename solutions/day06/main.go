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
	lines := parseInput(input)
	p1 := solve_part1(lines)
	p2 := solve_part2(lines)
	return strconv.Itoa(p1), strconv.Itoa(p2)
}

type matrix [][]byte

func (m matrix) Walk(y, x int, dir direction, visit func(y, x int, d direction) bool) {
	for {
		var dy, dx int
		switch dir {
		case dirUp:
			dy = -1
		case dirRight:
			dx = 1
		case dirDown:
			dy = 1
		case dirLeft:
			dx = -1
		}

		for {
			if ok := visit(y, x, dir); !ok {
				return
			}

			ny, nx := y+dy, x+dx
			if ny < 0 || ny >= len(m) || nx < 0 || nx >= len(m[ny]) {
				return
			}

			if m[ny][nx] == '#' {
				break
			}

			y, x = ny, nx
		}

		dir = dir.Next()
	}
}

func (m matrix) Clone() matrix {
	m2 := make(matrix, len(m))
	for row, b := range m {
		m2[row] = bytes.Clone(b)
	}
	return m2
}

type direction int

const (
	dirUp = direction(iota)
	dirRight
	dirDown
	dirLeft
	dirMax
)

func (d direction) Next() direction {
	return (d + 1) % dirMax
}

func (d direction) String() string {
	switch d {
	case dirUp:
		return "^"
	case dirRight:
		return ">"
	case dirDown:
		return "v"
	case dirLeft:
		return "<"
	default:
		return ""
	}
}

func dirFromByte(b byte) direction {
	switch b {
	case '^':
		return dirUp
	case '>':
		return dirRight
	case 'v':
		return dirDown
	case '<':
		return dirLeft
	default:
		panic(fmt.Sprintf("invalid direction: %c", b))
	}
}

type point [2]int

type edge struct {
	p point
	d direction
}

func (e edge) String() string {
	return fmt.Sprintf("%d-%d-%s", e.p[0], e.p[1], e.d)
}

func parseInput(input []byte) matrix {
	lines := bytes.Split(bytes.TrimSpace(input), []byte{'\n'})
	return lines
}

func solve_part1(lines matrix) int {
	var (
		sy, sx int
		dir    direction
	)
	for row, line := range lines {
		col := bytes.IndexAny(line, "^>v<")
		if col == -1 {
			continue
		}

		sy, sx = row, col
		dir = dirFromByte(line[col])
	}

	seen := make(map[point]struct{})

	lines.Walk(sy, sx, dir, func(y, x int, _ direction) bool {
		seen[point{y, x}] = struct{}{}
		return true
	})

	return len(seen)
}

func solve_part2(lines matrix) int {
	var (
		sy, sx int
		sd     direction
	)
	for row, line := range lines {
		col := bytes.IndexAny(line, "^>v<")
		if col == -1 {
			continue
		}

		sy, sx = row, col
		sd = dirFromByte(line[col])
	}

	m := lines.Clone()

	obstacles := make(map[point]struct{})
	lines.Walk(sy, sx, sd, func(y, x int, d direction) bool {
		seen := make(map[edge]struct{})

		oy, ox := y, x
		switch d {
		case dirUp:
			oy--
		case dirRight:
			ox++
		case dirDown:
			oy++
		case dirLeft:
			ox--
		}
		if oy < 0 || oy >= len(m) || ox < 0 || ox >= len(m[oy]) {
			// This is the exit, continue.
			return true
		}

		ch := m[oy][ox]
		if ch == '#' {
			// Already an obstacle, continue.
			return true
		}
		m[oy][ox] = '#'

		m.Walk(sy, sx, sd, func(ny, nx int, nd direction) bool {
			e := edge{point{ny, nx}, nd}
			if _, ok := seen[e]; ok {
				obstacles[point{oy, ox}] = struct{}{}
				return false
			}

			seen[e] = struct{}{}
			return true
		})

		// Remove added obstacle before the next iteration.
		m[oy][ox] = ch
		return true
	})

	return len(obstacles)
}
