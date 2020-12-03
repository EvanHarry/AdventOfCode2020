package puzzles

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type day2Puzzle1Entry struct {
	min      int
	max      int
	letter   string
	password string
	valid    bool
}

// DayTwoPuzzleOne - https://adventofcode.com/2020/day/2
func DayTwoPuzzleOne() {
	file, err := os.Open("data/day2.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var entries []day2Puzzle1Entry

	for scanner.Scan() {
		strV := scanner.Text()
		cols := strings.Split(strV, " ")

		minMax := strings.Split(cols[0], "-")
		min, _ := strconv.Atoi(minMax[0])
		max, _ := strconv.Atoi(minMax[1])

		letter := strings.ReplaceAll(cols[1], ":", "")

		password := cols[2]

		entry := day2Puzzle1Entry{
			min,
			max,
			letter,
			password,
			false,
		}

		entries = append(entries, entry)
	}

	count := 0
	for _, v := range entries {
		repetitions := strings.Count(v.password, v.letter)

		if repetitions >= v.min && repetitions <= v.max {
			v.valid = true

			count++
		}
	}

	fmt.Printf("There are %v valid entries\n", count)
}
