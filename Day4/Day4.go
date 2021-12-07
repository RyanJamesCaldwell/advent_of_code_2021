package Day4

import (
	f "adventOfCode/fileReader"
	"fmt"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
)

func Header() {
	fmt.Println("AoC Day4 Puzzle Solution")
}

func Solve() {
	absolutePath, _ := filepath.Abs("./Day4/input.txt")
	strDataSlice := strings.Split(f.ReadFile(absolutePath), "\n")

	fmt.Println("Part 1 solution:", part1(strDataSlice))
	fmt.Println("Part 2 solution:", part2(strDataSlice))
}

func part1(data []string) int {
	drawnNumbers := buildDrawnNumbers(data)
	boards := buildBoardsFromRawBoardData(data[2:])
	var firstWinningBoard [][]map[string]int
	var numberCausingWin int

out:
	for _, currentNumber := range drawnNumbers {
		for idx, _ := range boards {
			currentBoard := boards[idx]

			for rowIdx, rowVal := range currentBoard {
				for colIdx, _ := range rowVal {
					if currentBoard[rowIdx][colIdx]["value"] == currentNumber {
						currentBoard[rowIdx][colIdx]["marked"] = 1
						if boardIsWinner(currentBoard) {
							firstWinningBoard = currentBoard
							numberCausingWin = currentNumber
							break out
						}
					}
				}
			}
		}
	}

	sumUnmarkedNumbers := sumOfUnmarkedNumbers(firstWinningBoard)

	return sumUnmarkedNumbers * numberCausingWin
}

func part2(data []string) int {
	drawnNumbers := buildDrawnNumbers(data)
	boards := buildBoardsFromRawBoardData(data[2:])
	var winningBoards [][][]map[string]int
	var numbersCausingWins []int

	for _, currentNumber := range drawnNumbers {
		for boardIdx, _ := range boards {
			if boardIdx >= len(boards) {
				continue
			}

			currentBoard := boards[boardIdx]

			if boardAlreadyWon(winningBoards, currentBoard) {
				continue
			}

			for rowIdx, rowVal := range currentBoard {
				for colIdx, _ := range rowVal {
					if currentBoard[rowIdx][colIdx]["value"] == currentNumber {
						currentBoard[rowIdx][colIdx]["marked"] = 1

						if boardIsWinner(currentBoard) {
							winningBoards = append(winningBoards, currentBoard)
							numbersCausingWins = append(numbersCausingWins, currentNumber)
						}
					}
				}
			}
		}
	}

	lastBoardToWin := winningBoards[len(winningBoards)-1]
	winningNumberForLastBoardToWin := numbersCausingWins[len(numbersCausingWins)-1]

	sumUnmarkedNumbers := sumOfUnmarkedNumbers(lastBoardToWin)

	return sumUnmarkedNumbers * winningNumberForLastBoardToWin
}

func boardAlreadyWon(winningBoards [][][]map[string]int, board [][]map[string]int) bool {
	boardAlreadyWon := false

	for _, winnerBoard := range winningBoards {
		if reflect.DeepEqual(winnerBoard, board) == true {
			boardAlreadyWon = true
			break
		}
	}

	return boardAlreadyWon
}

func removeBoardFromBoardsSlice(boards [][][]map[string]int, index int) [][][]map[string]int {
	return append(boards[:index], boards[index+1:]...)
}

func sumOfUnmarkedNumbers(board [][]map[string]int) int {
	sum := 0

	for rowIdx, _ := range board {
		for colIdx, _ := range board[rowIdx] {
			if board[rowIdx][colIdx]["marked"] == 0 {
				sum += board[rowIdx][colIdx]["value"]
			}
		}
	}

	return sum
}

func boardIsWinner(board [][]map[string]int) bool {
	boardIsWinner := false

	winsByRows := boardWinsByRows(board)
	winsByCols := boardWinsByCols(board)

	if winsByRows == true || winsByCols == true {
		boardIsWinner = true
	}

	return boardIsWinner
}

func boardWinsByRows(board [][]map[string]int) bool {
	winsByRows := false

	for rowIdx, _ := range board {
		allRowNumbersMarked := allNumbersMarked(board[rowIdx])
		if allRowNumbersMarked == true {
			winsByRows = true
			break
		}
	}

	return winsByRows
}

func boardWinsByCols(board [][]map[string]int) bool {
	transposedBoard := transpose(board)

	return boardWinsByRows(transposedBoard)
}

func transpose(a [][]map[string]int) [][]map[string]int {
	newArr := make([][]map[string]int, len(a))
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			newArr[j] = append(newArr[j], a[i][j])
		}
	}
	return newArr
}

func allNumbersMarked(numbersMap []map[string]int) bool {
	allMarkedCount := 5
	actualMarkedCount := 0

	for idx, _ := range numbersMap {
		if numbersMap[idx]["marked"] == 1 {
			actualMarkedCount++
		}
	}

	return actualMarkedCount == allMarkedCount
}

func printBoard(board [][]map[string]int) {
	fmt.Println("Board")
	fmt.Println("==========")

	for _, rowValue := range board {
		rowSlice := []string{}

		for _, col := range rowValue {
			str := strconv.Itoa(col["value"]) + "(" + strconv.Itoa(col["marked"]) + ")"
			rowSlice = append(rowSlice, str)
		}

		fmt.Println(strings.Join(rowSlice, ", "))
	}
	fmt.Println("")
}

func buildBoardsFromRawBoardData(rawBoardData []string) [][][]map[string]int {
	var boards [][][]map[string]int

	for i := 0; i < len(rawBoardData)-1; i += 5 {
		if len(rawBoardData[i]) == 0 {
			i++
		}
		board := buildBoard(rawBoardData[i : i+5])
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
			boardPositionMap["marked"] = 0

			result[rowIdx][colIdx] = boardPositionMap
		}
	}

	return result
}
