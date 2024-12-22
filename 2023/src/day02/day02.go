package day02

import (
	"advent-of-code-2023/util/aoc"
	"fmt"
	"strconv"

	"github.com/dlclark/regexp2"
)

func Part01(lines []string) {
	gameR := regexp2.MustCompile(`^Game\s(\d+)`, 0)
	dieR := regexp2.MustCompile(`(\d+)\s(red|green|blue)`, 0)
	dieLimits := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	idSum := 0

	for _, line := range lines {
		gameIdM, _ := gameR.FindStringMatch(line)
		gameId, _ := strconv.Atoi(gameIdM.GroupByNumber(1).String())
		valid := true
		dieM, _ := dieR.FindStringMatch(line)
		for dieM != nil {
			color := dieM.GroupByNumber(2).String()
			cnt, _ := strconv.Atoi(dieM.GroupByNumber(1).String())

			if cnt > dieLimits[color] {
				valid = false
				break
			}
			dieM, _ = dieR.FindNextMatch(dieM)
		}

		if valid {
			idSum += gameId
		}
	}

	fmt.Println(idSum)
}

func Part02(lines []string) {
	dieR := regexp2.MustCompile(`(\d+)\s(red|green|blue)`, 0)

	sum := 0

	for _, line := range lines {
		dieMinMap := map[string]int{
			"red":   1,
			"green": 1,
			"blue":  1,
		}
		dieM, _ := dieR.FindStringMatch(line)
		for dieM != nil {
			color := dieM.GroupByNumber(2).String()
			cnt, _ := strconv.Atoi(dieM.GroupByNumber(1).String())

			if cnt > dieMinMap[color] {
				dieMinMap[color] = cnt
			}

			dieM, _ = dieR.FindNextMatch(dieM)
		}

		sum += dieMinMap["red"] * dieMinMap["green"] * dieMinMap["blue"]
	}

	fmt.Println(sum)
}

func Day02() *aoc.Day {
	return aoc.NewDay("02", "day02/input.txt", Part01, Part02)
}
