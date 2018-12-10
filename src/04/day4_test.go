package day4

import (
	"fmt"
	"helper"
	"regexp"
	"sort"
	"strconv"
	"testing"
)

func TestPartA(t *testing.T) {
	guards := GuardList()
	var sleepiestGuard Guard
	for _, guard := range guards {
		if guard.TotalMinutesAsleep() > sleepiestGuard.TotalMinutesAsleep() {
			sleepiestGuard = guard
		}
	}

	sleepiestMinute, _ := sleepiestGuard.SleepiestMinute()
	fmt.Printf("Part A: %v\n", sleepiestGuard.id*sleepiestMinute)
}

func TestPartB(t *testing.T) {
	guards := GuardList()

	var sleepiestGuard Guard
	maxTimesAsleep := 0
	for _, guard := range guards {
		_, timesAsleep := guard.SleepiestMinute()

		if timesAsleep > maxTimesAsleep {
			sleepiestGuard = guard
			maxTimesAsleep = timesAsleep
		}
	}

	sleepiestMinute, _ := sleepiestGuard.SleepiestMinute()
	fmt.Printf("Part A: %v\n", sleepiestGuard.id*sleepiestMinute)
}

type Guard struct {
	id            int
	minutesAsleep map[int]int
}

func IsGuardDuplicate(g int, guards []Guard) (bool, Guard) {
	for _, guard := range guards {
		if g == guard.id {
			return true, guard
		}
	}
	return false, Guard{id: g, minutesAsleep: make(map[int]int)}
}

func makeRange(min, max int) []int {
	a := make([]int, max-min)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func (g *Guard) TotalMinutesAsleep() (total int) {
	for _, count := range g.minutesAsleep {
		total += count
	}
	return total
}

func (g *Guard) SleepiestMinute() (sleepiestMinute int, timesAsleep int) {
	for key, value := range g.minutesAsleep {
		if value > timesAsleep {
			timesAsleep = value
			sleepiestMinute = key
		}
	}
	return sleepiestMinute, timesAsleep
}

func GuardList() []Guard {
	lines, _ := helper.ReadLines("input")

	sort.Strings(lines)
	var guards []Guard

	var currentID int64
	var startMinute int64 = -1
	for _, line := range lines {
		idRegex := regexp.MustCompile("Guard #([0-9]+)")
		matches := idRegex.FindStringSubmatch(line)
		if matches != nil {
			currentID, _ = strconv.ParseInt(matches[1], 10, 32)
			continue
		}

		minuteRegex := regexp.MustCompile("\\[[0-9]{4}-[0-9]{2}-[0-9]{2}.[0-9]{2}:([0-9]{2})\\]")
		time := minuteRegex.FindStringSubmatch(line)
		awake, _ := regexp.MatchString("wakes up", line)

		if !awake {
			startMinute, _ = strconv.ParseInt(time[1], 10, 32)
		} else {
			endMinute, _ := strconv.ParseInt(time[1], 10, 32)
			duplicate, currentGuard := IsGuardDuplicate(int(currentID), guards)

			minutes := makeRange(int(startMinute), int(endMinute))

			if duplicate {
				for _, minute := range minutes {
					currentGuard.minutesAsleep[minute]++
				}
			} else {
				guards = append(guards, currentGuard)
				for _, minute := range minutes {
					currentGuard.minutesAsleep[minute] = 1
				}
			}

			startMinute = -1
		}
	}

	return guards
}
