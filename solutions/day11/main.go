package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
		os.Exit(1)
	}
	p1, p2 := solve(string(input))
	fmt.Printf("%s\n%s\n", p1, p2)
}

func solve(input string) (string, string) {
	p1 := blink(parseInput(input), 25)
	p2 := blink(parseInput(input), 75)
	return strconv.FormatInt(p1, 10), strconv.FormatInt(p2, 10)
}

func parseInput(input string) []int64 {
	elems := strings.Split(strings.TrimSpace(input), " ")
	stones := make([]int64, len(elems))
	for i, s := range elems {
		stones[i], _ = strconv.ParseInt(s, 10, 64)
	}
	return stones
}

func blink(stones []int64, nblinks int) int64 {
	var (
		total int64
		cache = make(map[pair]int64)
	)
	for _, stone := range stones {
		total += transform(stone, nblinks, cache)
	}
	return total
}

type pair struct {
	stone   int64
	nblinks int
}

func transform(stone int64, nblinks int, cache map[pair]int64) (r int64) {
	p := pair{stone, nblinks}
	if v, ok := cache[p]; ok {
		return v
	}
	defer func() {
		cache[p] = r
	}()

	if nblinks == 0 {
		return 1
	}

	digits := strconv.FormatInt(stone, 10)
	if len(digits)%2 != 0 {
		next := stone * 2024
		if next == 0 {
			next = 1
		}
		return transform(next, nblinks-1, cache)
	}

	half := len(digits) / 2
	s1, _ := strconv.ParseInt(digits[:half], 10, 64)
	s2, _ := strconv.ParseInt(digits[half:], 10, 64)
	return transform(s1, nblinks-1, cache) + transform(s2, nblinks-1, cache)
}
