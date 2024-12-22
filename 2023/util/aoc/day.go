package aoc

import (
	"advent-of-code-2023/util/parsel"
	"fmt"
	"os"
)

type Day struct {
	Label string
	File  []string
	Part1 func([]string)
	Part2 func([]string)
}

func getFileContent(path string) ([]string, error) {
	dir, _ := os.Getwd()
	return parsel.ReadFile(dir + "/src/" + path)
}

func NewDay(label string, path string, part1 func([]string), part2 func([]string)) *Day {
	file, err := getFileContent(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &Day{
		Label: label,
		File:  file,
		Part1: part1,
		Part2: part2,
	}
}

func (d *Day) Run() {
	fmt.Println("Start of Day", d.Label)
	fmt.Println("Part 1:")
	d.Part1(d.File)
	fmt.Println()
	fmt.Println("Part 2:")
	d.Part2(d.File)
	fmt.Println("End of Day", d.Label)
}
