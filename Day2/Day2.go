package Day2

import (
	f "adventOfCode/fileReader"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
)

func Header() {
	fmt.Println("AoC Day2 Puzzle Solution")
}

func Solve() {
	absolutePath, _ := filepath.Abs("./Day2/input.txt")
	strDataSlice := strings.Split(f.ReadFile(absolutePath), "\n")

	fmt.Println("Part 1 Solution:", part1(strDataSlice))
	fmt.Println("Part 2 Solution:", part2(strDataSlice))
}

func part1(data []string) int {
	depth, horizontalPosition := 0, 0

	for _, line := range data {
		if len(line) == 0 {
			continue
		}

		lineSlice := strings.Split(line, " ")
		command := lineSlice[0]
		distance, _ := strconv.Atoi(lineSlice[1])

		switch command {
		case "forward":
			horizontalPosition += distance
		case "up":
			depth -= distance
		case "down":
			depth += distance
		}
	}

	return depth * horizontalPosition
}

func part2(data []string) int {
	depth, horizontalPosition, aim := 0, 0, 0

	for _, line := range data {
		if len(line) == 0 {
			continue
		}

		lineSlice := strings.Split(line, " ")
		command := lineSlice[0]
		distance, _ := strconv.Atoi(lineSlice[1])

		switch command {
		case "forward":
			horizontalPosition += distance
			depth += (aim * distance)
		case "up":
			aim -= distance
		case "down":
			aim += distance
		}
	}

	return depth * horizontalPosition
}
