package day5

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
	"unicode"
)

func TestPartA(t *testing.T) {
	input, _ := ioutil.ReadFile("input")
	runes := bytes.Runes(input)

	collapsedPolymer := CollapsePolymer(runes)

	fmt.Printf("Part A: %v\n", len(collapsedPolymer))
}

func CollapsePolymer(s []rune) []rune {
	var newArray []rune
	anyCollapsed := false
	for i := 0; i < len(s); i++ {
		if i+1 == len(s) {
			newArray = append(newArray, s[i])
			break
		}
		if s[i] == s[i+1] {
			newArray = append(newArray, s[i])
			continue
		}
		if unicode.ToUpper(s[i]) != unicode.ToUpper(s[i+1]) {
			newArray = append(newArray, s[i])
		} else {
			anyCollapsed = true
			i++
		}
	}

	if anyCollapsed {
		return CollapsePolymer(newArray)
	}
	return newArray
}

func TestPartB(t *testing.T) {
	input, _ := ioutil.ReadFile("input")
	runes := bytes.Runes(input)

	alpha := []rune("abcdefghijklmnopqrstuvwxyz")

	smallestPolymer := runes
	for _, r := range alpha {
		testRunes := RemoveRuneType(r, runes)
		testPolymer := CollapsePolymer(testRunes)
		if len(testPolymer) < len(smallestPolymer) {
			smallestPolymer = testPolymer
		}
	}

	fmt.Printf("Part A: %v\n", len(smallestPolymer))
}

func RemoveRuneType(r rune, runes []rune) []rune {
	var newArray []rune
	for _, c := range runes {
		if unicode.ToUpper(c) != unicode.ToUpper(r) {
			newArray = append(newArray, c)
		}
	}
	return newArray
}
