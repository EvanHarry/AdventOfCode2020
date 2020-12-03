package puzzles

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type day2Puzzle2Entry struct {
	first    int
	last     int
	letter   string
	password string
	valid    bool
}

// DayTwoPuzzleTwo - https://adventofcode.com/2020/day/2
func DayTwoPuzzleTwo() {
	file, err := os.Open("data/day2.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var entries []day2Puzzle2Entry

	for scanner.Scan() {
		strV := scanner.Text()
		cols := strings.Split(strV, " ")

		firstLast := strings.Split(cols[0], "-")
		first, _ := strconv.Atoi(firstLast[0])
		last, _ := strconv.Atoi(firstLast[1])

		letter := strings.ReplaceAll(cols[1], ":", "")

		password := cols[2]

		entry := day2Puzzle2Entry{
			first,
			last,
			letter,
			password,
			false,
		}

		entries = append(entries, entry)
	}

	count := 0
	for _, v := range entries {
		if string(v.password[v.first-1]) == v.letter {
			v.valid = true
		}

		if string(v.password[v.last-1]) == v.letter {
			if v.valid {
				v.valid = false
			} else {
				v.valid = true
			}
		}

		if v.valid {
			count++
		}
	}

	fmt.Printf("There are %v valid entries\n", count)
}
