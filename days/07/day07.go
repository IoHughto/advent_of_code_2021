package day07

import (
	"advent_of_code_2021/shared"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Run() error {
	rawData, err := shared.ReadData()
	if err != nil {
		return err
	}

	data, err := convertToInts(rawData)
	if err != nil {
		return err
	}

	err = partA(data)
	if err != nil {
		return err
	}

	err = partB(data)
	if err != nil {
		return err
	}

	return nil
}

func convertToInts(data []string) ([]int, error) {
	var intData []int

	for _, line := range data {
		numberStrings := strings.Split(line, ",")
		for _, numberString := range numberStrings {
			value, err := strconv.Atoi(numberString)
			if err != nil {
				return nil, err
			}
			intData = append(intData, value)
		}
	}

	return intData, nil
}

func partA(data []int) error {
	maxPosition := findMaxPosition(data)

	minFuel := math.MaxInt
	minFuelPosition := -1
	for i := 0; i <= maxPosition; i++ {
		fuel := computeFuelCosts(data, i, func(a int) int { return a })
		if fuel < minFuel {
			minFuel = fuel
			minFuelPosition = i
		}
	}

	fmt.Printf("Part A:\n%d fuel at position %d\n", minFuel, minFuelPosition)

	return nil
}

func partB(data []int) error {
	maxPosition := findMaxPosition(data)

	minFuel := math.MaxInt
	minFuelPosition := -1
	for i := 0; i <= maxPosition; i++ {
		fuel := computeFuelCosts(data, i, cumulative)
		if fuel < minFuel {
			minFuel = fuel
			minFuelPosition = i
		}
	}

	fmt.Printf("Part A:\n%d fuel at position %d\n", minFuel, minFuelPosition)

	return nil
}

func findMaxPosition(data []int) int {
	maxDistance := 0
	for _, value := range data {
		if value > maxDistance {
			maxDistance = value
		}
	}
	return maxDistance
}

func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func computeFuelCosts(data []int, position int, fuelFunction func(int) int) int {
	fuel := 0
	for _, value := range data {
		fuel += fuelFunction(abs(position - value))
	}

	return fuel
}

func cumulative(value int) int {
	return value * (1 + value) / 2
}
