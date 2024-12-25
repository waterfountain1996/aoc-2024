package main

import (
	"fmt"
	"maps"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	b, err := os.ReadFile("inputs/day24.txt")
	if err != nil {
		panic(err)
	}
	p1, p2 := solve(string(b))
	fmt.Printf("%s\n%s\n", p1, p2)
}

func solve(input string) (string, string) {
	wires, conns := parseInput(input)
	p1 := solve_part1(wires, conns)
	p2 := solve_part2(wires, conns)
	return strconv.FormatUint(p1, 10), p2
}

type gate struct {
	W1, W2 string
	Output string
	Op     string
}

func (c gate) String() string {
	return fmt.Sprintf("%s %s %s -> %s", c.W1, c.Op, c.W2, c.Output)
}

func parseInput(input string) (map[string]bool, map[string]gate) {
	input = strings.TrimSpace(input)
	sw, sg, _ := strings.Cut(input, "\n\n")

	wires := make(map[string]bool)
	for _, line := range strings.Split(sw, "\n") {
		name, bit, _ := strings.Cut(line, ": ")
		value, _ := strconv.ParseUint(bit, 10, 1)
		b := false
		if value > 0 {
			b = true
		}
		wires[name] = b
	}

	re := regexp.MustCompile(`(.*) (AND|OR|XOR) (.*) -> (.*)`)
	conns := make(map[string]gate)
	for _, line := range strings.Split(sg, "\n") {
		matches := re.FindAllStringSubmatch(line, -1)[0]
		w1, op, w2, wout := matches[1], matches[2], matches[3], matches[4]
		conns[wout] = gate{
			W1:     w1,
			W2:     w2,
			Output: wout,
			Op:     op,
		}
	}
	return wires, conns
}

var FAULTY []string

func resolveWire(levels map[string]bool, gates map[string]gate, w string) (out bool) {
	if l, ok := levels[w]; ok {
		return l
	}
	defer func() {
		levels[w] = out
	}()

	g := gates[w]

	if w[0] == 'z' && w != "z45" && g.Op != "XOR" {
		FAULTY = append(FAULTY, g.Output)
		fmt.Println("rule 1:", g)
	} else if w[0] != 'z' {
		if g.W1[0] != 'x' && g.W1[0] != 'y' && g.W2[0] != 'x' && g.W2[0] != 'y' {
			if g.Op == "XOR" {
				FAULTY = append(FAULTY, g.Output)
				fmt.Println("rule 2:", g)
			}
		}
	}

	a, b := resolveWire(levels, gates, g.W1), resolveWire(levels, gates, g.W2)
	switch g.Op {
	case "AND":
		return a && b
	case "OR":
		return a || b
	case "XOR":
		return a != b
	default:
		panic(fmt.Sprintf("unknown op: %s", g.Op))
	}
}

func solve_part1(wires map[string]bool, gates map[string]gate) uint64 {
	levels := maps.Clone(wires)
	var z uint64
	for w := range gates {
		if w[0] != 'z' {
			continue
		}

		if high := resolveWire(levels, gates, w); high {
			i, _ := strconv.Atoi(strings.TrimPrefix(w, "z"))
			mask := uint64(1 << i)
			z |= mask
		}
	}
	return z
}

func solve_part2(wires map[string]bool, gates map[string]gate) string {
	xor := findFaultyXOR(gates)
	FAULTY = append(FAULTY, xor.Output)
	fmt.Println("rule 3:", xor)

	and := findFaultyAND(gates)
	FAULTY = append(FAULTY, and.Output)
	fmt.Println("rule 4:", and)

	slices.Sort(FAULTY)
	return strings.Join(FAULTY, ",")
}

func findFaultyXOR(gates map[string]gate) gate {
	var possible []gate
	for _, g := range gates {
		if isXY(g.W1) && isXY(g.W2) && g.Op == "XOR" {
			if g.W1 == "x00" || g.W1 == "y00" || g.W2 == "x00" || g.W2 == "y00" {
				continue
			}
			possible = append(possible, g)
		}
	}

	found := make(map[gate]bool)
	for _, next := range gates {
		if next.Op == "XOR" {
			for _, g := range possible {
				if next.W1 == g.Output || next.W2 == g.Output {
					found[g] = true
				}
			}
		}
	}

	for _, g := range possible {
		if !found[g] {
			return g
		}
	}
	panic("unreachable")
}

func findFaultyAND(gates map[string]gate) gate {
	var possible []gate
	for _, g := range gates {
		if isXY(g.W1) && isXY(g.W2) && g.Op == "AND" {
			if g.W1 == "x00" || g.W1 == "y00" || g.W2 == "x00" || g.W2 == "y00" {
				continue
			}
			possible = append(possible, g)
		}
	}

	found := make(map[gate]bool)
	for _, next := range gates {
		if next.Op == "OR" {
			for _, g := range possible {
				if next.W1 == g.Output || next.W2 == g.Output {
					found[g] = true
				}
			}
		}
	}

	for _, g := range possible {
		if !found[g] {
			return g
		}
	}
	panic("unreachable")
}

func isXY(w string) bool {
	return w[0] == 'x' || w[0] == 'y'
}

func getXY(wires map[string]bool) (uint64, uint64) {
	var x, y uint64
	for w, high := range wires {
		prefix := w[0]
		if prefix != 'x' && prefix != 'y' {
			continue
		}

		if high {
			i, _ := strconv.Atoi(strings.TrimLeft(w, "xy"))
			mask := uint64(1 << i)
			switch prefix {
			case 'x':
				x |= mask
			case 'y':
				y |= mask
			}
		}
	}
	return x, y
}
