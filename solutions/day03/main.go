package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var instrRegexp = regexp.MustCompile(`mul\((\d{1,3},\d{1,3})\)|do\(\)|don't\(\)`)

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
		p1 = doSolve(input, true)
		p2 = doSolve(input, false)
	)
	return strconv.Itoa(p1), strconv.Itoa(p2)
}

func doSolve(input string, part1 bool) int {
	var (
		submatch = instrRegexp.FindAllStringSubmatch(input, -1)
		enabled  = true
		total    int
	)
	for _, sm := range submatch {
		switch sm[0] {
		case "do()":
			enabled = true
		case "don't()":
			enabled = false
		default:
			if part1 || enabled {
				s1, s2, _ := strings.Cut(sm[1], ",")
				a, _ := strconv.Atoi(s1)
				b, _ := strconv.Atoi(s2)
				total += a * b
			}
		}
	}
	return total
}
