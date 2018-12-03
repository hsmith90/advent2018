package day1

var Frequency int = 0

func NewFrequency(change int) int {
	Frequency += change
	return Frequency
}
