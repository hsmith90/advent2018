package day3

import (
	"fmt"
	"helper"
	"regexp"
	"strconv"
	"testing"
)

func TestPartA(t *testing.T) {
	lines, _ := helper.ReadLines("input")
	var claims []Claim

	for _, line := range lines {
		var newClaim Claim
		re := regexp.MustCompile("[0-9]+")

		items := re.FindAllString(line, -1)
		newClaim.id = items[0]
		xStart, _ := strconv.ParseInt(items[1], 10, 32)
		yStart, _ := strconv.ParseInt(items[2], 10, 32)
		xLength, _ := strconv.ParseInt(items[3], 10, 32)
		yLength, _ := strconv.ParseInt(items[4], 10, 32)

		for xLength > 0 {
			yl := yLength
			ys := yStart
			for yl > 0 {
				newPoint := Point{int(xStart), int(ys)}
				newClaim.Points = append(newClaim.Points, newPoint)
				ys++
				yl--
			}
			xStart++
			xLength--
		}
		claims = append(claims, newClaim)
	}

	var duplicatePoints []Point

	for i1, claim1 := range claims {
		uniqueClaim := true
		for i2, claim2 := range claims {
			if i1 == i2 {
				continue
			}

			for _, point := range claim1.Points {
				if IsPointDuplicate(point, claim2.Points) {
					uniqueClaim = false
					if !IsPointDuplicate(point, duplicatePoints) {
						duplicatePoints = append(duplicatePoints, point)
					}
				}
			}
		}
		if uniqueClaim {
			fmt.Println("Unique ID: ", claim1.id)
		}
	}

	fmt.Println("Duplicate Points: ", len(duplicatePoints))
}

type Claim struct {
	id     string
	Points []Point
}

type Point struct {
	x int
	y int
}

func IsPointDuplicate(p Point, points []Point) bool {
	for _, point := range points {
		if p.x == point.x && p.y == point.y {
			return true
		}
	}

	return false
}
