package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	b, err := os.ReadFile("inputs/day17.txt")
	if err != nil {
		panic(err)
	}
	p1, p2 := solve(string(b))
	fmt.Printf("%s\n%s\n", p1, p2)
}

func solve(input string) (string, string) {
	comp := parseInput(input)
	p1 := solve_part1(comp)
	p2 := solve_part2(comp)
	return p1, strconv.FormatUint(p2, 10)
}

const (
	opADV = uint8(0)
	opBXL = uint8(1)
	opBST = uint8(2)
	opJNZ = uint8(3)
	opBXC = uint8(4)
	opOUT = uint8(5)
	opBDV = uint8(6)
	opCDV = uint8(7)
)

type computer struct {
	regA         uint64
	regB         uint64
	regC         uint64
	instructions []uint8
	pc           uint8
}

func (comp *computer) Reset(reg uint64) {
	comp.regA = reg
	comp.pc = 0
}

func (comp computer) Combo() uint64 {
	arg := comp.Literal()
	switch arg {
	case 4:
		arg = comp.regA
	case 5:
		arg = comp.regB
	case 6:
		arg = comp.regC
	}
	return arg
}

func (comp computer) Literal() uint64 {
	return uint64(comp.instructions[comp.pc+1])
}

func (comp computer) Emit() uint64 {
	comp.regB = comp.regA % 8
	comp.regB = comp.regB ^ 1
	comp.regC = uint64(float64(comp.regA) / math.Pow(2.0, float64(comp.regB)))
	comp.regA >>= 3
	comp.regB = comp.regB ^ comp.regC
	comp.regB = comp.regB ^ 6
	return comp.regB % 8
}

func parseInput(input string) computer {
	input = strings.TrimSpace(input)
	s1, s2, _ := strings.Cut(input, "\n\n")

	var comp computer

	regs := strings.SplitN(s1, "\n", 3)
	comp.regA, _ = strconv.ParseUint(strings.TrimPrefix(regs[0], "Register A: "), 10, 64)
	comp.regB, _ = strconv.ParseUint(strings.TrimPrefix(regs[0], "Register B: "), 10, 64)
	comp.regC, _ = strconv.ParseUint(strings.TrimPrefix(regs[0], "Register C: "), 10, 64)

	s2 = strings.TrimPrefix(s2, "Program: ")
	for _, v := range strings.Split(s2, ",") {
		n, _ := strconv.ParseUint(v, 10, 7)
		comp.instructions = append(comp.instructions, uint8(n))
	}
	return comp
}

func execute(comp computer) []uint64 {
	var out []uint64
	for int(comp.pc) < len(comp.instructions) {
		switch comp.instructions[comp.pc] {
		case opADV:
			arg := comp.Combo()
			res := float64(comp.regA) / math.Pow(2.0, float64(arg))
			comp.regA = uint64(res)
		case opBXL:
			arg := comp.Literal()
			comp.regB ^= arg
		case opBST:
			arg := comp.Combo()
			comp.regB = arg & 0b111
		case opJNZ:
			if comp.regA != 0 {
				comp.pc = uint8(comp.Literal())
				continue
			}
		case opBXC:
			arg1, arg2 := comp.regB, comp.regC
			comp.regB = arg1 ^ arg2
		case opOUT:
			arg := comp.Combo()
			out = append(out, arg&0b111)
		case opBDV:
			arg := comp.Combo()
			res := float64(comp.regA) / math.Pow(2.0, float64(arg))
			comp.regB = uint64(res)
		case opCDV:
			arg := comp.Combo()
			res := float64(comp.regA) / math.Pow(2.0, float64(arg))
			comp.regC = uint64(res)
		}

		comp.pc += 2
	}
	return out
}

func solve_part1(comp computer) string {
	out := execute(comp)
	var b strings.Builder
	for _, n := range out {
		if b.Len() > 0 {
			b.WriteRune(',')
		}
		b.WriteString(strconv.FormatUint(n, 10))
	}
	return b.String()
}

func solve_part2(comp computer) uint64 {
	rprog := slices.Clone(comp.instructions)
	slices.Reverse(rprog)

	var reg uint64

	n := len(comp.instructions)
	for i := range n {
		want := make([]uint64, 0, i+1)
		for j := n - 1 - i; j < n; j++ {
			want = append(want, uint64(comp.instructions[j]))
		}

		for j := 0; true; j++ {
			a := (reg << 3) + uint64(j)
			comp.Reset(a)

			out := execute(comp)
			if slices.Equal(out, want) {
				reg = a
				break
			}
		}
	}

	return reg
}
