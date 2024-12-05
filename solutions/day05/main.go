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
	p1, p2 := solve(strings.TrimSpace(string(input)))
	fmt.Printf("%s\n%s\n", p1, p2)
}

func solve(input string) (string, string) {
	rules, updates := parseInput(input)
	p1 := solve_part1(rules, updates)
	p2 := solve_part2(rules, updates)
	return strconv.Itoa(p1), strconv.Itoa(p2)
}

func parseInput(input string) (map[int][]int, [][]int) {
	rawRules, rawUpdates, _ := strings.Cut(strings.TrimSpace(input), "\n\n")

	rules := make(map[int][]int)
	for _, line := range strings.Split(rawRules, "\n") {
		var (
			a, b, _     = strings.Cut(line, "|")
			page, _     = strconv.Atoi(a)
			lessThan, _ = strconv.Atoi(b)
		)
		rules[lessThan] = append(rules[lessThan], page)
	}

	var updates [][]int
	for _, line := range strings.Split(rawUpdates, "\n") {
		elems := strings.Split(line, ",")
		up := make([]int, len(elems))
		for i, elem := range elems {
			up[i], _ = strconv.Atoi(elem)
		}
		updates = append(updates, up)
	}

	return rules, updates
}

func solve_part1(rules map[int][]int, updates [][]int) int {
	var res int
	for _, up := range updates {
		if ok := isCorrectUpdate(rules, up); ok {
			res += up[len(up)/2]
		}
	}
	return res
}

func solve_part2(rules map[int][]int, updates [][]int) int {
	var res int
	for _, up := range updates {
		if ok := isCorrectUpdate(rules, up); !ok {
			slices.SortFunc(up, func(a, b int) int {
				lessThan := rules[a]
				if slices.Contains(lessThan, b) {
					return -1
				}
				return 0
			})
			res += up[len(up)/2]
		}
	}
	return res
}

func isCorrectUpdate(rules map[int][]int, up []int) bool {
	for i := range up {
		sub := up[i+1:]
		if len(sub) == 0 {
			break
		}

		contains := slices.ContainsFunc(sub, func(v int) bool {
			return slices.Contains(rules[up[i]], v)
		})
		if contains {
			return false
		}
	}
	return true
}
