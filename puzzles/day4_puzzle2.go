package puzzles

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type day4Puzzle2Passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func (p day4Puzzle2Passport) isValid() bool {
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

func (p day4Puzzle2Passport) strictValidation() bool {
	invalidFields := 0

	// byr - 4 digits; at least 1920 and at most 2002.
	byr, err := strconv.ParseInt(p.byr, 0, 64)
	if err != nil {
		invalidFields++
	}

	if byr < 1920 || byr > 2002 {
		invalidFields++
	}

	// iyr - 4 digits; at least 2010 and at most 2020.
	iyr, err := strconv.ParseInt(p.iyr, 0, 64)
	if err != nil {
		invalidFields++
	}

	if iyr < 2010 || iyr > 2020 {
		invalidFields++
	}

	// eyr - 4 digits; at least 2020 and at most 2030.
	eyr, err := strconv.ParseInt(p.eyr, 0, 64)
	if err != nil {
		invalidFields++
	}

	if eyr < 2020 || eyr > 2030 {
		invalidFields++
	}

	// hgt - a number followed by either cm or in
	// if cm, the number must be at least 150 or at most 193.
	// if in, the number must be at least 59 and at most 76.
	if strings.Contains(p.hgt, "cm") {
		cmStr := p.hgt[:len(p.hgt)-2]
		cm, err := strconv.ParseInt(cmStr, 0, 64)
		if err != nil {
			invalidFields++
		}

		if cm < 150 || cm > 193 {
			invalidFields++
		}
	} else if strings.Contains(p.hgt, "in") {
		inStr := p.hgt[:len(p.hgt)-2]
		in, err := strconv.ParseInt(inStr, 0, 64)
		if err != nil {
			invalidFields++
		}

		if in < 59 || in > 76 {
			invalidFields++
		}
	} else {
		invalidFields++
	}

	// hcl - a # followed by exactly 6 characters 0-9 or a-f.
	if strings.Index(p.hcl, "#") == 0 && len(p.hcl) == 7 {
		hcl := strings.TrimPrefix(p.hcl, "#")
		if strings.Contains(hcl, "ghijklmnopqrstuvwxyz") {
			invalidFields++
		}
	} else {
		invalidFields++
	}

	// ecl - exactly one of amb blu brn gry grn hzl oth
	eyeColours := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	eclValid := false
	for _, v := range eyeColours {
		if v == p.ecl {
			eclValid = true
		}
	}
	if !eclValid {
		invalidFields++
	}

	// pid - a nine-digit number, including leading zeroes
	if len(p.pid) == 9 {
		_, err := strconv.ParseFloat(p.pid, 64)
		if err != nil {
			invalidFields++
		}
	} else {
		invalidFields++
	}

	return invalidFields == 0
}

type day4Puzzle2PassportEntry struct {
	key   string
	value string
}

// Day4Puzzle2 - https://adventofcode.com/2020/day/4
func Day4Puzzle2() {
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

	var passports []day4Puzzle2Passport
	for _, v := range rows {
		passport := day4Puzzle2Passport{}
		for _, _v := range v {
			parse := strings.Split(_v, ":")
			entry := day4Puzzle2PassportEntry{
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
		if v.isValid() && v.strictValidation() {
			count++
		}
	}

	fmt.Printf("There are %v valid passports\n", count)
}
