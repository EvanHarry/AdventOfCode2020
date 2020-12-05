package puzzles

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type day4Puzzle1Passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func (p day4Puzzle1Passport) isValid() bool {
	missingFields := 0

	if p.byr == "" {
		missingFields++
	}
	if p.iyr == "" {
		missingFields++
	}
	if p.eyr == "" {
		missingFields++
	}
	if p.hgt == "" {
		missingFields++
	}
	if p.hcl == "" {
		missingFields++
	}
	if p.ecl == "" {
		missingFields++
	}
	if p.pid == "" {
		missingFields++
	}
	if p.cid == "" {
		missingFields++
	}

	if missingFields == 0 {
		return true
	} else if missingFields == 1 && p.cid == "" {
		return true
	} else {
		return false
	}
}

type day4Puzzle1PassportEntry struct {
	key   string
	value string
}

// Day4Puzzle1 - https://adventofcode.com/2020/day/4
func Day4Puzzle1() {
	file, err := os.Open("data/day4.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var rows [][]string
	blanks := 0

	var current []string
	for scanner.Scan() {
		row := scanner.Text()

		if row == "" {
			rows = append(rows, current)
			current = []string{}
			blanks++
			continue
		}

		cols := strings.Split(row, " ")

		for _, v := range cols {
			current = append(current, v)
		}
	}

	// We have to add the last item to the array manually - no blank line to trigger in loop.
	rows = append(rows, current)

	var passports []day4Puzzle1Passport
	for _, v := range rows {
		passport := day4Puzzle1Passport{}
		for _, _v := range v {
			parse := strings.Split(_v, ":")
			entry := day4Puzzle1PassportEntry{
				parse[0],
				parse[1],
			}

			switch entry.key {
			case "byr":
				passport.byr = entry.value
			case "iyr":
				passport.iyr = entry.value
			case "eyr":
				passport.eyr = entry.value
			case "hgt":
				passport.hgt = entry.value
			case "hcl":
				passport.hcl = entry.value
			case "ecl":
				passport.ecl = entry.value
			case "pid":
				passport.pid = entry.value
			case "cid":
				passport.cid = entry.value
			}
		}

		passports = append(passports, passport)
	}

	count := 0
	for _, v := range passports {
		valid := v.isValid()

		if valid {
			count++
		}
	}

	fmt.Printf("There are %v valid passports\n", count)
}
