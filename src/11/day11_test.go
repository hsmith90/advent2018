package day10

import (
	"errors"
	"fmt"
	"testing"
)

const gridSN = 6042

// const gridSN = 18

func TestPartA(t *testing.T) {
	maxFuelCell := FuelCell{}

	for y := 1; y <= 300; y++ {
		for x := 1; x <= 300; x++ {
			for size := 1; size < 300; size++ {
				powerLevel, err := GetSquarePowerLevel(x, y, size)
				if err != nil {
					break
				}

				if powerLevel > maxFuelCell.powerLevel {
					maxFuelCell.x = x
					maxFuelCell.y = y
					maxFuelCell.powerLevel = powerLevel
					maxFuelCell.size = size
				}
			}
		}
	}

	fmt.Printf("Part A: %v,%v,%v; TotalPower: %v\n", maxFuelCell.x, maxFuelCell.y, maxFuelCell.size, maxFuelCell.powerLevel)

}

type FuelCell struct {
	x, y       int
	powerLevel int
	size       int
}

func GetCellPowerLevel(x, y int) int {
	rackID := x + 10
	prePL := ((rackID * y) + gridSN) * rackID

	return ((prePL / 100) % 10) - 5
}

func GetSquarePowerLevel(xInit, yInit, size int) (int, error) {
	x := xInit
	y := yInit
	if x+size >= 300 || y+size >= 300 {
		return 0, errors.New("Square not completely in grid")
	}

	total := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			total += GetCellPowerLevel(x, y)
			y++
		}
		x++
		y = yInit
	}

	return total, nil
}
