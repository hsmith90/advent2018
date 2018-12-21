package day15

import (
	"fmt"
	"testing"
)

func TestPartA(t *testing.T) {
	cm := importMap("test")

	targetsRemain := true
	rounds := 0

	for targetsRemain {
		var movedUnits []*unit
		for y, line := range cm {
			for x, tile := range line {
				if tile.unit != nil && !contains(tile.unit, movedUnits) {
					var targetRace race
					switch tile.unit.race {
					case goblin:
						targetRace = elf
					case elf:
						targetRace = goblin
					}
					var step coords
					step, targetsRemain = cm.findNextStep(coords{x: x, y: y}, targetRace)
					cm[y][x].unit = nil
					cm[step.y][step.x].unit = tile.unit
					movedUnits = append(movedUnits, tile.unit)
				}
			}
		}
		rounds++
	}

	remainingUnits := append(cm.findAllTargets(elf), cm.findAllTargets(goblin)...)

	hpRemain := 0
	for _, pt := range remainingUnits {
		_, unit := cm.getTileInfo(pt)

		hpRemain += unit.health
	}

	fmt.Printf("Part A: rounds: %v, totalHP: %v, Total: %v\n", rounds, hpRemain, rounds*hpRemain)

}

func contains(u *unit, list []*unit) bool {
	for _, l := range list {
		if u == l {
			return true
		}
	}
	return false
}
