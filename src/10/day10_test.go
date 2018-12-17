package day10

import (
	"fmt"
	"helper"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func TestPartA(t *testing.T) {
	points := GetInput("input")

	allAdjacent := false
	seconds := 0

	for !allAdjacent {
		seconds++
		for _, p := range points {
			p.x += p.xVelocity
			p.y += p.yVelocity
		}

		allAdjacent = true
		for _, p1 := range points {
			if !p1.IsAdjacent(points) {
				allAdjacent = false
				break
			}
		}
	}

	var xMax, xMin, yMax, yMin = points[0].x, points[0].x, points[0].y, points[0].y
	for _, p := range points {
		if p.x > xMax {
			xMax = p.x
		}
		if p.x < xMin {
			xMin = p.x
		}
		if p.y > yMax {
			yMax = p.y
		}
		if p.y < yMin {
			yMin = p.y
		}
	}

	var strBuild strings.Builder

	for y := yMin; y <= yMax; y++ {
		for x := xMin; x <= xMax; x++ {
			isPoint := false
			for _, p := range points {
				if p.x == x && p.y == y {
					isPoint = true
					strBuild.WriteString("#")
					break
				}
			}
			if !isPoint {
				strBuild.WriteString(" ")
			}
		}
		strBuild.WriteString("\n")
	}

	fmt.Print(strBuild.String())
	fmt.Printf("Seconds: %v\n", seconds)
}

type Point struct {
	x, y                 int
	xVelocity, yVelocity int
}

func GetInput(file string) []*Point {
	lines, _ := helper.ReadLines(file)
	var points []*Point

	regex := regexp.MustCompile("<(-*\\s*[0-9]+),\\s(-*\\s*[0-9]+)>")
	for _, line := range lines {
		pts := regex.FindAllStringSubmatch(line, 2)
		x, _ := strconv.Atoi(strings.TrimSpace(pts[0][1]))
		y, _ := strconv.Atoi(strings.TrimSpace(pts[0][2]))
		xVel, _ := strconv.Atoi(strings.TrimSpace(pts[1][1]))
		yVel, _ := strconv.Atoi(strings.TrimSpace(pts[1][2]))

		newPoint := Point{x: x, y: y, xVelocity: xVel, yVelocity: yVel}

		points = append(points, &newPoint)
	}

	return points
}

func (p *Point) IsAdjacent(allPoints []*Point) bool {
	for _, p2 := range allPoints {
		xDiff := p.x - p2.x
		yDiff := p.y - p2.y
		if (-1 <= xDiff && xDiff <= 1) && (-1 <= yDiff && yDiff <= 1) && !(xDiff == 0 && yDiff == 0) {
			return true
		}
	}
	return false
}
