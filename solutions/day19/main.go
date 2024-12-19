package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	b, err := os.ReadFile("inputs/day19.txt")
	if err != nil {
		panic(err)
	}
	p1, p2 := solve(string(b))
	fmt.Printf("%s\n%s\n", p1, p2)
}

func solve(input string) (string, string) {
	patterns, designs := parseInput(input)
	p1 := solve_part1(patterns, designs)
	p2 := solve_part2(patterns, designs)
	return strconv.Itoa(p1), strconv.Itoa(p2)
}

func parseInput(input string) ([]string, []string) {
	input = strings.TrimSpace(input)
	s1, s2, _ := strings.Cut(input, "\n\n")
	return strings.Split(s1, ", "), strings.Split(s2, "\n")
}

type trie struct {
	r        rune
	end      bool
	children [26]*trie
}

func (t *trie) Match(s string) []string {
	var (
		matches []string
		prefix  []rune
	)

	node := t
	for _, r := range s {
		key := r - 'a'
		child := node.children[key]
		if child == nil {
			break
		}

		node = child
		prefix = append(prefix, node.r)

		if node.end {
			matches = append(matches, string(prefix))
		}
	}

	return matches
}

func (t *trie) Put(p string) {
	node := t
	for _, r := range p {
		key := r - 'a'
		child := node.children[key]
		if child == nil {
			child = &trie{r: r}
			node.children[key] = child
		}
		node = child
	}
	node.end = true
}

func solve_part1(patterns, designs []string) int {
	var total int

	var t trie
	for _, p := range patterns {
		t.Put(p)
	}

	for _, d := range designs {
		if dfs(&t, d) {
			total++
		}
	}

	return total
}

func dfs(t *trie, s string) bool {
	if s == "" {
		return true
	}

	for _, p := range t.Match(s) {
		if dfs(t, strings.TrimPrefix(s, p)) {
			return true
		}
	}
	return false
}

func solve_part2(patterns, designs []string) int {
	var t trie
	for _, p := range patterns {
		t.Put(p)
	}

	cache := make(map[string]int)

	for _, d := range designs {
		for i := len(d) - 1; i >= 0; i-- {
			sub := d[i:]
			if _, ok := cache[sub]; ok {
				continue
			}

			type state struct {
				s string
				n int
			}

			queue := []string{sub}
			for len(queue) > 0 {
				item := queue[0]
				queue = queue[1:]

				if item == "" {
					cache[sub]++
					continue
				}

				if v, ok := cache[item]; ok {
					cache[sub] += v
					continue
				}

				for _, p := range t.Match(item) {
					queue = append(queue, strings.TrimPrefix(item, p))
				}
			}
		}
	}

	var total int
	for _, k := range designs {
		total += cache[k]
	}
	return total
}
