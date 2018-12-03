package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"testing"
)

func TestPartA(t *testing.T) {
	lines, _ := readLines("input")

	for _, value := range lines {
		i, _ := strconv.ParseInt(value, 10, 32)
		NewFrequency(int(i))
	}

	fmt.Printf("Part A: %v \n", Frequency)
}

func TestPartB(t *testing.T) {
	Frequency = 0
	lines, _ := readLines("input")

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

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func IsUnique(value int, list []int) bool {
	for _, item := range list {
		if item == value {
			return false
		}
	}

	return true
}
