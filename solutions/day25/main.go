package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	b, err := os.ReadFile("inputs/day25.txt")
	if err != nil {
		panic(err)
	}
	p1, p2 := solve(string(b))
	fmt.Printf("%s\n%s\n", p1, p2)
}

func solve(input string) (string, string) {
	schematics := parseInput(input)
	p1 := solve_part1(schematics)
	return strconv.Itoa(p1), ""
}

func parseInput(input string) []string {
	input = strings.TrimSpace(input)
	return strings.Split(input, "\n\n")
}

type heights [5]int

func solve_part1(schematics []string) int {
	var locks, keys []heights
	for _, sch := range schematics {
		var (
			rows = strings.Split(sch, "\n")
			n    = len(rows)
			h    heights
		)
		for _, row := range rows[1 : n-1] {
			for j, col := range row {
				if col == '#' {
					h[j]++
				}
			}
		}
		if strings.IndexRune(rows[0], '#') == -1 {
			keys = append(keys, h)
		} else {
			locks = append(locks, h)
		}
	}

	var total int
	for _, key := range keys {
	NextLock:
		for _, lock := range locks {
			for i := 0; i < 5; i++ {
				if key[i]+lock[i] > 5 {
					continue NextLock
				}
			}
			total++
		}
	}
	return total
}
