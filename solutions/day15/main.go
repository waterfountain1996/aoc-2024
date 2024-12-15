package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	p1, p2 := solve(b)
	fmt.Printf("%s\n%s\n", p1, p2)
}

func solve(input []byte) (string, string) {
	tiles, moves := parseInput(bytes.Clone(input))
	p1 := solve_part1(tiles, moves)
	tiles, moves = parseInput(bytes.Clone(input))
	p2 := solve_part2(tiles, moves)
	return strconv.Itoa(p1), strconv.Itoa(p2)
}

type direction int

const (
	dirUp = direction(iota)
	dirRight
	dirDown
	dirLeft
)

func (dir direction) String() string {
	switch dir {
	case dirUp:
		return "^"
	case dirRight:
		return ">"
	case dirDown:
		return "v"
	case dirLeft:
		return "<"
	default:
		return ""
	}
}

type tileMap [][]byte

func parseInput(input []byte) (tileMap, []byte) {
	rawmap, moves, _ := bytes.Cut(bytes.TrimSpace(input), []byte{'\n', '\n'})
	return bytes.Split(rawmap, []byte{'\n'}), moves
}

func solve_part1(tiles tileMap, moves []byte) int {
	var sy, sx int
	for y, row := range tiles {
		for x, col := range row {
			if col == '@' {
				sy, sx = y, x
				break
			}
		}
	}

	for _, m := range moves {
		var dy, dx int
		switch m {
		case '^':
			dy = -1
		case '>':
			dx = 1
		case 'v':
			dy = 1
		case '<':
			dx = -1
		default:
			continue
		}

		y, x := sy, sx
		var boxes bool
	Push:
		for {
			y += dy
			x += dx

			switch tiles[y][x] {
			case 'O':
				boxes = true
			case '#':
				break Push
			case '.':
				tiles[sy][sx] = '.'
				tiles[sy+dy][sx+dx] = '@'

				sy += dy
				sx += dx

				if boxes {
					tiles[y][x] = 'O'
				}
				break Push
			}
		}
	}

	var total int
	for y, row := range tiles {
		for x, col := range row {
			if col == 'O' {
				gps := 100*y + x
				total += gps
			}
		}
	}
	return total
}

func widenMap(tiles tileMap) tileMap {
	newMap := make(tileMap, len(tiles))
	for y, row := range tiles {
		newRow := make([]byte, len(row)*2)

		for x, col := range row {
			var a, b byte
			switch col {
			case '#':
				a, b = '#', '#'
			case 'O':
				a, b = '[', ']'
			case '.':
				a, b = '.', '.'
			case '@':
				a, b = '@', '.'
			}
			newRow[x*2] = a
			newRow[x*2+1] = b
		}

		newMap[y] = newRow
	}
	return newMap
}

func solve_part2(tiles tileMap, moves []byte) int {
	tiles = widenMap(tiles)

	var sy, sx int
	for y, row := range tiles {
		for x, col := range row {
			if col == '@' {
				sy, sx = y, x
				break
			}
		}
	}

	for _, m := range moves {
		var dir direction
		switch m {
		case '^':
			dir = dirUp
		case '>':
			dir = dirRight
		case 'v':
			dir = dirDown
		case '<':
			dir = dirLeft
		default:
			continue
		}

		if ny, nx, ok := maybePush(tiles, sy, sx, dir); ok {
			sy, sx = ny, nx
		}
	}

	var total int
	for y, row := range tiles {
		for x, col := range row {
			if col == '[' {
				total += 100*y + x
			}
		}
	}
	return total
}

type point struct {
	Y, X int
}

func maybePush(tiles tileMap, y, x int, dir direction) (int, int, bool) {
	seen := make(map[point]struct{})

	queue := []point{{y, x}}
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		seen[p] = struct{}{}

		switch dir {
		case dirUp:
			switch tiles[p.Y-1][p.X] {
			case '#':
				return 0, 0, false
			case '[':
				queue = append(queue, point{p.Y - 1, p.X}, point{p.Y - 1, p.X + 1})
			case ']':
				queue = append(queue, point{p.Y - 1, p.X}, point{p.Y - 1, p.X - 1})
			}
		case dirRight:
			switch tiles[p.Y][p.X+1] {
			case '#':
				return 0, 0, false
			case '[':
				seen[point{p.Y, p.X + 1}] = struct{}{}
				queue = append(queue, point{p.Y, p.X + 2})
			}
		case dirDown:
			switch tiles[p.Y+1][p.X] {
			case '#':
				return 0, 0, false
			case '[':
				queue = append(queue, point{p.Y + 1, p.X}, point{p.Y + 1, p.X + 1})
			case ']':
				queue = append(queue, point{p.Y + 1, p.X}, point{p.Y + 1, p.X - 1})
			}
		case dirLeft:
			switch tiles[p.Y][p.X-1] {
			case '#':
				return 0, 0, false
			case ']':
				seen[point{p.Y, p.X - 1}] = struct{}{}
				queue = append(queue, point{p.Y, p.X - 2})
			}
		}
	}

	var ny, nx int

	replace := make(map[point]byte)
	for p := range seen {

		var dy, dx int
		switch dir {
		case dirUp:
			dy = -1
		case dirRight:
			dx = 1
		case dirDown:
			dy = 1
		case dirLeft:
			dx = -1
		}

		if tiles[p.Y][p.X] == '@' {
			ny, nx = p.Y+dy, p.X+dx
		}

		prev := point{p.Y - dy, p.X - dx}
		if _, ok := seen[prev]; ok {
			replace[p] = tiles[prev.Y][prev.X]
		} else {
			replace[p] = '.'
		}

		next := point{p.Y + dy, p.X + dx}
		if _, ok := seen[next]; !ok {
			replace[next] = tiles[p.Y][p.X]
		}
	}

	for p, r := range replace {
		tiles[p.Y][p.X] = r
	}

	return ny, nx, true
}
