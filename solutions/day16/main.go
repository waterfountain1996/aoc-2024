package main

import (
	"bytes"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

func main() {
	b, err := os.ReadFile("inputs/day16.txt")
	if err != nil {
		panic(err)
	}
	p1, p2 := solve(b)
	fmt.Printf("%s\n%s\n", p1, p2)
}

func solve(input []byte) (string, string) {
	grid := bytes.Split(bytes.TrimSpace(input), []byte{'\n'})
	p1 := solve_part1(grid)
	p2 := solve_part2(grid)
	return strconv.Itoa(p1), strconv.Itoa(p2)
}

type point struct {
	y, x int
}

func (p point) Advance(dir direction) point {
	return [4]point{
		{p.y, p.x + 1}, // East
		{p.y + 1, p.x}, // South
		{p.y, p.x - 1}, // West
		{p.y - 1, p.x}, // North
	}[dir]
}

type direction int

const (
	dirEast = direction(iota)
	dirSouth
	dirWest
	dirNorth
	dirCount
)

func (d direction) Prev() direction {
	if d == 0 {
		return dirCount - 1
	}
	return d - 1
}

func (d direction) Next() direction {
	return (d + 1) % dirCount
}

func (d direction) Reverse() direction {
	return (d + 2) % dirCount
}

type dirpoint struct {
	p point
	d direction
}

type state struct {
	point point
	dir   direction
	steps []dirpoint
	dist  int
	index int
}

type heapQueue []*state

func (hq heapQueue) Len() int { return len(hq) }

func (pq heapQueue) Less(i, j int) bool {
	return pq[i].dist < pq[j].dist
}

func (pq heapQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *heapQueue) Push(x interface{}) {
	item := x.(*state)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *heapQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[:n-1]
	return item
}

func solve_part1(grid [][]byte) int {
	var (
		start point
		end   point
		dir   direction
	)
	for y, row := range grid {
		for x, col := range row {
			if col == 'S' {
				start = point{y, x}
			} else if col == 'E' {
				end = point{y, x}
			}
		}
	}

	topscore, _ := solve2(grid, start, end, dir)
	return topscore
}

func solve_part2(grid [][]byte) int {
	var (
		start point
		end   point
		dir   direction
	)
	for y, row := range grid {
		for x, col := range row {
			if col == 'S' {
				start = point{y, x}
			} else if col == 'E' {
				end = point{y, x}
			}
		}
	}

	topscore, bestRoutes := solve2(grid, start, end, dir)

	seen := make(map[point]bool)
	for _, p := range bestRoutes[topscore] {
		seen[p.p] = true
	}
	return len(seen) + 1
}

func solve2(grid [][]byte, start, end point, dir direction) (int, map[int][]dirpoint) {
	var hq heapQueue
	heap.Init(&hq)
	heap.Push(&hq, &state{
		point: start,
		dir:   dir,
		dist:  0,
	})

	var (
		topscore   int = -1
		visited        = make(map[dirpoint]int)
		bestRoutes     = make(map[int][]dirpoint)
	)

	for hq.Len() > 0 {
		item := heap.Pop(&hq).(*state)

		if item.point == end {
			if topscore == -1 || item.dist <= topscore {
				route := append(bestRoutes[item.dist], item.steps...)
				bestRoutes[item.dist] = route
				topscore = item.dist
			}
			continue
		}

		for _, d := range []direction{item.dir, item.dir.Prev(), item.dir.Next()} {
			np := item.point.Advance(d)
			if grid[np.y][np.x] == '#' {
				continue
			}

			weight := item.dist + 1
			if d != item.dir {
				weight += 1000
			}

			dp := dirpoint{np, d}

			if v, ok := visited[dp]; ok {
				if v < weight {
					continue
				}
			}
			visited[dp] = weight

			steps := make([]dirpoint, len(item.steps))
			copy(steps, item.steps)

			heap.Push(&hq, &state{
				point: np,
				dir:   d,
				steps: append(steps, dp),
				dist:  weight,
			})
		}
	}
	return topscore, bestRoutes
}
