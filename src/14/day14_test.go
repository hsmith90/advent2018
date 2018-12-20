package day14

import (
	"fmt"
	"testing"
)

var input = 110201

type Recipe struct {
	score int
}

func TestPartA(t *testing.T) {
	allRecipes := []int{3, 7}
	elf1, elf2 := 0, 1

	for len(allRecipes) <= input+10 {
		allRecipes = MakeNewRecipes(allRecipes, &elf1, &elf2)
	}

	lastRecipes := allRecipes[input : input+10]

	fmt.Print("Part A: ")
	for _, r := range lastRecipes {
		fmt.Printf("%v", r)
	}
	fmt.Println()

}

func TestPartB(t *testing.T) {
	allRecipes := []int{3, 7}
	elf1, elf2 := 0, 1
	seqMatch := false
	recipeCount := 0
	searchArray := TurnIntToSeq(input)
	for !seqMatch {
		allRecipes = MakeNewRecipes(allRecipes, &elf1, &elf2)
		recipeCount, seqMatch = SearchForSequence(allRecipes, searchArray, recipeCount)
	}

	fmt.Printf("Part B: %v\n", recipeCount)
}

func MakeNewRecipes(recipes []int, e1, e2 *int) (updateRecipes []int) {
	recipeSum := recipes[*e1] + recipes[*e2]

	for recipeSum > -1 {
		if recipeSum > 9 {
			recipes = append(recipes, recipeSum/10)
			recipeSum = recipeSum % 10
		} else {
			recipes = append(recipes, recipeSum)
			recipeSum = -1
		}
	}

	*e1 = (*e1 + recipes[*e1] + 1) % len(recipes)
	*e2 = (*e2 + recipes[*e2] + 1) % len(recipes)

	return recipes
}

func SearchForSequence(recipes []int, seq []int, startIndex int) (int, bool) {

	if len(seq) > len(recipes) {
		return 0, false
	}

	for i := startIndex; i < len(recipes); i++ {
		match := true
		if i+len(seq) > len(recipes) {
			return i, false
		}
		for j, s := range seq {
			if recipes[i+j] != s {
				match = false
				break
			}
		}
		if match {
			return i, true
		}
	}

	return 0, false
}

func TurnIntToSeq(seq int) []int {
	var searchArray []int
	for ; seq > 0; seq /= 10 {
		searchArray = append([]int{seq % 10}, searchArray...)
	}
	return searchArray
}
