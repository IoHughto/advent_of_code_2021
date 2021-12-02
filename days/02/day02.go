package day02

import (
	"advent_of_code_2021/shared"
	"fmt"
	"strconv"
	"strings"
)

type instruction struct {
	direction string
	amount    int
}

func Run() error {
	rawData, err := shared.ReadData()
	if err != nil {
		return err
	}

	instructions, err := convertToInstructions(rawData)

	err = partA(instructions)
	if err != nil {
		return err
	}

	err = partB(instructions)

	return err
}

func convertToInstructions(rawData []string) ([]instruction, error) {
	var instructions []instruction

	for _, value := range rawData {
		words := strings.Split(value, " ")
		if len(words) != 2 {
			return nil, fmt.Errorf("%w: %d", errWrongNumberOfItems, len(words))
		}
		direction := words[0]
		if !isValidDirection(direction) {
			return nil, fmt.Errorf("%w: %s", errUnexpectedDirection, direction)
		}
		amount, err := strconv.Atoi(words[1])
		if err != nil {
			return nil, err
		}
		newInstruction := instruction{
			direction: direction,
			amount:    amount,
		}
		instructions = append(instructions, newInstruction)
	}

	return instructions, nil
}

func isValidDirection(direction string) bool {
	for _, value := range expectedDirections {
		if value == direction {
			return true
		}
	}
	return false
}

func partA(instructions []instruction) error {
	horizontal := 0
	depth := 0

	for _, value := range instructions {
		switch value.direction {
		case "forward":
			horizontal += value.amount
		case "up":
			depth += value.amount
		case "down":
			depth -= value.amount
		}
	}

	fmt.Printf("Horizontal: %d\n", horizontal)
	fmt.Printf("     Depth: %d\n", depth)
	fmt.Printf("  Combined: %d\n", horizontal*depth)

	return nil
}

func partB(instructions []instruction) error {
	horizontal := 0
	depth := 0
	aim := 0

	for _, value := range instructions {
		switch value.direction {
		case "forward":
			horizontal += value.amount
			depth += value.amount * aim
		case "up":
			aim -= value.amount
		case "down":
			aim += value.amount
		}
	}

	fmt.Printf("Horizontal: %d\n", horizontal)
	fmt.Printf("     Depth: %d\n", depth)
	fmt.Printf("  Combined: %d\n", horizontal*depth)

	return nil
}

// Literals
var expectedDirections = []string{"forward", "down", "up"}

// Sentinel errors
var errWrongNumberOfItems = fmt.Errorf("unexpected number of items in line")
var errUnexpectedDirection = fmt.Errorf("unexpected direction")
