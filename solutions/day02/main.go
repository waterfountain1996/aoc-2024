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
	var (
		reports = parseInput(input)
		p1      = solve_part1(reports)
		p2      = solve_part2(reports)
	)
	return strconv.Itoa(p1), strconv.Itoa(p2)
}

func parseInput(input string) [][]int {
	lines := strings.Split(input, "\n")

	reports := make([][]int, len(lines))
	for i, line := range lines {
		levels := strings.Split(line, " ")
		reports[i] = make([]int, len(levels))

		for j, level := range levels {
			n, err := strconv.Atoi(level)
			if err != nil {
				panic(fmt.Sprintf("invalid input on line %d", i))
			}
			reports[i][j] = n
		}
	}
	return reports
}

func solve_part1(reports [][]int) int {
	var total int
	for _, report := range reports {
		if ok := isSafe(report); ok {
			total++
		}
	}

	return total
}

func isSafe(report []int) bool {
	var (
		safe       = true
		decreasing bool
	)
	for i, v := range report[1:] {
		prev := report[i]

		if i == 0 {
			decreasing = v < prev
		}

		if (decreasing && v > prev) || (!decreasing && v < prev) {
			safe = false
			break
		}

		diff := v - prev
		if decreasing {
			diff = -diff
		}

		if !(1 <= diff && diff <= 3) {
			safe = false
			break
		}
	}
	return safe
}

func solve_part2(reports [][]int) int {
	var total int
	for _, report := range reports {
		var safe bool
		for i := range report {
			r := slices.Concat(report[:i], report[i+1:])
			if ok := isSafe(r); ok {
				safe = true
				break
			}
		}

		if safe {
			total++
		}
	}

	return total
}
