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
	absolutePath, _ := filepath.Abs("./Day3/input.txt")
	strDataSlice := strings.Split(f.ReadFile(absolutePath), "\n")

	//fmt.Println("Part 1 Solution:", part1(strDataSlice))
	fmt.Println("Part 2 Solution:", part2(strDataSlice))
}

func part1(data []string) int {
	gammaRate, epsilonRate := gammaRate(data), epsilonRate(data)

	return gammaRate * epsilonRate
}

func part2(data []string) int {
	dataCopy := make([]string, len(data))

	for rowIdx, rowValue := range data {
		dataCopy[rowIdx] = rowValue
	}

	oxygenGeneratorRating := oxygenGeneratorRating(data)

	co2ScrubberRating := co2ScrubberRating(dataCopy)

	return oxygenGeneratorRating * co2ScrubberRating
}

func co2ScrubberRating(data []string) int {
	//remove superfluous strings ¯\_(ツ)_/¯
	//starting data seems off
	for rowIdx, rowVal := range data {
		if len(rowVal) < 2 {
			data = removeElement(data, rowIdx)
		}
	}
	fmt.Println("Starting data", data)

	var result string
	for i := 0; i < 12; i++ {
		fmt.Println("next column iter")
		bitCountByColumn := calculateBitCountByColumn(data)
		//column has all 1 bits, the 1's are getting removed
		fmt.Println("bitCounts", bitCountByColumn[i]["0"], bitCountByColumn[i]["1"])
		if bitCountByColumn[i]["0"] < bitCountByColumn[i]["1"] || bitCountByColumn[i]["0"] == bitCountByColumn[i]["1"] {
			//remove all 1's
			for dataRowIdx, removalCount, lenData := 0, 0, len(data); dataRowIdx < lenData; dataRowIdx++ {
				j := dataRowIdx - removalCount
				if string(data[j][i]) == "1" {
					fmt.Println("removing all 1's from col", i)
					data = removeElement(data, j)
					removalCount++
				}
			}
		} else {
			//remove all 0's
			for dataRowIdx, removalCount, lenData := 0, 0, len(data); dataRowIdx < lenData; dataRowIdx++ {
				j := dataRowIdx - removalCount
				if string(data[j][i]) == "0" {
					fmt.Println("removing all 0's from col", i)
					data = removeElement(data, j)
					removalCount++
				}
			}
		}

		fmt.Println("After iteration", data)

		if len(data) == 1 {
			result = data[0]
		}
	}

	fmt.Println("Result", result)

	oxygenIntSlice := make([]int, 12) //12
	for idx, _ := range result {
		intVal, _ := strconv.Atoi(string(result[idx]))
		oxygenIntSlice[idx] = intVal
	}

	fmt.Println("co2ScrubberSlice", oxygenIntSlice)

	return intSliceToDecimal(oxygenIntSlice)
}

func oxygenGeneratorRating(data []string) int {
	//remove superfluous strings ¯\_(ツ)_/¯
	for rowIdx, rowVal := range data {
		if len(rowVal) < 2 {
			data = removeElement(data, rowIdx)
		}
	}
	var result string
	for i := 0; i < 12; i++ {
		bitCountByColumn := calculateBitCountByColumn(data)
		if bitCountByColumn[i]["0"] > bitCountByColumn[i]["1"] {
			//remove all 1's
			for dataRowIdx, removalCount, lenData := 0, 0, len(data); dataRowIdx < lenData; dataRowIdx++ {
				j := dataRowIdx - removalCount
				if string(data[j][i]) == "1" {
					data = removeElement(data, j)
					removalCount++
				}
			}
		} else {
			//remove all 0's
			for dataRowIdx, removalCount, lenData := 0, 0, len(data); dataRowIdx < lenData; dataRowIdx++ {
				j := dataRowIdx - removalCount
				if string(data[j][i]) == "0" {
					data = removeElement(data, j)
					removalCount++
				}
			}
		}

		if len(data) == 1 {
			result = data[0]
		}
	}

	oxygenIntSlice := make([]int, 12) //12
	for idx, _ := range result {
		intVal, _ := strconv.Atoi(string(result[idx]))
		oxygenIntSlice[idx] = intVal
	}

	fmt.Println("oxygenIntSlice", oxygenIntSlice)

	return intSliceToDecimal(oxygenIntSlice)
}

func calculateBitCountByColumn(data []string) []map[string]int {
	bitCountByColumn := make([]map[string]int, 12) //12

	for _, row := range data {
		for colIdx, colValue := range strings.Split(row, "") {
			if bitCountByColumn[colIdx] == nil {
				bitCountByColumn[colIdx] = map[string]int{"0": 0, "1": 0}
			}

			bitCountByColumn[colIdx][colValue] += 1
		}
	}

	return bitCountByColumn
}

func removeElement(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
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
