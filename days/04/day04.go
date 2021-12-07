package day04

import (
	"fmt"
	"github.com/fatih/color"
	"strconv"
	"strings"
)

type square struct {
	value  int
	marked bool
}

type board struct {
	values    [][]*square
	winner    bool
	finalCall int
}

var someoneHasWon = false

func Run(rawData []string) error {
	instructions, boards, err := convertToInstructionsAndBoards(rawData)
	if err != nil {
		return err
	}

	markAllBoards(instructions, boards)

	return nil
}

func convertToInstructionsAndBoards(rawData []string) ([]int, []*board, error) {
	// Convert instructions
	var instructions []int
	instructionStrings := strings.Split(rawData[0], ",")
	for _, value := range instructionStrings {
		instruction, err := strconv.Atoi(value)
		if err != nil {
			return nil, nil, err
		}
		instructions = append(instructions, instruction)
	}

	// Convert boards
	var boards []*board

	for i := 0; i < (len(rawData)-1)/6; i++ {
		var newBoardValues [][]*square
		for j := 6*i + 2; j < 6*i+7; j++ {
			var newBoardRow []*square
			values := strings.Fields(rawData[j])
			for _, value := range values {
				squareValue, err := strconv.Atoi(value)
				if err != nil {
					return nil, nil, err
				}
				newBoardRow = append(newBoardRow, &square{squareValue, false})
			}
			newBoardValues = append(newBoardValues, newBoardRow)
		}
		boards = append(boards, &board{newBoardValues, false, -1})
	}

	return instructions, boards, nil
}

func everybodyHasWon(boards []*board) bool {
	for _, board := range boards {
		if !board.winner {
			return false
		}
	}

	return true
}

func markAllBoards(instructions []int, boards []*board) {
	for _, instruction := range instructions {
		for _, board := range boards {
			if !board.winner {
				board.markValue(instruction)
			}
			if everybodyHasWon(boards) {
				fmt.Println("The last board has won!")
				board.printString()
				return
			}
		}
	}
}

func (b board) printString() {
	bold := color.New(color.FgWhite).Add(color.Bold)

	fmt.Printf("Board: %d\n", b.computeScore())
	for _, row := range b.values {
		for _, square := range row {
			if square.marked {
				bold.Printf("%2d ", square.value)
			} else {
				fmt.Printf("%2d ", square.value)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (b *board) isWinner(finalValue int) bool {
	// Check rows
	for i := 0; i < len(b.values); i++ {
		count := 0
		for j := 0; j < len(b.values[0]); j++ {
			if b.values[i][j].marked {
				count++
			}
		}
		if count == len(b.values[0]) {
			b.finalCall = finalValue
			return true
		}
	}

	// Check columns
	for i := 0; i < len(b.values[0]); i++ {
		count := 0
		for j := 0; j < len(b.values); j++ {
			if b.values[j][i].marked {
				count++
			}
		}
		if count == len(b.values[0]) {
			b.finalCall = finalValue
			return true
		}
	}

	return false
}

func (b *board) markSquare(i, j int) {
	b.values[i][j].marked = true
}

func (b *board) markValue(value int) {
	for _, row := range b.values {
		for _, square := range row {
			if square.value == value {
				square.marked = true
			}
		}
	}
	b.winner = b.isWinner(value)
	if b.winner && !someoneHasWon {
		fmt.Println("We have a winner!")
		b.printString()
		someoneHasWon = true
	}
}

func (b board) computeScore() int {
	score := 0
	for _, row := range b.values {
		for _, square := range row {
			if !square.marked {
				score += square.value
			}
		}
	}
	return score * b.finalCall
}
