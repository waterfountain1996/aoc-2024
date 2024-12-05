package main

import "testing"

var testInput = `
47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47
`

func TestSolve_Part1(t *testing.T) {
	rules, updates := parseInput(testInput)
	want := 143
	if res := solve_part1(rules, updates); res != want {
		t.Fatalf("unexpected result:\nwant: %d\nhave: %d", want, res)
	}
}

func TestSolve_Part2(t *testing.T) {
	rules, updates := parseInput(testInput)
	want := 123
	if res := solve_part2(rules, updates); res != want {
		t.Fatalf("unexpected result:\nwant: %d\nhave: %d", want, res)
	}
}
