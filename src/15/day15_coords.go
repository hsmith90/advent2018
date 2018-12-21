package day15

import "errors"

type coords struct {
	x, y int
}

func (c *coords) up() (coords, error) {
	if c.y != 0 {
		return coords{x: c.x, y: c.y - 1}, nil
	}
	return coords{}, errors.New("Not a valid point")
}

func (c *coords) down(cm *caveMap) (coords, error) {
	if c.y < len(*cm)-1 {
		return coords{x: c.x, y: c.y + 1}, nil
	}
	return coords{}, errors.New("Not a valid point")
}

func (c *coords) left() (coords, error) {
	if c.x != 0 {
		return coords{x: c.x - 1, y: c.y}, nil
	}
	return coords{}, errors.New("Not a valid point")
}

func (c *coords) right(cm caveMap) (coords, error) {
	if c.x < len(cm[0])-1 {
		return coords{x: c.x + 1, y: c.y}, nil
	}
	return coords{}, errors.New("Not a valid point")
}

func (c *coords) isAdjacent(test coords) bool {
	if (c.x == test.x && (test.y == c.y+1 || test.y == c.y-1)) ||
		(c.y == test.y && (test.x == c.x+1 || test.x == c.x-1)) {
		return true
	}
	return false
}

func (c *coords) equals(test coords) bool {
	if c.x == test.x && c.y == test.y {
		return true
	}
	return false
}

func getHigherPriority(point1, point2 coords) coords {
	if point1.equals(point2) {
		return point1
	}
	if point1.y < point2.y {
		return point1
	}
	if point1.y == point2.y && point1.x < point2.x {
		return point1
	}
	return point2
}
