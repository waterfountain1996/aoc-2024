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
	p1 := solve_part1(parseInput(input))
	p2 := solve_part2(parseInput(input))
	return strconv.Itoa(p1), strconv.FormatInt(p2, 10)
}

func parseInput(input string) diskMap {
	input = strings.TrimSpace(input)
	diskMap := make([]int, len(input))
	for i, r := range input {
		diskMap[i] = int(r - '0')
	}
	return diskMap
}

type diskMap []int

func (d diskMap) Filesize(id int) int {
	return d[id*2]
}

func (d diskMap) MaxFile() int {
	n := len(d)
	if n%2 != 0 {
		n -= 1
	}
	return n / 2
}

func solve_part1(disk diskMap) int {
	// Get last non-empty file.
	end := disk.MaxFile()
	for disk.Filesize(end) == 0 {
		end--
	}

	var (
		start    int
		cursor   int
		checksum int
	)
	for i := 0; i < len(disk); i++ {

		if i%2 == 0 {
			start = i / 2
			if start > end {
				break
			}
			nblocks := disk.Filesize(start)
			for j := 0; j < nblocks; j++ {
				checksum += cursor * start
				cursor++
			}
		} else {
			nblocks := disk[i] // Size of current free sector.
			for j := 0; j < nblocks; j++ {
				if end <= start {
					break
				}
				checksum += cursor * end

				// Decrement blocksize of the last file
				endPos := end * 2
				disk[endPos]--

				// If it's fully rearranged, decrement the pointer.
				for disk.Filesize(end) == 0 {
					end--
				}
				cursor++
			}
		}
	}
	return checksum
}

func solve_part2(disk diskMap) int64 {
	sectors := make([][]int, len(disk)/2+1)
	for id := disk.MaxFile(); id > 0; id-- {
		for i := 1; i < len(disk) && i < id*2; i += 2 {
			free := disk[i]
			if size := disk.Filesize(id); free >= size {
				disk[i] -= size
				sectors[(i-1)/2] = append(sectors[(i-1)/2], id)
				break
			}
		}
	}

	var (
		seen     = make(map[int]bool)
		cursor   int64
		checksum int64
	)
	for i := 0; i < len(disk); i++ {
		if i%2 == 0 {
			var (
				id      = i / 2
				nblocks = disk.Filesize(id)
			)
			if seen[id] {
				cursor += int64(nblocks)
				continue
			}
			for j := 0; j < nblocks; j++ {
				checksum += cursor * int64(id)
				cursor++
			}
		} else {
			for _, id := range sectors[(i-1)/2] {
				seen[id] = true
				for j := 0; j < disk.Filesize(id); j++ {
					checksum += cursor * int64(id)
					cursor++
				}
			}
			left := disk[i]
			cursor += int64(left)
		}
	}
	return checksum
}
