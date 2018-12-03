package day2

import (
	"fmt"
	"helper"
	"testing"
)

func TestPartA(t *testing.T) {
	twoLetters, threeLetters := 0, 0
	lines, _ := helper.ReadLines("input")

	for _, value := range lines {
		letterCounts := CountAllLetters(value)
		two, three := false, false

		for _, count := range letterCounts {
			if count == 2 {
				two = true
			}
			if count == 3 {
				three = true
			}
		}

		if two {
			twoLetters++
		}
		if three {
			threeLetters++
		}
	}

	checksum := twoLetters * threeLetters

	fmt.Println("Part A: ", checksum)
}

func CountAllLetters(s string) map[rune]int {
	letterMap := make(map[rune]int)

	for _, letter := range s {
		_, letterExists := letterMap[letter]
		if letterExists {
			letterMap[letter]++
		} else {
			letterMap[letter] = 1
		}
	}
	return letterMap
}

func TestPartB(t *testing.T) {
	lines, _ := helper.ReadLines("input")

	for i1, line1 := range lines {
		for i2, line2 := range lines {
			if i2 <= i1 {
				continue
			}

			differentLetters := 0
			var diffList []int

			for j := range line1 {
				if line1[j] != line2[j] {
					differentLetters++
					diffList = append(diffList, j)
				}
			}

			if differentLetters == 1 {
				fmt.Print("Part B: ")
				for j, let := range line1 {
					if j != diffList[0] {
						fmt.Printf("%c", let)
					}
				}

			}
		}
	}

	fmt.Println()
}
