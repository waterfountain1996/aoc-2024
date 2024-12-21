package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	b, err := os.ReadFile("inputs/day21.txt")
	if err != nil {
		panic(err)
	}
	p1, p2 := solve(string(b))
	fmt.Printf("%s\n%s\n", p1, p2)
}

func solve(input string) (string, string) {
	codes := parseInput(input)
	p1 := solve2(codes, 3)
	p2 := solve2(codes, 26)
	return strconv.FormatInt(p1, 10), strconv.FormatInt(p2, 10)
}

func parseInput(input string) []string {
	input = strings.TrimSpace(input)
	return strings.Split(input, "\n")
}

type move struct {
	Button rune
	Dir    rune
}

type keypad func(r rune) []move

func numericKeypad(r rune) []move {
	switch r {
	case 'A':
		return []move{
			{'3', '^'},
			{'0', '<'},
		}
	case '0':
		return []move{
			{'2', '^'},
			{'A', '>'},
		}
	case '1':
		return []move{
			{'4', '^'},
			{'2', '>'},
		}
	case '2':
		return []move{
			{'5', '^'},
			{'3', '>'},
			{'0', 'v'},
			{'1', '<'},
		}
	case '3':
		return []move{
			{'6', '^'},
			{'A', 'v'},
			{'2', '<'},
		}
	case '4':
		return []move{
			{'7', '^'},
			{'5', '>'},
			{'1', 'v'},
		}
	case '5':
		return []move{
			{'8', '^'},
			{'6', '>'},
			{'2', 'v'},
			{'4', '<'},
		}
	case '6':
		return []move{
			{'9', '^'},
			{'3', 'v'},
			{'5', '<'},
		}
	case '7':
		return []move{
			{'8', '>'},
			{'4', 'v'},
		}
	case '8':
		return []move{
			{'9', '>'},
			{'5', 'v'},
			{'7', '<'},
		}
	case '9':
		return []move{
			{'6', 'v'},
			{'8', '<'},
		}
	default:
		panic(fmt.Sprintf("%c on a numeric keypad", r))
	}
}

func directionKeypad(r rune) []move {
	switch r {
	case 'A':
		return []move{
			{'>', 'v'},
			{'^', '<'},
		}
	case '^':
		return []move{
			{'A', '>'},
			{'v', 'v'},
		}
	case '>':
		return []move{
			{'A', '^'},
			{'v', '<'},
		}
	case 'v':
		return []move{
			{'^', '^'},
			{'<', '<'},
			{'>', '>'},
		}
	case '<':
		return []move{
			{'v', '>'},
		}
	default:
		panic(fmt.Sprintf("%c on a numeric keypad", r))
	}
}

func solve2(codes []string, n int) int64 {
	var total int64
	for _, code := range codes {
		var length int64
		v, _ := strconv.ParseInt(strings.Trim(code, "A"), 10, 32)

		start := 'A'
		for _, end := range code {
			var r int64 = 1<<63 - 1
			for _, seq := range bfs(start, end, numericKeypad) {
				if nr := cheapestN(seq, n); nr < r {
					r = nr
				}
			}
			length += r
			start = end
		}

		total += length * v
	}
	return total
}

func cheapestN(seq string, n int) int64 {
	if n == 1 {
		return int64(len(seq))
	}

	var answer int64

	start := 'A'
	for _, end := range seq {
		answer += cheapest(start, end, n)
		start = end
	}

	return answer
}

func cheapest(start, end rune, n int) (r int64) {
	tup := cachetup{start, end, n}
	if v, ok := cache[tup]; ok {
		return v
	}
	defer func() {
		cache[tup] = r
	}()

	answer := int64(1<<63 - 1)
	for _, seq := range bfs(start, end, directionKeypad) {
		if v := cheapestN(seq, n-1); v < answer {
			answer = v
		}
	}
	return answer
}

type cachetup struct {
	s rune
	e rune
	n int
}

var cache = make(map[cachetup]int64)

type state struct {
	Button rune
	Moves  []move
}

func bfs(start, end rune, kp keypad) []string {
	var (
		result []string
		queue  = []state{{Button: start}}
	)
	for len(queue) > 0 {
		it := queue[0]
		queue = queue[1:]

		if it.Button == end {
			n := len(it.Moves) + 1
			way := make([]rune, 0, n)
			for _, mv := range it.Moves {
				way = append(way, mv.Dir)
			}
			way = append(way, 'A')
			result = append(result, string(way))
			continue
		}

		for _, to := range kp(it.Button) {
			if slices.ContainsFunc(it.Moves, func(m move) bool {
				return m.Button == to.Button
			}) {
				continue
			}

			moves := append([]move(nil), it.Moves...)
			moves = append(moves, to)

			queue = append(queue, state{
				Button: to.Button,
				Moves:  moves,
			})
		}
	}
	return result
}
