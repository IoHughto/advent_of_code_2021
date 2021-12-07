package day03

import (
	"fmt"
	"strconv"
	"strings"
)

func Run(rawData []string) error {
	data, err := convertToIntSlice(rawData)
	if err != nil {
		return err
	}

	fmt.Println("Part A")
	err = partA(data)
	if err != nil {
		return err
	}

	fmt.Println("Part B")
	err = partB(data)
	if err != nil {
		return err
	}

	return err
}

func computeGammaAndEpsilon(sumArray []int, length int) (int64, int64, error) {
	var gammaString, epsilonString string
	for _, bit := range sumArray {
		if bit > length-bit {
			gammaString += "1"
			epsilonString += "0"
		} else {
			gammaString += "0"
			epsilonString += "1"
		}
	}

	gamma, err := strconv.ParseInt(gammaString, 2, 64)
	if err != nil {
		return 0, 0, err
	}

	epsilon, err := strconv.ParseInt(epsilonString, 2, 64)
	if err != nil {
		return 0, 0, err
	}

	return gamma, epsilon, nil
}

func convertToIntSlice(rawData []string) ([][]int, error) {
	var data [][]int

	for _, line := range rawData {
		var intSlice []int
		for _, value := range strings.Split(line, "") {
			intValue, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			if intValue != 0 && intValue != 1 {
				return nil, fmt.Errorf("%w: %d", errNotBinary, intValue)
			}
			intSlice = append(intSlice, intValue)
		}
		data = append(data, intSlice)
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("%w", errNoData)
	}

	return data, nil
}

func partA(data [][]int) error {
	sumArray := make([]int, len(data[0]))

	for _, value := range data {
		for i, bit := range value {
			sumArray[i] += bit
		}
	}

	gamma, epsilon, err := computeGammaAndEpsilon(sumArray, len(data))
	if err != nil {
		return err
	}

	fmt.Printf("  Gamma: %d\n", gamma)
	fmt.Printf("Epsilon: %d\n", epsilon)
	fmt.Printf("  Value: %d\n", gamma*epsilon)

	return nil
}

func ratingCalculator(data [][]int, override int, selectorTest func(int, int) bool) (int64, error) {
	previousList := data
	for i := range data[0] {
		// Find the most common bit
		sumTotal := 0
		for _, value := range previousList {
			sumTotal += value[i]
		}
		keeper := -1
		// Use the selectorTest function to extract the preferred value
		if selectorTest(sumTotal, len(previousList)-sumTotal) {
			// Most common is 1
			keeper = 1
		} else if sumTotal == len(previousList)-sumTotal {
			// If they're tied, use the override
			keeper = override
		} else {
			// Most common is 0
			keeper = 0
		}
		// Remove all uncommon values
		var newList [][]int
		for _, value := range previousList {
			if value[i] == keeper {
				newList = append(newList, value)
			}
		}

		// Stop when there's only one value left
		if len(newList) == 1 {
			var stringValue string
			for _, bit := range newList[0] {
				stringValue += strconv.Itoa(bit)
			}
			return strconv.ParseInt(stringValue, 2, 64)
		} else {
			previousList = newList
		}

	}

	return 0, errNoValueFound
}

func partB(data [][]int) error {
	o2Rating, err := ratingCalculator(data, 1, func(a, b int) bool { return a > b })
	if err != nil {
		return err
	}
	co2Rating, err := ratingCalculator(data, 0, func(a, b int) bool { return a < b })
	if err != nil {
		return err
	}

	fmt.Printf(" O2 Rating: %d\n", o2Rating)
	fmt.Printf("CO2 Rating: %d\n", co2Rating)
	fmt.Printf("     Value: %d\n", o2Rating*co2Rating)

	return nil
}

// Sentinel errors
var errNoData = fmt.Errorf("no data parsed from file")
var errNotBinary = fmt.Errorf("bit was not 1 or 0")
var errNoValueFound = fmt.Errorf("no rating value found")
