package Day3

import (
	f "adventOfCode/fileReader"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
)

func Header() {
	fmt.Println("AoC Day3 Puzzle Solution")
}

func Solve() {
	absolutePath, _ := filepath.Abs("./Day3/sample_input.txt")
	strDataSlice := strings.Split(f.ReadFile(absolutePath), "\n")

	fmt.Println("Part 1 Solution:", part1(strDataSlice))
	fmt.Println("Part 2 Solution:", part2(strDataSlice))
}

func part1(data []string) int {
	gammaRate, epsilonRate := gammaRate(data), epsilonRate(data)

	return gammaRate * epsilonRate
}

func part2(data []string) int {
	oxygenGeneratorRating, co2ScrubberRating := 0, 0

	return oxygenGeneratorRating * co2ScrubberRating
}

func gammaRate(data []string) int {
	bitCountByColumn := make([]map[string]int, 12)

	for _, row := range data {
		for colIdx, colValue := range strings.Split(row, "") {
			if bitCountByColumn[colIdx] == nil {
				bitCountByColumn[colIdx] = map[string]int{"0": 0, "1": 0}
			}

			bitCountByColumn[colIdx][colValue] += 1
		}
	}

	gammaRateMap := make([]int, 12)

	for idx, colMap := range bitCountByColumn {
		if colMap["0"] > colMap["1"] {
			gammaRateMap[idx] = 0
		} else {
			gammaRateMap[idx] = 1
		}
	}

	return intSliceToDecimal(gammaRateMap)
}

func epsilonRate(data []string) int {
	bitCountByColumn := make([]map[string]int, 12)

	for _, row := range data {
		for colIdx, colValue := range strings.Split(row, "") {
			if bitCountByColumn[colIdx] == nil {
				bitCountByColumn[colIdx] = map[string]int{"0": 0, "1": 0}
			}

			bitCountByColumn[colIdx][colValue] += 1
		}
	}

	epsilonRateMap := make([]int, 12)

	for idx, colMap := range bitCountByColumn {
		if colMap["0"] > colMap["1"] {
			epsilonRateMap[idx] = 1
		} else {
			epsilonRateMap[idx] = 0
		}
	}

	return intSliceToDecimal(epsilonRateMap)
}

func intSliceToDecimal(intSlice []int) int {
	str := []string{}
	for _, val := range intSlice {
		str = append(str, strconv.Itoa(val))
	}

	builtStr := strings.Join(str, "")

	binaryInt64, _ := strconv.ParseInt(builtStr, 2, 64)
	formattedInt := strconv.FormatInt(binaryInt64, 10)

	intResult, _ := strconv.Atoi(formattedInt)

	return intResult
}
