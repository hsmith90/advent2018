package day6

import (
	"fmt"
	"helper"
	"math"
	"strconv"
	"strings"
	"testing"
)

func TestPartA(t *testing.T) {
	lines, _ := helper.ReadLines("test")
	var locations []Location

	for i, line := range lines {
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(strings.TrimSpace(coords[0]))
		y, _ := strconv.Atoi(strings.TrimSpace(coords[1]))

		locations = append(locations, Location{name: strconv.Itoa(i), point: Point{x: x, y: y}})
	}

	var largestX, largestY int = locations[0].point.x, locations[0].point.y
	var smallestX, smallestY int = locations[0].point.x, locations[0].point.y
	for _, l := range locations {
		if l.point.x > largestX {
			largestX = l.point.x
		}
		if l.point.y > largestY {
			largestY = l.point.y
		}
		if l.point.x < smallestX {
			smallestX = l.point.x
		}
		if l.point.y < smallestY {
			smallestY = l.point.y
		}
	}

	for i, l := range locations {
		if l.point.x == largestX || l.point.x == smallestX || l.point.y == largestY || l.point.y == smallestY {
			locations[i].infinte = true
		} else {
			locations[i].infinte = false
		}
	}

	var allPoints []Point

	for x := smallestX; x <= largestX; x++ {
		for y := smallestY; y <= largestY; y++ {
			allPoints = append(allPoints, Point{x: x, y: y})
		}
	}

	for i, p := range allPoints {
		closestDistance := math.MaxInt32
		for _, l := range locations {
			if l.point.x == p.x && l.point.y == p.y {
				allPoints[i].closestLocation = "location"
				break
			}

			distance := int(math.Abs(float64(l.point.x-p.x)) + math.Abs(float64(l.point.y-p.y)))

			if distance < closestDistance {
				closestDistance = distance
			}
		}

		if allPoints[i].closestLocation != "location" {
			for _, l := range locations {
				distance := int(math.Abs(float64(l.point.x-p.x)) + math.Abs(float64(l.point.y-p.y)))

				if distance == closestDistance {
					if allPoints[i].closestLocation != "" {
						allPoints[i].closestLocation = "multi"
					} else {
						allPoints[i].closestLocation = l.name
					}
				}
			}
		}
	}

	for _, p := range allPoints {
		if p.closestLocation == "multi" {
			continue
		}

		for i, l := range locations {
			if l.name == p.closestLocation {
				locations[i].area++
				break
			}
		}
	}

	var largestArea int
	for _, l := range locations {
		if !l.infinte {
			if l.area > largestArea {
				largestArea = l.area
			}
		}
	}

	fmt.Printf("Part A: %v\n", largestArea+1)

}

type Location struct {
	name    string
	point   Point
	infinte bool
	area    int
}

type Point struct {
	x, y            int
	closestLocation string
}
