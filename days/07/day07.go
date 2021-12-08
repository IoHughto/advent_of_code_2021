package day07

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Run(rawData []string) error {
	if len(rawData) != 1 {
		return fmt.Errorf("%w: %d", errUnexpectedData, len(rawData))
	}

	data, err := convertToInts(rawData[0])
	if err != nil {
		return err
	}

	fmt.Printf("Part A: %d\n", findMinFuel(data, simple))
	fmt.Printf("Part B: %d\n", findMinFuel(data, cumulative))

	return nil
}

func convertToInts(data string) ([]int, error) {
	var intData []int

	numberStrings := strings.Split(data, ",")
	for _, numberString := range numberStrings {
		value, err := strconv.Atoi(numberString)
		if err != nil {
			return nil, fmt.Errorf("%w: %s", errCantParseInt, numberStrings)
		}
		intData = append(intData, value)
	}

	return intData, nil
}

func findMinFuel(data []int, fuelFunction func(int) int) int {
	maxPosition := findMaxPosition(data)

	minFuel := math.MaxInt
	for i := 0; i <= maxPosition; i++ {
		fuel := computeFuelCosts(data, i, fuelFunction)
		if fuel < minFuel {
			minFuel = fuel
		}
	}

	return minFuel
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

func simple(value int) int {
	return value
}

// Sentinel errors
var errUnexpectedData = fmt.Errorf("unexpected number of lines in input data")
var errCantParseInt = fmt.Errorf("can't parse as int")
