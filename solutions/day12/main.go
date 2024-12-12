package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
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
	return strconv.Itoa(p1), ""
}

type point [2]int

func (p point) String() string {
	return fmt.Sprintf("(%d; %d)", p[0], p[1])
}

type direction int

const (
	dirUp direction = iota
	dirRight
	dirDown
	dirLeft
	dirCount
)

func (d direction) Next() direction {
	return (d + 1) % dirCount
}

func (d direction) Prev() direction {
	if d == 0 {
		return dirCount - 1
	}
	return d - 1
}

type region map[point]struct{}

func (r region) Area() int {
	return len(r)
}

func (r region) Perimeter() int {
	var total int
	for p := range r {
		y, x := p[0], p[1]
		pts := [4]point{
			{y - 1, x}, // Up
			{y, x + 1}, // Right
			{y + 1, x}, // Down
			{y, x - 1}, // Left
		}

		v := 4
		for _, n := range pts {
			if _, ok := r[n]; ok {
				v--
			}
		}
		total += v
	}
	return total
}

type segment struct {
	Start  int
	Length int
}

func (s segment) End() int {
	return s.Start + s.Length - 1
}

func (s segment) String() string {
	return fmt.Sprintf("[%d; %d)", s.Start, s.Start+s.Length)
}

func (r region) Sides() int {
	var (
		ymin, ymax int
		xmin, xmax int
	)
	for p := range r {
		y, x := p[0], p[1]

		ymin, ymax = min(y, ymin), max(y, ymax)
		xmin, xmax = min(x, xmin), max(x, xmax)
	}

	levels := make([][]segment, 0, ymax-ymin)
	for y := ymin; y <= ymax; y++ {
		var lvl []segment

		for x := xmin; x <= xmax; x++ {
			start := x
			for {
				p := point{y, x}
				if _, ok := r[p]; !ok {
					if x-start > 0 {
						lvl = append(lvl, segment{start, x - start})
					}
					break
				}
				x++
			}
		}
		if len(lvl) > 0 {
			levels = append(levels, lvl)
		}
	}
	return 0
}

// MergeSegments combines and sorts the points from two levels.

func FindNonOverlappingSegments(level1, level2 []segment) []segment {
	type Point struct {
		Position int
		Change   int
	}
	points := []Point{}

	// Convert segments to points
	for _, seg := range level1 {
		points = append(points, Point{Position: seg.Start, Change: 1})
		points = append(points, Point{Position: seg.Start + seg.Length, Change: -1})
	}
	for _, seg := range level2 {
		points = append(points, Point{Position: seg.Start, Change: 2})
		points = append(points, Point{Position: seg.Start + seg.Length, Change: -2})
	}

	// Sort points by position, with ties broken by Change value
	sort.Slice(points, func(i, j int) bool {
		if points[i].Position == points[j].Position {
			return points[i].Change < points[j].Change
		}
		return points[i].Position < points[j].Position
	})

	var result []segment
	currentLevel1 := 0
	currentLevel2 := 0
	lastPosition := -1

	for _, point := range points {
		// If we have moved to a new position, evaluate the interval
		if lastPosition != -1 && point.Position > lastPosition {
			if (currentLevel1 > 0 && currentLevel2 == 0) || (currentLevel2 > 0 && currentLevel1 == 0) {
				result = append(result, segment{
					Start:  lastPosition,
					Length: point.Position - lastPosition,
				})
			}
		}

		// Update level counts based on the change
		switch point.Change {
		case 1:
			currentLevel1++
		case -1:
			currentLevel1--
		case 2:
			currentLevel2++
		case -2:
			currentLevel2--
		}

		lastPosition = point.Position
	}

	return result
}

func abs(i int) int {
	if i < 0 {
		i = -i
	}
	return i
}

func solve_part1(lines [][]byte) int {
	return calculateTotal(lines, func(r region) int {
		return r.Area() * r.Perimeter()
	})
}

func solve_part2(lines [][]byte) int {
	return calculateTotal(lines, func(r region) int {
		return r.Area() * r.Sides()
	})
}

func calculateTotal(lines [][]byte, f func(region) int) int {
	var total int
	seen := make(map[point]bool)
	for y, line := range lines {
		for x, plant := range line {
			if plant != 'I' {
				// continue
			}

			p := point{y, x}
			if seen[p] {
				continue
			}

			reg := bfs(lines, y, x)
			for k := range reg {
				seen[k] = true
			}
			total += f(reg)
		}
	}
	return total
}

func bfs(lines [][]byte, y, x int) region {
	isOOB := func(p point) bool {
		y, x := p[0], p[1]
		return y < 0 || y >= len(lines) || x < 0 || x >= len(lines[y])
	}

	plant := lines[y][x]
	queue := []point{{y, x}}
	reg := make(map[point]struct{})

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		if _, ok := reg[p]; ok {
			continue
		}

		y, x := p[0], p[1]
		pts := [4]point{
			{y - 1, x}, // Up
			{y, x + 1}, // Right
			{y + 1, x}, // Down
			{y, x - 1}, // Left
		}
		for _, next := range pts {
			if isOOB(next) {
				continue
			} else if _, ok := reg[next]; ok {
				continue
			} else if lines[next[0]][next[1]] != plant {
				continue
			}
			queue = append(queue, next)
		}
		reg[p] = struct{}{}
	}
	return reg
}
