package day01

import (
	"fmt"
	"strconv"
)

func Run(rawData []string) error {
	data, err := convertToInt(rawData)
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

	return err
}

func convertToInt(data []string) ([]int, error) {
	var intData []int

	for _, line := range data {
		value, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		intData = append(intData, value)
	}

	return intData, nil
}

func countChanges(data []int) (int, int, int) {
	var increases, decreases, sames, previousValue int

	for i, value := range data {
		if i != 0 {
			if value < previousValue {
				decreases++
			}
			if value > previousValue {
				increases++
			}
			if value == previousValue {
				sames++
			}
		}
		previousValue = value
	}
	return increases, decreases, sames
}

func getSumData(data []int) []int {
	var sumData []int
	for i, value := range data {
		if i < len(data)-2 {
			sumData = append(sumData, value)
		}
		if i > 0 && i < len(data)-1 {
			sumData[i-1] += value
		}
		if i > 1 {
			sumData[i-2] += value
		}
	}
	return sumData
}

func partA(data []int) error {
	increases, decreases, sames := countChanges(data)
	fmt.Printf("Increases: %d\n", increases)
	fmt.Printf("Decreases: %d\n", decreases)
	fmt.Printf("    Sames: %d\n", sames)
	return nil
}

func partB(data []int) error {
	sumData := getSumData(data)
	increases, decreases, sames := countChanges(sumData)
	fmt.Printf("Increases: %d\n", increases)
	fmt.Printf("Decreases: %d\n", decreases)
	fmt.Printf("    Sames: %d\n", sames)
	return nil
}
