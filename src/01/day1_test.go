package day1

import (
	"fmt"
	"strconv"
	"testing"

	"helper"
)

func TestPartA(t *testing.T) {
	lines, _ := helper.ReadLines("input")

	for _, value := range lines {
		i, _ := strconv.ParseInt(value, 10, 32)
		NewFrequency(int(i))
	}

	fmt.Printf("Part A: %v \n", Frequency)
}

func TestPartB(t *testing.T) {
	Frequency = 0
	lines, _ := helper.ReadLines("input")

	var frequencies []int
	frequencies = append(frequencies, Frequency)

	for true {
		for _, value := range lines {
			i, _ := strconv.ParseInt(value, 10, 32)
			NewFrequency(int(i))
			if IsUnique(Frequency, frequencies) {
				frequencies = append(frequencies, Frequency)
			} else {
				fmt.Printf("Part B: %v \n", Frequency)
				return
			}
		}
	}
}

func IsUnique(value int, list []int) bool {
	for _, item := range list {
		if item == value {
			return false
		}
	}

	return true
}
