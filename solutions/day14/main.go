package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	p1 := solve(string(input))
	fmt.Printf("%s\n", p1)
}

func solve(input string) string {
	robots := parseInput(input)
	p1 := solve_part1(robots, 100, 101, 103)
	return strconv.Itoa(p1)
}

type point struct {
	Y, X int
}

func parseInput(input string) [][2]point {
	input = strings.TrimSpace(input)

	pat := regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)

	var res [][2]point
	for _, line := range strings.Split(input, "\n") {
		var pos, vel point
		submatch := pat.FindAllStringSubmatch(line, 2)

		pos.X, _ = strconv.Atoi(submatch[0][1])
		pos.Y, _ = strconv.Atoi(submatch[0][2])
		vel.X, _ = strconv.Atoi(submatch[0][3])
		vel.Y, _ = strconv.Atoi(submatch[0][4])
		res = append(res, [2]point{pos, vel})
	}

	return res
}

func solve_part1(robots [][2]point, seconds, width, height int) int {
	var (
		halfwidth  = (width - 1) / 2
		halfheight = (height - 1) / 2
	)
	var quadrants [4]int
	for _, rob := range robots {
		pos, vel := rob[0], rob[1]
		for range seconds {
			pos.X += vel.X
			if pos.X < 0 {
				pos.X = width + pos.X
			}
			pos.X %= width

			pos.Y += vel.Y
			if pos.Y < 0 {
				pos.Y = height + pos.Y
			}
			pos.Y %= height
		}

		if pos.X == halfwidth || pos.Y == halfheight {
			continue
		}

		switch {
		case pos.X < halfwidth && pos.Y < halfheight:
			quadrants[0]++
		case pos.X > halfwidth && pos.Y < halfheight:
			quadrants[1]++
		case pos.X < halfwidth && pos.Y > halfheight:
			quadrants[2]++
		case pos.X > halfwidth && pos.Y > halfheight:
			quadrants[3]++
		}
	}
	return quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
}
