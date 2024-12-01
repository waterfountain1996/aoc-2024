package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
		os.Exit(1)
	}
	p1, p2 := solve2(strings.TrimSpace(string(input)))
	fmt.Printf("%s\n%s\n", p1, p2)
}

type pair [2]int

func parseInput(input string) []pair {
	lines := strings.Split(input, "\n")

	lists := make([]pair, 0, len(lines))
	for i, line := range lines {
		before, after, _ := strings.Cut(line, " ")

		left, err := strconv.Atoi(strings.TrimSpace(before))
		if err != nil {
			panic(fmt.Sprintf("invalid input on line %d", i+1))
		}

		right, err := strconv.Atoi(strings.TrimSpace(after))
		if err != nil {
			panic(fmt.Sprintf("invalid input on line %d", i+1))
		}

		lists = append(lists, pair{left, right})
	}
	return lists
}

func solve2(input string) (string, string) {
	var (
		lists = parseInput(input)
		p1    = solve_part1(lists)
		p2    = solve_part2(lists)
	)
	return strconv.Itoa(p1), strconv.Itoa(p2)
}

func solve_part1(lists []pair) int {
	left, right := make([]int, len(lists)), make([]int, len(lists))
	for i, p := range lists {
		left[i], right[i] = p[0], p[1]
	}

	slices.Sort(left)
	slices.Sort(right)

	var total int
	for i := range left {
		diff := left[i] - right[i]
		if diff < 0 {
			diff = -diff
		}
		total += diff
	}
	return total
}

func solve_part2(lists []pair) int {
	entries := make(map[int]int)
	for _, p := range lists {
		v := p[1]
		entries[v]++
	}

	var total int
	for _, p := range lists {
		v := p[0]
		total += v * entries[v]
	}
	return total
}
