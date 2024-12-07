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
	eqs := parseInput(input)
	p1 := solve_part1(eqs)
	p2 := solve_part2(eqs)
	return strconv.FormatInt(p1, 10), strconv.FormatInt(p2, 10)
}

func parseInput(input string) []equation {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	eqs := make([]equation, len(lines))
	for i, line := range lines {
		v, tail, _ := strings.Cut(line, ": ")

		eqs[i].Value, _ = strconv.ParseInt(v, 10, 64)

		elems := strings.Split(tail, " ")
		eqs[i].Arguments = make([]int64, len(elems))
		for j, e := range elems {
			eqs[i].Arguments[j], _ = strconv.ParseInt(e, 10, 64)
		}
	}
	return eqs
}

type equation struct {
	Value     int64
	Arguments []int64
}

type operator int

const (
	opMul = operator(iota)
	opAdd
	opConcat
	opMax
)

func (op operator) String() string {
	switch op {
	case opMul:
		return "*"
	case opAdd:
		return "+"
	case opConcat:
		return "||"
	default:
		return ""
	}
}

func solve_part1(eqs []equation) int64 {
	var total int64
	for _, eq := range eqs {
		if isValid(eq) {
			total += eq.Value
		}
	}
	return total
}

func isValid(eq equation) bool {
	nops := len(eq.Arguments) - 1
	for i := 0; i < (1 << nops); i++ {
		acc := eq.Arguments[0]
		for j := range nops {
			mask := 1 << j
			op := opMul
			if (i & mask) > 0 {
				op = opAdd
			}
			switch op {
			case opMul:
				acc *= eq.Arguments[1+j]
			case opAdd:
				acc += eq.Arguments[1+j]
			}
		}
		if acc == eq.Value {
			return true
		}
	}
	return false
}

func solve_part2(eqs []equation) int64 {
	var total int64
	for _, eq := range eqs {
		if isValid2(eq) {
			total += eq.Value
		}
	}
	return total
}

func isValid2(eq equation) bool {
	nops := len(eq.Arguments) - 1

	prod, ok := prodCache[nops]
	if !ok {
		prod = product([]operator{opMul, opAdd, opConcat}, nops)
		prodCache[nops] = prod
	}

	for _, ops := range prod {
		acc := eq.Arguments[0]

		for i, arg := range eq.Arguments[1:] {
			switch ops[i] {
			case opAdd:
				acc += arg
			case opMul:
				acc *= arg
			case opConcat:
				acc = concat(acc, arg)
			}
		}

		if acc == eq.Value {
			return true
		}
	}

	return false
}

var prodCache = make(map[int][][]operator)

func concat(a, b int64) int64 {
	temp := b
	for temp > 0 {
		a *= 10
		temp /= 10
	}
	return a + b
}

// translated Python's itertools.product implementation.
func product(ops []operator, n int) [][]operator {
	if n <= 0 {
		panic("ur crazy")
	}

	var res [][]operator = make([][]operator, n)

	for i := 0; i < n; i++ {
		var tmpRes [][]operator
		for _, comb := range res {
			for _, op := range ops {
				newComb := append([]operator(nil), comb...)
				newComb = append(newComb, op)
				tmpRes = append(tmpRes, newComb)
			}
			res = tmpRes
		}
	}

	return res
}
