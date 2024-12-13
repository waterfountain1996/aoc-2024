package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
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
	rules := parseInput(input)
	p1 := solve_part1(rules)
	p2 := solve_part2(rules)
	return strconv.Itoa(p1), strconv.FormatInt(p2, 10)
}

type point struct {
	Y, X int
}

func parseInput(input string) [][3]point {
	input = strings.TrimSpace(input)

	pat := regexp.MustCompile(`[XY][+=](\d+)`)

	var res [][3]point
	for _, elem := range strings.Split(input, "\n\n") {
		var rule [3]point
		for i, line := range strings.SplitN(elem, "\n", 3) {
			submatch := pat.FindAllStringSubmatch(line, 2)

			rule[i].X, _ = strconv.Atoi(submatch[0][1])
			rule[i].Y, _ = strconv.Atoi(submatch[1][1])
		}
		res = append(res, rule)
	}

	return res
}

func solve_part1(rules [][3]point) int {
	var total int
	for _, rule := range rules {
		btnA, btnB, prize := rule[0], rule[1], rule[2]
		a1, b1, c1 := float64(btnA.X), float64(btnB.X), float64(prize.X)
		a2, b2, c2 := float64(btnA.Y), float64(btnB.Y), float64(prize.Y)

		x, y := lineq(a1, b1, c1, a2, b2, c2)
		if x > 100 || math.Mod(x, 1.0) != 0.0 || y > 100 || math.Mod(y, 1.0) != 0.0 {
			continue
		}

		total += 3*int(x) + int(y)
	}
	return total
}

func solve_part2(rules [][3]point) int64 {
	var total int64
	for _, rule := range rules {
		btnA, btnB, prize := rule[0], rule[1], rule[2]
		a1, b1, c1 := float64(btnA.X), float64(btnB.X), float64(prize.X)+10000000000000
		a2, b2, c2 := float64(btnA.Y), float64(btnB.Y), float64(prize.Y)+10000000000000

		x, y := lineq(a1, b1, c1, a2, b2, c2)
		if math.Mod(x, 1.0) != 0.0 || math.Mod(y, 1.0) != 0.0 {
			continue
		}

		total += 3*int64(x) + int64(y)
	}
	return total
}

func lineq(a1, b1, c1, a2, b2, c2 float64) (float64, float64) {
	mul1, mul2 := b2, b1

	da := (a1 * mul1) - (a2 * mul2)
	dc := (c1 * mul1) - (c2 * mul2)

	x := dc / da
	y := (c1 - a1*x) / b1
	return x, y
}
