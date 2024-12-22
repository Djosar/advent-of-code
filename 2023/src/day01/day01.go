package day01

import (
	"advent-of-code-2023/util/aoc"
	"advent-of-code-2023/util/parsel"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var numbersMap = map[string]int{

	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func getLines() ([]string, error) {
	dir, _ := os.Getwd()
	filename := "/src/day01/input.txt"
	return parsel.ReadFile(dir + filename)
}

func Part01() {
	lines, err := getLines()
	if err != nil {
		fmt.Println(err)
		return
	}

	var pattern = `\d`
	r := regexp.MustCompile(pattern)
	sum := 0

	for _, line := range lines {
		matches := r.FindAllString(line, -1)
		if matches != nil {
			first := matches[0]
			last := matches[len(matches)-1]
			val, err := strconv.Atoi(first + last)
			if err != nil {
				fmt.Println(err)
				continue
			}
			sum += val
		}
	}

	fmt.Println(sum)
	return
}

func Part02() {
	var p = `\d|one|two|three|four|five|six|seven|eight|nine|zero`
	lines, err := getLines()
	if err != nil {
		fmt.Println(err)
		return
	}

	sum := 0
	r := regexp.MustCompile(p)
	for _, line := range lines {
		matches := r.FindAllString(line, -1)
		if matches != nil {
			first := matches[0]
			last := matches[len(matches)-1]

			val := 10*numbersMap[first] + numbersMap[last]
			sum += val
		}
	}

	fmt.Println(sum)
	return
}

func Day01() *aoc.Day {
	return aoc.NewDay("01", Part01, Part02)
}
