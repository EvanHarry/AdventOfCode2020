package puzzles

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// DayOnePuzzleOne - https://adventofcode.com/2020/day/1
func DayOnePuzzleOne() {
	file, err := os.Open("data/day1.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var entries []int

	for scanner.Scan() {
		strV := scanner.Text()
		intV, err := strconv.Atoi(strV)

		if err == nil {
			entries = append(entries, intV)
		}
	}

	for k, v := range entries {
		for _k, _v := range entries {
			if k == _k {
				continue
			}

			sum := v + _v

			if sum == 2020 {
				product := v * _v

				fmt.Printf("%v and %v have a sum of %v and a product of %v\n", v, _v, sum, product)

				return
			}
		}
	}
}
