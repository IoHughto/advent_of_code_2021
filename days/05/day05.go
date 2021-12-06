package day05

import (
	"advent_of_code_2021/shared"
	"fmt"
)

type point struct {
	x int
	y int
}

type segment struct {
	start point
	end   point
}

type pipeMap [][]int

func Run() error {
	rawData, err := shared.ReadData()
	if err != nil {
		return err
	}

	segments, err := convertToSegments(rawData)
	if err != nil {
		return err
	}

	max := findMaximumPoint(segments)

	pipeMap := initializePipeMap(max)

	partA(segments, &pipeMap)
	partB(segments, &pipeMap)

	return nil
}

func partA(segments []segment, pipeMap *pipeMap) {
	for _, segment := range segments {
		if segment.isFlat() {
			pipeMap.markSegment(segment)
		}
	}

	fmt.Printf("Part A: %d\n", pipeMap.countCrossingsOverValue(2))
}

func partB(segments []segment, pipeMap *pipeMap) {
	for _, segment := range segments {
		if !segment.isFlat() {
			pipeMap.markSegment(segment)
		}
	}

	fmt.Printf("Part B: %d\n", pipeMap.countCrossingsOverValue(2))
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func convertToSegments(rawData []string) ([]segment, error) {
	var segments []segment
	for _, line := range rawData {
		var x1, y1, x2, y2 int
		n, err := fmt.Sscanf(line, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		if err != nil {
			return nil, err
		}
		if n != 4 {
			return nil, fmt.Errorf("%w: %s", errUnexpectedArguments, n)
		}
		newSegment := segment{
			start: point{x1, y1},
			end:   point{x2, y2},
		}
		segments = append(segments, newSegment)
	}

	return segments, nil
}

func findMaximumPoint(segments []segment) point {
	var maxX, maxY int

	for _, segment := range segments {
		if segment.findMaximumPoint().x > maxX {
			maxX = segment.findMaximumPoint().x
		}
		if segment.findMaximumPoint().y > maxY {
			maxY = segment.findMaximumPoint().y
		}
	}

	return point{maxX, maxY}
}

func initializePipeMap(max point) pipeMap {
	pipeMap := make([][]int, max.y+1)
	for y := range pipeMap {
		pipeMap[y] = make([]int, max.x+1)
	}
	return pipeMap
}

func (s segment) findMaximumPoint() point {
	var x, y int
	if s.start.x > s.end.x {
		x = s.start.x
	} else {
		x = s.end.x
	}
	if s.start.y > s.end.y {
		y = s.start.y
	} else {
		y = s.end.y
	}

	return point{x, y}
}

func (s segment) isFlat() bool {
	return s.start.x == s.end.x || s.start.y == s.end.y
}

func (s segment) slope() point {
	var x, y int
	if s.end.x != s.start.x {
		x = (s.end.x - s.start.x) / abs(s.end.x-s.start.x)
	}
	if s.end.y != s.start.y {
		y = (s.end.y - s.start.y) / abs(s.end.y-s.start.y)
	}
	return point{x, y}
}

func (s segment) length() int {
	if s.start.x != s.end.x {
		return abs(s.end.x - s.start.x)
	} else if s.start.y != s.end.y {
		return abs(s.end.y - s.start.y)
	}
	return 0
}

func (p pipeMap) markSegment(segment segment) {
	for currentPoint := segment.start; currentPoint != segment.end; currentPoint = currentPoint.plus(segment.slope()) {
		p[currentPoint.y][currentPoint.x]++
	}
	p[segment.end.y][segment.end.x]++
}

func (p pipeMap) prettyPrint() {
	for _, row := range p {
		for _, cell := range row {
			if cell != 0 {
				fmt.Printf("%d ", cell)
			} else {
				fmt.Printf(". ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (p pipeMap) findMaximumCrossings() int {
	var max int
	for _, row := range p {
		for _, cell := range row {
			if cell > max {
				max = cell
			}
		}
	}

	return max
}

func (p pipeMap) countCrossingsOverValue(value int) int {
	count := 0
	for _, row := range p {
		for _, cell := range row {
			if cell >= value {
				count++
			}
		}
	}

	return count
}

func (p point) plus(p2 point) point {
	return point{
		x: p.x + p2.x,
		y: p.y + p2.y,
	}
}

// Sentinel errors
var errUnexpectedArguments = fmt.Errorf("unexpected number of arguments in line")
