package aoc

import "fmt"

type Day struct {
	Label string
	Part1 func()
	Part2 func()
}

func NewDay(label string, part1 func(), part2 func()) *Day {
	return &Day{
		Label: label,
		Part1: part1,
		Part2: part2,
	}
}

func (d *Day) Run() {
	fmt.Println("Start of Day", d.Label)
	fmt.Println("Part 1:")
	d.Part1()
	fmt.Println()
	fmt.Println("Part 2:")
	d.Part2()
	fmt.Println("End of Day", d.Label)
}
