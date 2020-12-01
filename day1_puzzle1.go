package main

import "fmt"

func main() {
	entries := []int{1721, 979, 366, 299, 675, 1456}

	for k, v := range entries {
		for _k, _v := range entries {
			if k == _k {
				continue
			}

			sum := v + _v

			if sum == 2020 {
				product := v * _v

				fmt.Printf("%v and %v have a sum of %v and a product of %v.", v, _v, sum, product)

				return
			}
		}
	}
}
