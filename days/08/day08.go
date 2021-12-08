package day08

import (
	"fmt"
	"math"
	"reflect"
	"sort"
	"strings"
)

type pattern struct {
	signalPatterns []string
	numberPatterns []string
	signalMap      map[string]string
	output         int
}

func Run(rawData []string) error {
	patterns, err := convertToPatterns(rawData)
	if err != nil {
		return err
	}

	fmt.Printf("Part A: %d\n", count1s4s7sAnd8s(patterns))

	makeAllMaps(patterns)

	computeAllNumbers(patterns)

	fmt.Println(computeSums(patterns))

	return nil
}

func computeSums(patterns []pattern) int {
	sum := 0
	for _, pattern := range patterns {
		sum += pattern.output
	}
	return sum
}

func computeAllNumbers(patterns []pattern) {
	for i := 0; i < len(patterns); i++ {
		patterns[i].computeNumber()
	}
}

func convertToPatterns(rawData []string) ([]pattern, error) {
	var patternList []pattern
	for _, line := range rawData {
		patternsAndNumbers := strings.Split(line, "|")
		if len(patternsAndNumbers) != 2 {
			return nil, fmt.Errorf("%w: %d", errUnexpectedInput, len(patternsAndNumbers))
		}
		patterns := strings.Split(strings.TrimSpace(patternsAndNumbers[0]), " ")
		if len(patterns) != 10 {
			return nil, fmt.Errorf("%w: %d", errUnexpectedPatternNumber, len(patterns))
		}
		numbers := strings.Split(strings.TrimSpace(patternsAndNumbers[1]), " ")
		if len(numbers) != 4 {
			return nil, fmt.Errorf("%w: %d", errUnexpectedNumberNumber, len(numbers))
		}
		patternList = append(patternList, pattern{patterns, numbers, nil, 0})
	}

	return patternList, nil
}

func count1s4s7sAnd8s(patterns []pattern) int {
	count := 0
	for _, pattern := range patterns {
		for _, number := range pattern.numberPatterns {
			length := len(number)
			switch length {
			// 1
			case 2:
				count++
			// 7
			case 3:
				count++
			// 4
			case 4:
				count++
			// 8
			case 7:
				count++
			}
		}
	}

	return count
}

func makeAllMaps(patterns []pattern) {
	for i := 0; i < len(patterns); i++ {
		patterns[i].makeMap()
	}
}

func (p *pattern) computeNumber() {
	value := 0
	for i, numberPattern := range p.numberPatterns {
		signals := projectLetters(numberPattern, p.signalMap)
		sort.Sort(sort.StringSlice(signals))
		value += computeDigit(signals) * int(math.Pow10(3-i))
	}
	p.output = value
}

func projectLetters(word string, letterMap map[string]string) []string {
	var signals []string
	letters := strings.Split(word, "")
	for _, letter := range letters {
		signals = append(signals, letterMap[letter])
	}
	return signals
}

func computeDigit(signals []string) int {
	for number, letters := range numberMap {
		if reflect.DeepEqual(signals, letters) {
			return number
		}
	}
	return -1
}

func (p *pattern) makeMap() {
	p.signalMap = make(map[string]string)
	listOfLetters := []string{"a", "b", "c", "d", "e", "f", "g"}
	for _, letter := range listOfLetters {
		var matches []int
		for _, signal := range p.signalPatterns {
			if strings.Contains(signal, letter) {
				matches = append(matches, len(signal))
			}
		}
		var matchingLetter string
		switch len(matches) {
		case 4:
			matchingLetter = "e"
		case 6:
			matchingLetter = "b"
		case 7:
			if contains(matches, 4) {
				matchingLetter = "d"
			} else {
				matchingLetter = "g"
			}
		case 8:
			if contains(matches, 2) {
				matchingLetter = "c"
			} else {
				matchingLetter = "a"
			}
		case 9:
			matchingLetter = "f"
		}
		p.signalMap[letter] = matchingLetter
	}
}

func contains(slice []int, value int) bool {
	for _, testValue := range slice {
		if testValue == value {
			return true
		}
	}

	return false
}

// Sentinel errors
var errUnexpectedInput = fmt.Errorf("unexpected number of elements in input")
var errUnexpectedPatternNumber = fmt.Errorf("unexpected number of signal patterns")
var errUnexpectedNumberNumber = fmt.Errorf("unexpected number of numberPatterns")

var numberMap = map[int][]string{
	0: {"a", "b", "c", "e", "f", "g"},
	1: {"c", "f"},
	2: {"a", "c", "d", "e", "g"},
	3: {"a", "c", "d", "f", "g"},
	4: {"b", "c", "d", "f"},
	5: {"a", "b", "d", "f", "g"},
	6: {"a", "b", "d", "e", "f", "g"},
	7: {"a", "c", "f"},
	8: {"a", "b", "c", "d", "e", "f", "g"},
	9: {"a", "b", "c", "d", "f", "g"},
}
