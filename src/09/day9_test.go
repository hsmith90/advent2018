package day9

import (
	"fmt"
	"testing"
)

var players = [439]int{}
var lastMarble = 7130700

func TestPartA(t *testing.T) {
	ring := Ring{}
	p := 0
	for i := 0; i <= lastMarble; i++ {

		score, didScore := ring.InsertMarble(&Marble{value: i})

		if didScore {
			players[p] += score
		}

		if p < len(players)-1 {
			p++
		} else {
			p = 0
		}
	}

	highScore := 0
	for _, play := range players {
		if play > highScore {
			highScore = play
		}
	}

	fmt.Printf("Part A: %v\n", highScore)

}

type Ring struct {
	currentMarble *Marble
}

type Marble struct {
	value    int
	next     *Marble
	previous *Marble
}

func (r *Ring) InsertMarble(marble *Marble) (score int, didScore bool) {
	if r.currentMarble == nil {
		r.currentMarble = marble
		marble.previous = marble
		marble.next = marble
		return 0, false
	}

	if marble.value%23 != 0 {
		firstMarbleClockwise := r.currentMarble.next
		secondMarbleClockwise := firstMarbleClockwise.next

		firstMarbleClockwise.next = marble
		marble.previous = firstMarbleClockwise

		secondMarbleClockwise.previous = marble
		marble.next = secondMarbleClockwise

		r.currentMarble = marble
		return 0, false
	}

	for i := 7; i > 0; i-- {
		r.currentMarble = r.currentMarble.previous
	}

	score = r.currentMarble.value + marble.value

	nextMarble := r.currentMarble.next
	prevMarble := r.currentMarble.previous

	prevMarble.next = nextMarble
	nextMarble.previous = prevMarble

	r.currentMarble = nextMarble

	return score, true
}
