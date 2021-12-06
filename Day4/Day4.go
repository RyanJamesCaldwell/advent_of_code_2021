package Day4

import (
	f "adventOfCode/fileReader"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
)

func Header() {
  fmt.Println("AoC Day4 Puzzle Solution")
}

func Solve() {
	absolutePath, _ := filepath.Abs("./Day4/sample_input.txt")
	strDataSlice := strings.Split(f.ReadFile(absolutePath), "\n")

  fmt.Println("Part 1 solution:", part1(strDataSlice))
  fmt.Println("Part 2 solution:", part2(strDataSlice))
}

func part1(data []string) int {
  drawnNumbers := buildDrawnNumbers(data)
  boards := buildBoardsFromRawBoardData(data[2:])

  fmt.Println("Drawn numbers:", drawnNumbers)
  fmt.Println("Boards:", boards)

  printBoard(boards[0])

  return 0
}

func part2(data []string) int {
  return 0
}

func printBoard(board [][]map[string]int) {
  fmt.Println("Board")
  fmt.Println("==========")

  for _, rowValue := range board {
    rowSlice := []string{}

    for _, col := range rowValue {
      str := strconv.Itoa(col["value"])
      rowSlice = append(rowSlice, str)
    }

    fmt.Println(strings.Join(rowSlice, ", "))
  }
}

func buildBoardsFromRawBoardData(rawBoardData []string) [][][]map[string]int {
  var boards [][][]map[string]int

  for i := 0; i < len(rawBoardData) - 1; i += 5 {
    if len(rawBoardData[i]) == 0 {
      i++
    }
    board := buildBoard(rawBoardData[i:i+5])
    boards = append(boards, board)
  }

  return boards
}

func buildDrawnNumbers(input []string) []int {
  strInputSlice := strings.Split(input[0], ",")
  intInputSlice := make([]int, len(strInputSlice))

  for idx, val := range strInputSlice {
    integerVal, _ := strconv.Atoi(val)
    intInputSlice[idx] = integerVal
  }

  return intInputSlice
}

func buildBoard(input []string) [][]map[string]int {
  result := make([][]map[string]int, 5)

  for i := 0; i < 5; i++ {
    result[i] = make([]map[string]int, 5)
  }

  for rowIdx, rowValue := range input {
    fixedRow := strings.TrimSpace(strings.Replace(rowValue, "  ", " ", -1))

    for colIdx, colValue := range strings.Split(fixedRow, " ") {
      boardPositionMap := make(map[string]int)
      intValue, _ := strconv.Atoi(colValue)

      boardPositionMap["value"] = intValue
      boardPositionMap["occupied"] = 0

      result[rowIdx][colIdx] = boardPositionMap
    }
  }

  return result
}

