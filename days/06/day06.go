package day06

import (
	"fmt"
	"strconv"
	"strings"
)

type fishByAge [9]int

func Run(rawData []string) error {
	fishByAge, err := convertToFish(rawData)
	if err != nil {
		return err
	}

	days := 80

	for i := 0; i < days; i++ {
		fishByAge = iterate(fishByAge)
	}

	fmt.Printf("%d fish after %d days\n", count(fishByAge), days)

	days = 256

	for i := 80; i < days; i++ {
		fishByAge = iterate(fishByAge)
	}

	fmt.Printf("%d fish after %d days\n", count(fishByAge), days)

	return nil
}

func convertToFish(rawData []string) (fishByAge, error) {
	var fishes fishByAge
	fishStrings := strings.Split(rawData[0], ",")
	for _, fishString := range fishStrings {
		age, err := strconv.Atoi(fishString)
		if err != nil {
			return fishByAge{}, err
		}
		fishes[age]++
	}

	return fishes, nil
}

func count(fishByAge fishByAge) int {
	numberOfFish := 0
	for _, fishCount := range fishByAge {
		numberOfFish += fishCount
	}
	return numberOfFish
}

func iterate(fishes fishByAge) fishByAge {
	var newFishes fishByAge
	for i := 0; i < len(fishes); i++ {
		if i == 0 {
			newFishes[6] += fishes[i]
			newFishes[8] += fishes[i]
		} else {
			newFishes[i-1] += fishes[i]
		}
	}
	return newFishes
}
