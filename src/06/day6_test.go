package day6

import (
	"fmt"
	"helper"
	"math"
	"strconv"
	"strings"
	"testing"
)

type Point struct {
	closestLocation        int
	isBaseLocation         bool
	distanceToAllLocations int
}

func TestPartA(t *testing.T) {
	locations, lCount := GetLocations("input")
	maxDistance := 10000

	for y, line := range locations {
		for x, p := range line {
			p.FindClosestLocation(x, y, locations)

		}
	}

	largestArea := 0

	for i := 0; i < lCount; i++ {
		area, isInfinite := GetArea(i, locations)

		if !isInfinite {
			if area > largestArea {
				largestArea = area
			}
		}
	}

	areaClosestToAll := 0

	for _, line := range locations {
		for _, p := range line {
			if p.distanceToAllLocations < maxDistance {
				areaClosestToAll++
			}
		}
	}

	fmt.Printf("Part A: %v\n", largestArea)
	fmt.Printf("Part B: %v\n", areaClosestToAll)

}

func GetLocations(file string) ([][]*Point, int) {
	lines, _ := helper.ReadLines(file)
	xMax, yMax, xMin, yMin := 0, 0, math.MaxInt32, math.MaxInt32

	for _, line := range lines {
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(strings.TrimSpace(coords[0]))
		y, _ := strconv.Atoi(strings.TrimSpace(coords[1]))

		if x < xMin {
			xMin = x
		}
		if x > xMax {
			xMax = x
		}
		if y > yMax {
			yMax = y
		}
		if y < yMin {
			yMin = y
		}
	}

	grid := make([][]*Point, yMax-yMin+1)
	for i := range grid {
		grid[i] = make([]*Point, xMax-xMin+1)
	}

	for i, line := range lines {
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(strings.TrimSpace(coords[0]))
		y, _ := strconv.Atoi(strings.TrimSpace(coords[1]))

		grid[y-yMin][x-xMin] = &Point{closestLocation: i, isBaseLocation: true}
	}

	return grid, len(lines)
}

func (p *Point) FindClosestLocation(x, y int, grid [][]*Point) {
	if p == nil {
		p = &Point{}
		grid[y][x] = p
	}

	p.FindTotalDistance(x, y, grid)

	if p.isBaseLocation {
		return
	}

	minDistance, multi := math.MaxInt32, false
	var closestLocation *Point

	for yTest, line := range grid {
		for xTest, pTest := range line {
			if pTest != nil {
				if pTest.isBaseLocation {
					distance := Abs(xTest-x) + Abs(yTest-y)
					if distance < minDistance {
						minDistance = distance
						closestLocation = pTest
						multi = false
					} else if distance == minDistance {
						multi = true
					}
				}
			}
		}
	}

	if multi {
		p.closestLocation = -1
	} else {
		p.closestLocation = closestLocation.closestLocation
	}

}

func (p *Point) FindTotalDistance(x, y int, grid [][]*Point) {
	if p == nil {
		p = &Point{}
		grid[y][x] = p
	}

	for yTest, line := range grid {
		for xTest, pTest := range line {
			if pTest != nil {
				if pTest.isBaseLocation {
					p.distanceToAllLocations += Abs(x-xTest) + Abs(y-yTest)
				}
			}
		}
	}
}

func GetArea(location int, grid [][]*Point) (area int, isInfinite bool) {

	for y, line := range grid {
		for x, p := range line {
			if p.closestLocation == location {
				area++

				if x == 0 || y == 0 || x == len(line)-1 || y == len(grid)-1 {
					return area, true
				}
			}
		}
	}

	return area, false
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
