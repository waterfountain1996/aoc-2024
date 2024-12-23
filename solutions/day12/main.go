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
	p2 := solve_part2(lines)
	return strconv.Itoa(p1), strconv.Itoa(p2)
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

type level []segment

func (r region) levels() (horizontal, vertical []level) {
	var (
		ymin, ymax int
		xmin, xmax int
	)
	for p := range r {
		y, x := p[0], p[1]
		ymin, ymax = min(y, ymin), max(y, ymax)
		xmin, xmax = min(x, xmin), max(x, xmax)
	}

	var lh []level
	for y := ymin; y <= ymax; y++ {
		var lvl level

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
			lh = append(lh, lvl)
		}
	}

	var lv []level
	for x := xmax; x >= xmin; x-- {
		var lvl level
		for y := ymin; y <= ymax; y++ {
			start := y
			for {
				p := point{y, x}
				if _, ok := r[p]; !ok {
					if y-start > 0 {
						lvl = append(lvl, segment{start, y - start})
					}
					break
				}
				y++
			}
		}
		if len(lvl) > 0 {
			lv = append(lv, lvl)
		}
	}

	return lh, lv
}

func (r region) Sides() int {
	var total int

	lh, lv := r.levels()
	for _, levels := range [][]level{lh, lv} {
		total += len(levels[0])

		for i := 1; i < len(levels); i++ {
			a, b := levels[i-1], levels[i]
			total += len(nonOverlapping(a, b))
		}

		n := len(levels)
		total += len(levels[n-1])
	}

	return total
}

func nonOverlapping(a, b []segment) []segment {
	var pts [][2]int
	for _, s := range a {
		pts = append(pts, [2]int{s.Start, 1}, [2]int{s.End() + 1, -1})
	}
	for _, s := range b {
		pts = append(pts, [2]int{s.Start, -1}, [2]int{s.End() + 1, 1})
	}
	sort.SliceStable(pts, func(i, j int) bool {
		return pts[i][0] < pts[j][0]
	})

	var state, start int
	var result []segment
	for _, p := range pts {
		x, d := p[0], p[1]
		state += d
		if state != 0 {
			start = x
		} else {
			if x != start {
				result = append(result, segment{
					Start:  start,
					Length: x - start,
				})
			}
		}
	}
	return result
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
			if plant != 'F' {
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
