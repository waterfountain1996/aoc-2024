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
	var (
		lines = normalizeInput(input)
		p1    = solve_part1(lines)
		p2    = solve_part2(lines)
	)
	return strconv.Itoa(p1), strconv.Itoa(p2)
}

func normalizeInput(input string) []string {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	rows := make([]string, len(lines)+6)
	for i, line := range lines {
		rows[i+3] = "..." + line + "..."
	}
	for i := range 3 {
		row := strings.Repeat(".", len(rows[i+3]))
		rows[i], rows[len(rows)-1-i] = row, row
	}
	return rows
}

func solve_part1(lines []string) int {
	seen := make(map[wordloc]struct{})

	for row, line := range lines {
		for col, ch := range line {
			if ch != 'X' {
				continue
			}

			if lines[row][col:col+4] == "XMAS" {
				seen[wordloc{point{row, col}, dirRight}] = struct{}{}
			}
			if lines[row][col-3:col+1] == "SAMX" {
				seen[wordloc{point{row, col}, dirLeft}] = struct{}{}
			}

			{
				word := [...]byte{
					lines[row-0][col],
					lines[row-1][col],
					lines[row-2][col],
					lines[row-3][col],
				}
				if s := string(word[:]); s == "XMAS" {
					seen[wordloc{point{row, col}, dirUp}] = struct{}{}
				}
			}

			{
				word := [...]byte{
					lines[row+0][col],
					lines[row+1][col],
					lines[row+2][col],
					lines[row+3][col],
				}
				if s := string(word[:]); s == "XMAS" {
					seen[wordloc{point{row, col}, dirDown}] = struct{}{}
				}
			}

			{
				word := [...]byte{
					lines[row-0][col-0],
					lines[row-1][col-1],
					lines[row-2][col-2],
					lines[row-3][col-3],
				}
				if s := string(word[:]); s == "XMAS" {
					seen[wordloc{point{row, col}, dirUpLeft}] = struct{}{}
				}
			}

			{
				word := [...]byte{
					lines[row-0][col+0],
					lines[row-1][col+1],
					lines[row-2][col+2],
					lines[row-3][col+3],
				}
				if s := string(word[:]); s == "XMAS" {
					seen[wordloc{point{row, col}, dirUpRight}] = struct{}{}
				}
			}

			{
				word := [...]byte{
					lines[row+0][col+0],
					lines[row+1][col+1],
					lines[row+2][col+2],
					lines[row+3][col+3],
				}
				if s := string(word[:]); s == "XMAS" {
					seen[wordloc{point{row, col}, dirDownRight}] = struct{}{}
				}
			}

			{
				word := [...]byte{
					lines[row+0][col-0],
					lines[row+1][col-1],
					lines[row+2][col-2],
					lines[row+3][col-3],
				}
				if s := string(word[:]); s == "XMAS" {
					seen[wordloc{point{row, col}, dirDownLeft}] = struct{}{}
				}
			}
		}
	}

	return len(seen)
}

func solve_part2(lines []string) int {
	var total int

	for row, line := range lines {
		for col, ch := range line {
			if ch != 'A' {
				continue
			}
			adjacent := map[direction]byte{
				dirUpLeft:    lines[row-1][col-1],
				dirUpRight:   lines[row-1][col+1],
				dirDownRight: lines[row+1][col+1],
				dirDownLeft:  lines[row+1][col-1],
			}
			var (
				isMAS1 = adjacent[dirUpLeft] == 'M' && adjacent[dirDownRight] == 'S'
				isSAM1 = adjacent[dirUpLeft] == 'S' && adjacent[dirDownRight] == 'M'
				isMAS2 = adjacent[dirUpRight] == 'M' && adjacent[dirDownLeft] == 'S'
				isSAM2 = adjacent[dirUpRight] == 'S' && adjacent[dirDownLeft] == 'M'
			)
			if (isMAS1 || isSAM1) && (isMAS2 || isSAM2) {
				total++
			}
		}
	}

	return total
}

type wordloc struct {
	Point point
	Dir   direction
}

type point [2]int

type direction int

const (
	dirLeft direction = iota
	dirRight
	dirUp
	dirDown
	dirUpLeft
	dirUpRight
	dirDownLeft
	dirDownRight
)
