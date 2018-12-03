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

	fmt.Println(Frequency)
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
