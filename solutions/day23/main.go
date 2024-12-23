package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	b, err := os.ReadFile("inputs/day23.txt")
	if err != nil {
		panic(err)
	}
	p1, p2 := solve(string(b))
	fmt.Printf("%s\n%s\n", p1, p2)
}

func solve(input string) (string, string) {
	g := parseInput(input)
	p1 := solve_part1(g)
	p2 := solve_part2(g)
	return strconv.Itoa(p1), p2
}

func parseInput(input string) map[string][]string {
	g := make(map[string][]string)
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		a, b, _ := strings.Cut(line, "-")
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}
	return g
}

func solve_part1(g map[string][]string) int {
	sets := make(map[string]struct{})
	for a, values := range g {
		for _, b := range values {
			for _, c := range values {
				if c == b {
					continue
				}
				var (
					ok1 = slices.Contains(g[b], a) && slices.Contains(g[b], c)
					ok2 = slices.Contains(g[c], a) && slices.Contains(g[c], b)
				)
				if ok1 && ok2 {
					if a[0] == 't' || b[0] == 't' || c[0] == 't' {
						s := []string{a, b, c}
						slices.Sort(s)
						sets[strings.Join(s, ",")] = struct{}{}
					}
				}
			}
		}
	}
	return len(sets)
}

func solve_part2(g map[string][]string) string {
	var R, P, X []string
	for node := range g {
		P = append(P, node)
	}

	var cliques [][]string
	findCliques(g, R, P, X, &cliques)

	var password []string
	for _, c := range cliques {
		if len(c) > len(password) {
			password = c
		}
	}
	slices.Sort(password)
	return strings.Join(password, ",")
}

func findCliques(graph map[string][]string, R, P, X []string, cliques *[][]string) {
	if len(P) == 0 && len(X) == 0 {
		clique := make([]string, len(R))
		copy(clique, R)
		*cliques = append(*cliques, clique)
		return
	}

	for _, v := range P {
		newR := append(R, v)
		newP := intersect(P, graph[v])
		newX := intersect(X, graph[v])
		findCliques(graph, newR, newP, newX, cliques)

		P = slices.DeleteFunc(P, func(e string) bool {
			return e == v
		})
		X = append(X, v)
	}
}

func intersect(a, b []string) []string {
	set := make(map[string]bool)
	for _, v := range b {
		set[v] = true
	}
	var result []string
	for _, v := range a {
		if set[v] {
			result = append(result, v)
		}
	}
	return result
}
