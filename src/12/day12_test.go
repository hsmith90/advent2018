package day12

import (
	"fmt"
	"helper"
	"regexp"
	"strings"
	"testing"
)

var plantByte = []byte("#")[0]
var allPots []*Pot

func TestPartA(t *testing.T) {
	lines, _ := helper.ReadLines("input")
	generations := 50000000000

	var state Row

	for i, p := range []byte(strings.TrimPrefix(strings.TrimSpace(lines[0]), "initial state: ")) {
		state.Append(&Pot{value: i, plant: p == plantByte})
	}

	var successConditions [][]bool
	regex := regexp.MustCompile("(.{5})\\s=>\\s#")

	for _, line := range lines {
		match := regex.FindSubmatch([]byte(line))

		if match != nil {
			var boolMatch []bool
			for _, m := range match[1] {
				if m == plantByte {
					boolMatch = append(boolMatch, true)
				} else {
					boolMatch = append(boolMatch, false)
				}
			}
			successConditions = append(successConditions, boolMatch)
		}
	}

	score, previousScore, diff, previousDiff := 0, 0, 0, 0
	for g := 1; g <= generations; g++ {
		score = state.UpdateState(successConditions)
		previousDiff = diff
		diff = score - previousScore
		previousScore = score

		fmt.Printf("Score: %v, Diff: %v\n", score, diff)

		if diff == previousDiff {
			score = score + ((generations - g) * diff)
			break
		}
	}

	fmt.Printf("Part A: %v\n", score)

}

type Row struct {
	firstPot, lastPot *Pot
}

type Pot struct {
	value                                int
	plant                                bool
	nextGenPlant                         bool
	previous, twoPrevious, next, twoNext *Pot
}

func (r *Row) Append(p *Pot) {
	allPots = append(allPots, p)
	if r.firstPot == nil {
		r.firstPot = p
		r.lastPot = p
		return
	}

	p.previous = r.lastPot
	p.twoPrevious = r.lastPot.previous

	r.lastPot.next = p
	if r.lastPot.previous != nil {
		r.lastPot.previous.twoNext = p
	}

	r.lastPot = p
}

func (r *Row) AppendBefore(p *Pot) {
	allPots = append(allPots, p)
	if r.firstPot == nil {
		r.firstPot = p
		r.lastPot = p
		return
	}

	p.next = r.firstPot
	p.twoNext = r.firstPot.next

	r.firstPot.previous = p
	if r.firstPot.next != nil {
		r.firstPot.next.twoPrevious = p
	}

	r.firstPot = p
}

func (p *Pot) UpdateNextGen(success [][]bool) {
	twoPrev, prev, next, twoNext := false, false, false, false
	if p.twoPrevious != nil {
		twoPrev = p.twoPrevious.plant
	}
	if p.previous != nil {
		prev = p.previous.plant
	}
	if p.next != nil {
		next = p.next.plant
	}
	if p.twoNext != nil {
		twoNext = p.twoNext.plant
	}

	for _, s := range success {
		if s[0] == twoPrev &&
			s[1] == prev &&
			s[2] == p.plant &&
			s[3] == next &&
			s[4] == twoNext {
			p.nextGenPlant = true
			return
		}
	}

	p.nextGenPlant = false
}

func (r *Row) UpdateState(success [][]bool) int {
	if r.firstPot.plant {
		for i := 1; i < 4; i++ {
			r.AppendBefore(&Pot{value: r.firstPot.value - 1, plant: false})
		}
	}

	if r.lastPot.plant {
		for i := 1; i < 4; i++ {
			r.Append(&Pot{value: r.lastPot.value + 1, plant: false})
		}
	}

	for _, p := range allPots {
		p.UpdateNextGen(success)
	}

	score := 0
	for _, p := range allPots {
		p.plant = p.nextGenPlant
		if p.plant {
			score += p.value
		}
	}

	return score

}
