package day13

import (
	"fmt"
	"helper"
	"testing"
)

const (
	Vertical Track = iota + 1
	Horizontal
	CurveLeft
	CurveRight
	Intersection

	Left TurnDirection = iota + 1
	Center
	Right

	North Direction = iota + 1
	East
	South
	West

	HorizontalMark = byte('-')
	VerticalMark   = byte('|')
	CurveLeftMark  = byte('\\')
	CurveRightMark = byte('/')
	IntersectMark  = byte('+')
	EastCart       = byte('>')
	WestCart       = byte('<')
	NorthCart      = byte('^')
	SouthCart      = byte('v')
)

type Track int
type TurnDirection int
type Direction int
type TrackMap [][]*MapPoint

type Cart struct {
	nextTurnDirection TurnDirection
	currentDirection  Direction
}

type MapPoint struct {
	track Track
	cart  *Cart
}

func TestPartA(t *testing.T) {
	trackMap := CreateTrackMap("input")

	continueUpdate := true
	for continueUpdate {
		continueUpdate = trackMap.UpdateMap()
	}
}

func CreateTrackMap(file string) TrackMap {
	var trackMap [][]*MapPoint
	lines, _ := helper.ReadLines(file)

	for _, line := range lines {
		bytes := []byte(line)
		var mapLine []*MapPoint

		for _, b := range bytes {
			switch b {
			case VerticalMark:
				mapLine = append(mapLine, &MapPoint{track: Vertical})
			case HorizontalMark:
				mapLine = append(mapLine, &MapPoint{track: Horizontal})
			case CurveLeftMark:
				mapLine = append(mapLine, &MapPoint{track: CurveLeft})
			case CurveRightMark:
				mapLine = append(mapLine, &MapPoint{track: CurveRight})
			case IntersectMark:
				mapLine = append(mapLine, &MapPoint{track: Intersection})
			case EastCart:
				mapLine = append(mapLine, &MapPoint{track: Horizontal, cart: &Cart{currentDirection: East, nextTurnDirection: Left}})
			case WestCart:
				mapLine = append(mapLine, &MapPoint{track: Horizontal, cart: &Cart{currentDirection: West, nextTurnDirection: Left}})
			case NorthCart:
				mapLine = append(mapLine, &MapPoint{track: Vertical, cart: &Cart{currentDirection: North, nextTurnDirection: Left}})
			case SouthCart:
				mapLine = append(mapLine, &MapPoint{track: Vertical, cart: &Cart{currentDirection: South, nextTurnDirection: Left}})
			default:
				mapLine = append(mapLine, &MapPoint{})
			}
		}

		trackMap = append(trackMap, mapLine)
	}

	return trackMap
}

func (m TrackMap) UpdateMap() bool {
	var prevCarts []*Cart
	var err error
	for y, xLine := range m {
		for x, point := range xLine {
			if point.cart != nil {
				if Contains(prevCarts, point.cart) {
					continue
				}
				switch point.cart.currentDirection {
				case East:
					err = point.cart.MoveCart(x+1, y, m)
				case West:
					err = point.cart.MoveCart(x-1, y, m)
				case North:
					err = point.cart.MoveCart(x, y-1, m)
				case South:
					err = point.cart.MoveCart(x, y+1, m)
				}
				if err != nil {
					fmt.Println(err)
				} else {
					prevCarts = append(prevCarts, point.cart)
				}
				point.cart = nil
			}
		}
	}

	var totalCarts []*Cart
	for _, xLine := range m {
		for _, point := range xLine {
			if point.cart != nil {
				totalCarts = append(totalCarts, point.cart)
			}
		}
	}

	if len(totalCarts) == 1 {
		for y, xLine := range m {
			for x, point := range xLine {
				if point.cart != nil {
					fmt.Printf("Last Cart: %v,%v\n", x, y)
				}
			}
		}
		return false
	}
	return true
}

func (c *Cart) MoveCart(x, y int, m TrackMap) error {
	if m[y][x].cart != nil {
		m[y][x].cart = nil
		return fmt.Errorf("Crash at %v,%v", x, y)
	}
	m[y][x].cart = c
	track := m[y][x].track

	switch track {
	case Intersection:
		c.ChangeDirectionOnIntersection()
	case CurveLeft, CurveRight:
		c.ChangeDirectionOnCurve(track)
	case Vertical, Horizontal:
		err := c.OnValidTrack(track, x, y)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Cart) ChangeDirectionOnCurve(t Track) {
	if t == CurveLeft {
		switch c.currentDirection {
		case South:
			c.currentDirection = East
		case East:
			c.currentDirection = South
		case West:
			c.currentDirection = North
		case North:
			c.currentDirection = West
		}
	} else {
		switch c.currentDirection {
		case South:
			c.currentDirection = West
		case East:
			c.currentDirection = North
		case West:
			c.currentDirection = South
		case North:
			c.currentDirection = East
		}
	}
}

func (c *Cart) ChangeDirectionOnIntersection() {
	if c.nextTurnDirection == Left {
		if c.currentDirection == North {
			c.currentDirection = West
		} else {
			c.currentDirection--
		}
	} else if c.nextTurnDirection == Right {
		if c.currentDirection == West {
			c.currentDirection = North
		} else {
			c.currentDirection++
		}
	}

	if c.nextTurnDirection == Right {
		c.nextTurnDirection = Left
	} else {
		c.nextTurnDirection++
	}
}

func (c *Cart) OnValidTrack(t Track, x, y int) error {
	switch c.currentDirection {
	case North, South:
		if t == Vertical {
			return nil
		}
	case East, West:
		if t == Horizontal {
			return nil
		}
	}

	return fmt.Errorf("Cart not on Valid Track at %v,%v", x, y)
}

func Contains(list []*Cart, c *Cart) bool {
	for _, l := range list {
		if c == l {
			return true
		}
	}
	return false
}
