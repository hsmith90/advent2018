package day15

const (
	elf race = iota + 1
	goblin
)

type race int

type unit struct {
	race        race
	attackPower int
	health      int
}

func (u *unit) attack(p coords, c *caveMap) {
	cm := *c
	targetRace := goblin
	if u.race == goblin {
		targetRace = elf
	}
	up, err := p.up()
	if err == nil {
		_, testUnit := c.getTileInfo(up)
		if testUnit != nil && testUnit.race == targetRace {
			testUnit.health -= u.attackPower

			if testUnit.health <= 0 {
				cm[up.y][up.x].unit = nil
			}
			return
		}
	}

	left, err := p.left()
	if err == nil {
		_, testUnit := c.getTileInfo(left)
		if testUnit != nil && testUnit.race == targetRace {
			testUnit.health -= u.attackPower
			if testUnit.health <= 0 {
				cm[left.y][left.x].unit = nil
			}
			return
		}
	}
	right, err := p.right(cm)
	if err == nil {
		_, testUnit := c.getTileInfo(right)
		if testUnit != nil && testUnit.race == targetRace {
			testUnit.health -= u.attackPower
			if testUnit.health <= 0 {
				cm[right.y][right.x].unit = nil
			}
			return
		}
	}

	down, err := p.down(c)
	if err == nil {
		_, testUnit := c.getTileInfo(down)
		if testUnit != nil && testUnit.race == targetRace {
			testUnit.health -= u.attackPower
			if testUnit.health <= 0 {
				cm[down.y][down.x].unit = nil
			}
			return
		}
	}
}
