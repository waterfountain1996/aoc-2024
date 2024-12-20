package main

import (
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	b, err := os.ReadFile("inputs/day20.txt")
	if err != nil {
		panic(err)
	}
	p1, p2 := solve(string(b))
	fmt.Printf("%s\n%s\n", p1, p2)
}

func solve(input string) (string, string) {
	grid := parseInput(input)
	p1 := solve_part1(grid)
	p2 := solve_part2(grid)
	return strconv.Itoa(p1), strconv.Itoa(p2)
}

func parseInput(input string) []string {
	input = strings.TrimSpace(input)
	return strings.Split(input, "\n")
}

func solve_part1(grid []string) int {
	return solve2(grid, 2)
}

func solve_part2(grid []string) int {
	return solve2(grid, 20)
}

func solve2(grid []string, picos int) int {
	var start, end point
	for y, row := range grid {
		for x, col := range row {
			p := point{y, x}
			switch col {
			case 'S':
				start = p
			case 'E':
				end = p
			}
		}
	}

	dist, prev := dijkstra(grid, start, end)

	unique := make(map[cheat]struct{})

	for ce := end; ce != start; ce = prev[ce] {
		var (
			visited = make(map[point]bool)
			queue   = []state{{point: ce}}
		)
		for len(queue) > 0 {
			it := queue[0]
			queue = queue[1:]

			if 2 <= it.dist && it.dist <= picos {
				cs := it.point

				if v, ok := dist[cs]; ok {
					if diff := dist[ce] - v - it.dist; diff >= 100 {
						unique[cheat{cs, ce}] = struct{}{}
					}
				}
			} else if it.dist > picos {
				continue
			}

			if visited[it.point] {
				continue
			}
			visited[it.point] = true

			for _, to := range []point{
				{it.point.Y - 1, it.point.X},
				{it.point.Y, it.point.X + 1},
				{it.point.Y + 1, it.point.X},
				{it.point.Y, it.point.X - 1},
			} {
				queue = append(queue, state{
					point: to,
					dist:  it.dist + 1,
				})
			}
		}
	}
	return len(unique)
}

type point struct {
	Y, X int
}

func (p point) Manhattan(other point) int {
	dy := p.Y - other.Y
	if dy < 0 {
		dy = -dy
	}

	dx := p.X - other.X
	if dx < 0 {
		dx = -dx
	}
	return dy + dx
}

type cheat struct {
	Start, End point
}

type state struct {
	point point
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

func dijkstra(grid []string, start, end point) (map[point]int, map[point]point) {
	var hq heapQueue
	heap.Push(&hq, &state{point: start, dist: 0})

	var (
		visited = make(map[point]bool)
		dist    = make(map[point]int)
		prev    = make(map[point]point)
	)

	for hq.Len() > 0 {
		item := heap.Pop(&hq).(*state)

		if item.point == end {
			break
		}

		curr := item.point
		if visited[curr] {
			continue
		}
		visited[curr] = true

		y, x := item.point.Y, item.point.X
		pts := [4]point{
			{y - 1, x}, // Up
			{y, x + 1}, // Right
			{y + 1, x}, // Down
			{y, x - 1}, // Left
		}

		for _, to := range pts {
			if to.Y < 0 || to.Y >= len(grid) || to.X < 0 || to.X >= len(grid[to.Y]) {
				continue
			}

			if grid[to.Y][to.X] == '#' {
				continue
			}

			alt := dist[curr] + 1
			if v, ok := dist[to]; !ok || alt < v {
				dist[to] = alt
				prev[to] = item.point
				heap.Push(&hq, &state{
					point: to,
					dist:  alt,
				})
			}
		}
	}
	return dist, prev
}
