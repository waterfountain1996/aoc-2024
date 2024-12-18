package main

import (
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	b, err := os.ReadFile("inputs/day18.txt")
	if err != nil {
		panic(err)
	}
	p1, p2 := solve(string(b))
	fmt.Printf("%s\n%s\n", p1, p2)
}

func solve(input string) (string, string) {
	bytes := parseInput(input)
	p1 := solve_part1(bytes, point{0, 0}, point{70, 70}, 1024)
	p2 := solve_part2(bytes, point{0, 0}, point{70, 70})
	return strconv.Itoa(p1), p2
}

type point struct {
	X int
	Y int
}

func parseInput(input string) []point {
	input = strings.TrimSpace(input)
	var bytes []point
	for _, line := range strings.Split(input, "\n") {
		xs, ys, _ := strings.Cut(line, ",")
		var p point
		p.X, _ = strconv.Atoi(xs)
		p.Y, _ = strconv.Atoi(ys)
		bytes = append(bytes, p)
	}
	return bytes
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

func solve_part1(bytes []point, start, end point, n int) int {
	dist := dijkstra(bytes, start, end, n)
	return dist[end]
}

func solve_part2(bytes []point, start, end point) string {
	for i := range bytes {
		dist := dijkstra(bytes, start, end, i+1)
		if _, ok := dist[end]; !ok {
			return fmt.Sprintf("%d,%d", bytes[i].X, bytes[i].Y)
		}
	}
	return "-1,-1"
}

func simulate(bytes []point, n int) map[point]struct{} {
	corrupted := make(map[point]struct{})
	for i := range n {
		corrupted[bytes[i]] = struct{}{}
	}
	return corrupted
}

func dijkstra(bytes []point, start, end point, n int) map[point]int {
	corrupted := simulate(bytes, n)

	var hq heapQueue
	heap.Push(&hq, &state{point: start, dist: 0})

	var (
		visited = make(map[point]bool)
		dist    = make(map[point]int)
	)

	for hq.Len() > 0 {
		item := heap.Pop(&hq).(*state)

		curr := item.point
		if visited[curr] {
			continue
		}
		visited[curr] = true

		y, x := item.point.Y, item.point.X
		pts := [4]point{
			{x, y - 1}, // Up
			{x + 1, y}, // Right
			{x, y + 1}, // Down
			{x - 1, y}, // Left
		}

		for _, to := range pts {
			if to.X < 0 || to.X > end.X || to.Y < 0 || to.Y > end.Y {
				continue
			}

			if _, ok := corrupted[to]; ok {
				continue
			}

			alt := dist[curr] + 1
			if v, ok := dist[to]; !ok || alt < v {
				dist[to] = alt
				heap.Push(&hq, &state{
					point: to,
					dist:  alt,
				})
			}
		}
	}
	return dist
}
