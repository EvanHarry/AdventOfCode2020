package puzzles

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

// Day3Puzzle1 - https://adventofcode.com/2020/day/3
func Day3Puzzle1() {
	file, err := os.Open("data/day3.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var rows []string

	for scanner.Scan() {
		row := scanner.Text()

		rows = append(rows, row)
	}

	xDiff := 3
	yDiff := 1
	xPos := 0
	yPos := 0
	hits := 0

	for k, v := range rows {
		// Re-size the matrix
		if (xPos + xDiff) >= len(v) {
			_xPos := float64(xPos)
			_xDiff := float64(xDiff)
			_vLen := float64(len(v))
			multiFactor := math.Ceil((_xPos + _xDiff) / _vLen)

			_multiFactor := int(multiFactor) + 1 // We add 1 to account for when (xPos + xDiff) == len(v)
			v = strings.Repeat(v, _multiFactor)
		}

		if (yPos + yDiff) != k {
			continue
		}

		xPos += xDiff
		yPos += yDiff

		symbol := fmt.Sprintf("%c", v[xPos])
		if symbol == "#" {
			hits++
		}
	}

	fmt.Printf("There were %v hits\n", hits)
}
