package puzzles

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type day3Puzzle2Slopes struct {
	xDiff int
	yDiff int
}

// Day3Puzzle2 - https://adventofcode.com/2020/day/3
func Day3Puzzle2() {
	slopes := []day3Puzzle2Slopes{
		{
			1,
			1,
		},
		{
			3,
			1,
		},
		{
			5,
			1,
		},
		{
			7,
			1,
		},
		{
			1,
			2,
		},
	}

	hits := 1
	for _, v := range slopes {
		_hits := countTrees(v.xDiff, v.yDiff)

		hits = hits * _hits
	}

	fmt.Printf("There were %v hits\n", hits)
}

func countTrees(xDiff, yDiff int) int {
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

	return hits
}
