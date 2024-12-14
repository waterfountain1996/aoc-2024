package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

const (
	width  = 101
	height = 103
)

func main() {
	robots := parseInput(input)

	scanner := bufio.NewScanner(os.Stdin)

	var seconds int
	for {
		for i, rob := range robots {
			pos, vel := rob[0], rob[1]

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
			robots[i][0] = pos
		}

		filename := fmt.Sprintf("images/%05d.jpeg", seconds)
		writeImage(filename, imageFromRobots(robots, width, height))

		seconds++

		if seconds%100 == 0 {
			fmt.Printf("Type 'break' if you found one:")
			scanner.Scan()
			if text := scanner.Text(); strings.ToLower(text) == "break" {
				break
			}
		}
	}
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

func imageFromRobots(robots [][2]point, width, height int) image.Image {
	im := image.NewGray(image.Rect(0, 0, width, height))
	for _, rob := range robots {
		pos := rob[0]
		im.SetGray(pos.X, pos.Y, color.Gray{Y: 255})
	}
	return im
}

func writeImage(filename string, im image.Image) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := jpeg.Encode(f, im, nil); err != nil {
		panic(err)
	}

	f.Sync()
}
