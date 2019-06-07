package converter

import (
	"fmt"
	"sort"
)

var numerals = map[int]string{
	1000: "M",
	500:  "D",
	100:  "C",
	50:   "L",
	10:   "X",
	5:    "V",
	1:    "I",
}

func Convert(number int) string {
	var rn string

	fmt.Printf("%v\n", sortedKeys())

	for {
		for _, d := range sortedKeys() {
			n := numerals[d]
			fmt.Printf("d: %d, n: %s\n", d, n)
			if number >= d {
				rn += n
				number -= d
				break
			} else if number >= d-1 {
				rn += n
				number -= d-2
				break
			}
		}

		fmt.Printf("number: %d\n", number)
		if number <= 0 {
			break
		}
	}

	return rn
}

func sortedKeys() []int {
	var k []int

	for d := range numerals {
		k = append(k, d)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(k)))

	return k
}