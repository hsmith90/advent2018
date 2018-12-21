package day15

import (
	"helper"
	"math"
)

const (
	elfMark    = byte('E')
	goblinMark = byte('G')
	wallMark   = byte('#')
	emptyMark  = byte('.')
)

type caveMap [][]*tile

type tile struct {
	wall bool
	unit *unit
}

func importMap(file string) caveMap {
	var cm [][]*tile
	lines, _ := helper.ReadLines(file)

	for _, line := range lines {
		bytes := []byte(line)
		var mapLine []*tile

		for _, b := range bytes {
			switch b {
			case elfMark:
				mapLine = append(mapLine, &tile{wall: false, unit: &unit{race: elf, attackPower: 3, health: 200}})
			case goblinMark:
				mapLine = append(mapLine, &tile{wall: false, unit: &unit{race: goblin, attackPower: 3, health: 200}})
			case wallMark:
				mapLine = append(mapLine, &tile{wall: true})
			case emptyMark:
				mapLine = append(mapLine, &tile{wall: false})
			}
		}

		cm = append(cm, mapLine)
	}

	return cm
}

func (c *caveMap) findAllTargets(targetRace race) []coords {
	var targetList []coords

	for y, line := range *c {
		for x, tile := range line {
			if tile.unit != nil && tile.unit.race == targetRace {
				targetList = append(targetList, coords{x: x, y: y})
			}
		}
	}
	return targetList
}

func (c *caveMap) findNextStep(player coords, targetRace race) (nextStep coords, targetsRemain bool) {
	cm := *c
	targetList := c.findAllTargets(targetRace)

	if len(targetList) == 0 {
		return coords{}, false
	}

	for _, target := range targetList {
		if player.isAdjacent(target) {
			cm[player.y][player.x].unit.attack(player, c)
			return player, true
		}
	}

	shortestDistance := math.MaxInt32
	firstStep := coords{}
	for _, target := range c.findTargetRange(player, targetList) {
		step, distance := c.findShortestPath(player, target)
		if distance == shortestDistance {
			firstStep = getHigherPriority(firstStep, step)
		}
		if distance < shortestDistance {
			shortestDistance = distance
			firstStep = step
		}
	}

	return firstStep, true
}

func (c *caveMap) findTargetRange(player coords, targets []coords) []coords {
	var inRange []coords

	for _, t := range targets {
		up, err := t.up()
		if err == nil {
			testWall, testUnit := c.getTileInfo(up)
			if !testWall && testUnit == nil {
				inRange = append(inRange, up)
			}
		}
		down, err := t.down(c)
		if err == nil {
			testWall, testUnit := c.getTileInfo(down)
			if !testWall && testUnit == nil {
				inRange = append(inRange, down)
			}
		}
		left, err := t.left()
		if err == nil {
			testWall, testUnit := c.getTileInfo(left)
			if !testWall && testUnit == nil {
				inRange = append(inRange, left)
			}
		}
		right, err := t.right(*c)
		if err == nil {
			testWall, testUnit := c.getTileInfo(right)
			if !testWall && testUnit == nil {
				inRange = append(inRange, right)
			}
		}
	}
	return inRange
}

func (c caveMap) getTileInfo(tile coords) (wall bool, unit *unit) {
	tileInfo := c[tile.y][tile.x]
	return tileInfo.wall, tileInfo.unit
}

func (c *caveMap) findShortestPath(player, target coords) (firstStep coords, distance int) {
	order := c.findTestStepOrder(player, target)

	firstStep, distance, _ = c.testStepOrder(order, target, coords{}, player, 0)

	return firstStep, distance
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (c *caveMap) findTestStepOrder(player, target coords) []coords {
	dy := player.y - target.y
	dx := player.x - target.x

	if dy > 0 && dx >= 0 {
		s1, _ := player.up()
		s2, _ := player.left()
		s3, _ := player.down(c)
		s4, _ := player.right(*c)

		return []coords{s1, s2, s3, s4}
	}
	if dy > 0 && dx < 0 {
		s1, _ := player.up()
		s2, _ := player.right(*c)
		s3, _ := player.left()
		s4, _ := player.down(c)

		return []coords{s1, s2, s3, s4}
	}
	if dy <= 0 && dx < 0 {
		s1, _ := player.right(*c)
		s2, _ := player.down(c)
		s3, _ := player.left()
		s4, _ := player.up()

		return []coords{s1, s2, s3, s4}
	}
	if dy <= 0 && dx > 0 {
		s1, _ := player.left()
		s2, _ := player.down(c)
		s3, _ := player.up()
		s4, _ := player.right(*c)

		return []coords{s1, s2, s3, s4}
	}
	if dy < 0 && dx == 0 {
		s1, _ := player.down(c)
		s2, _ := player.left()
		s3, _ := player.right(*c)
		s4, _ := player.up()

		return []coords{s1, s2, s3, s4}
	}

	return []coords{}
}

func (c *caveMap) testStepOrder(order []coords, target, parent, current coords, stepNo int) (coords, int, bool) {

	for _, p := range order {
		if p.equals(target) {
			return p, stepNo, true
		}
		if p.equals(parent) {
			continue
		}
		pWall, pUnit := c.getTileInfo(p)
		if pWall || pUnit != nil {
			continue
		}
		newOrder := c.findTestStepOrder(p, target)
		goodPath := false
		_, stepNo, goodPath = c.testStepOrder(newOrder, target, current, p, stepNo)

		if goodPath {
			stepNo++
			return p, stepNo, true
		}

	}

	return coords{}, 0, false
}
