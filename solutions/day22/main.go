package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	b, err := os.ReadFile("inputs/day22.txt")
	if err != nil {
		panic(err)
	}
	p1, p2 := solve(string(b))
	fmt.Printf("%s\n%s\n", p1, p2)
}

func solve(input string) (string, string) {
	secrets := parseInput(input)
	p1, p2 := solve2(secrets)
	return strconv.FormatInt(p1, 10), strconv.FormatInt(p2, 10)
}

func parseInput(input string) []int64 {
	input = strings.TrimSpace(input)
	var out []int64
	for _, line := range strings.Split(input, "\n") {
		n, _ := strconv.ParseInt(line, 10, 64)
		out = append(out, n)
	}
	return out
}

var cache = make(map[int64]int64)

func next(secret int64) int64 {
	if r, ok := cache[secret]; ok {
		return r
	}
	result := secret

	v := result * 64
	result = (v ^ result) % 16777216

	v = result / 32
	result = (v ^ result) % 16777216

	v = result * 2048
	result = (v ^ result) % 16777216

	cache[secret] = result
	return result
}

func hash4(a, b, c, d int64) int64 {
	a += 9
	b += 9
	c += 9
	d += 9
	return a*19*19*19 + b*19*19 + c*19 + d
}

func solve2(secrets []int64) (int64, int64) {
	var total int64

	results := make(map[int64]int64)

	for _, secret := range secrets {
		var (
			a, b, c, d int64
			seen       = make(map[int64]bool)
			prev       = secret % 10
		)
		for i := range 2000 {
			secret = next(secret)
			var (
				price  = secret % 10
				change = price - prev
			)

			a, b, c, d = b, c, d, change
			if i >= 3 {
				key := hash4(a, b, c, d)
				if !seen[key] {
					results[key] += price
					seen[key] = true
				}
			}

			prev = price
		}
		total += secret
	}

	var bananas int64 = -1
	for _, v := range results {
		if bananas == -1 || v > bananas {
			bananas = v
		}
	}
	return total, bananas
}
