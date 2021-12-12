package Day5

import (
	f "adventOfCode/fileReader"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
)

func Header() {
	fmt.Println("AoC Day5 Puzzle Solution")
}

func Solve() {
	absolutePath, _ := filepath.Abs("./Day5/input.txt")
	//absolutePath, _ := filepath.Abs("./Day5/sample_input.txt")
	strDataSlice := strings.Split(f.ReadFile(absolutePath), "\n")

	builtLines := buildLines(strDataSlice)
	width, height := getBoardDimensions(builtLines)
	emptyDiagram := makeEmptyDiagram(width, height)

	fmt.Println("Part 1 Solution:", part1(builtLines, emptyDiagram))

	newEmpty := makeEmptyDiagram(width, height)
	fmt.Println("Part 2 Solution:", part2(builtLines, newEmpty))
}

// types
type Point struct {
	X, Y int
}

type Line struct {
	Start, End Point
}

// part 1
func part1(lines []Line, diagram [][]int) int {
	for _, line := range lines {
		if isHorizontalLine(line) {
			diagram = plotHorizontalLine(line, diagram)
		} else if isVerticalLine(line) {
			diagram = plotVerticalLine(line, diagram)
		}
	}
	return getNumPointsGreaterThanTwo(diagram)
}

// part 2
func part2(lines []Line, diagram [][]int) int {
	for _, line := range lines {
		if isHorizontalLine(line) {
			diagram = plotHorizontalLine(line, diagram)
		} else if isVerticalLine(line) {
			diagram = plotVerticalLine(line, diagram)
		} else if isDiagonalLine(line) {
			diagram = plotDiagonalLine(line, diagram)
		}
	}
	return getNumPointsGreaterThanTwo(diagram)
}

func getNumPointsGreaterThanTwo(diagram [][]int) int {
	count := 0

	for rowIdx, rowValue := range diagram {
		for colIdx, _ := range rowValue {
			if diagram[rowIdx][colIdx] >= 2 {
				count += 1
			}
		}
	}

	return count
}

func plotDiagonalLine(line Line, diagram [][]int) [][]int {
	xIncrease, yIncrease := 0, 0

	if line.Start.X < line.End.X {
		xIncrease = 1
	} else {
		xIncrease = -1
	}

	if line.Start.Y < line.End.Y {
		yIncrease = 1
	} else {
		yIncrease = -1
	}

	var x int
	var y int

	for x, y = line.Start.X, line.Start.Y; x != line.End.X; x, y = x+xIncrease, y+yIncrease {
		diagram[y][x] += 1
	}

	diagram[y][x] += 1

	return diagram
}

func plotHorizontalLine(line Line, diagram [][]int) [][]int {
	// want to work from left to right
	var actualStartPoint Point
	var actualEndPoint Point

	if line.Start.X < line.End.X {
		actualStartPoint = line.Start
		actualEndPoint = line.End
	} else {
		actualStartPoint = line.End
		actualEndPoint = line.Start
	}

	// ex: 0,9 -> 5,9
	for i := actualStartPoint.X; i <= actualEndPoint.X; i++ {
		diagram[actualStartPoint.Y][i] += 1
	}

	return diagram
}

func plotVerticalLine(line Line, diagram [][]int) [][]int {
	//want to work from top to bottom
	var actualStartPoint Point
	var actualEndPoint Point

	if line.Start.Y < line.End.Y {
		actualStartPoint = line.Start
		actualEndPoint = line.End
	} else {
		actualStartPoint = line.End
		actualEndPoint = line.Start
	}

	for i := actualStartPoint.Y; i <= actualEndPoint.Y; i++ {
		diagram[i][actualStartPoint.X] += 1
	}

	return diagram
}

func isHorizontalLine(line Line) bool {
	if line.Start.Y == line.End.Y {
		return true
	}

	return false
}

func isVerticalLine(line Line) bool {
	if line.Start.X == line.End.X {
		return true
	}

	return false
}

func isDiagonalLine(line Line) bool {
	result := line.Start.X != line.End.X && line.Start.Y != line.End.Y
	return result
}

func printDiagram(diagram [][]int) {
	for _, rowVal := range diagram {
		fmt.Println(rowVal)
	}
}

func makeEmptyDiagram(width int, height int) [][]int {
	var emptyDiagram [][]int

	for i := 0; i < height; i++ {
		emptyDiagram = append(emptyDiagram, make([]int, width))
	}

	return emptyDiagram
}

func getBoardDimensions(lines []Line) (int, int) {
	maxWidth, maxHeight := 0, 0

	for _, line := range lines {
		lineMaxWidth := max(line.Start.X, line.End.X)
		lineMaxHeight := max(line.Start.Y, line.End.Y)

		if lineMaxWidth > maxWidth {
			maxWidth = lineMaxWidth
		}

		if lineMaxHeight > maxHeight {
			maxHeight = lineMaxHeight
		}
	}

	return maxWidth + 1, maxHeight + 1
}

func max(int1 int, int2 int) int {
	if int1 > int2 {
		return int1
	} else {
		return int2
	}
}

func buildLines(data []string) []Line {
	var lines []Line

	for _, lineVal := range data {
		if len(lineVal) == 0 {
			continue
		}
		splitStr := strings.Split(lineVal, " -> ")
		lines = append(lines, *newLine(splitStr[0], splitStr[1]))
	}

	return lines
}

func newLine(rawPoint1 string, rawPoint2 string) *Line {
	var orderedIntPoints []int

	line := Line{
		Start: Point{},
		End:   Point{},
	}

	splitPoint1 := strings.Split(rawPoint1, ",")
	splitPoint2 := strings.Split(rawPoint2, ",")

	for _, strVal := range splitPoint1 {
		intVal, _ := strconv.Atoi(strVal)
		orderedIntPoints = append(orderedIntPoints, intVal)
	}

	for _, strVal := range splitPoint2 {
		intVal, _ := strconv.Atoi(strVal)
		orderedIntPoints = append(orderedIntPoints, intVal)
	}

	line.Start.X = orderedIntPoints[0]
	line.Start.Y = orderedIntPoints[1]
	line.End.X = orderedIntPoints[2]
	line.End.Y = orderedIntPoints[3]

	return &line
}
