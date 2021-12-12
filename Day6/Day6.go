package Day6

import (
	f "adventOfCode/fileReader"
	"fmt"
  "path/filepath"
  "strconv"
	"strings"
)

func Header() {
	fmt.Println("AoC Day6 Puzzle Solution")
}

func Solve() {
	absolutePath, _ := filepath.Abs("./Day6/input.txt")
  //absolutePath, _ := filepath.Abs("./Day6/sample_input.txt")
	strDataSlice := strings.Split(f.ReadFile(absolutePath), ",")
	fishSlice := stringSliceToFishSlice(strDataSlice)
	intSlice := stringSliceToIntSlice(strDataSlice)

	fmt.Println("Part 1 solution:", part1(fishSlice, 80))
	fmt.Println("Part 2 solution:", part2(intSlice, 256))
}

type Fish struct {
	DaysUntilNewOffspring int
}

func part1(fishSlice []Fish, numDays int) int {
	for dayIter := 0; dayIter < numDays; dayIter++ {
		fishCountAtStartOfDay := len(fishSlice)
		numFishToAdd := 0

		for fishIter := 0; fishIter < fishCountAtStartOfDay; fishIter++ {
			fish := fishSlice[fishIter]

			if fish.DaysUntilNewOffspring == 0 {
				numFishToAdd += 1
				fishSlice[fishIter].DaysUntilNewOffspring = 6
			} else {
				fishSlice[fishIter].DaysUntilNewOffspring -= 1
			}
		}

		fishSlice = appendFish(fishSlice, numFishToAdd)
	}

	return len(fishSlice)
}

// sample answer: 26984457539
func part2(intSlice []int, numDays int) int {
	for currentDay := 0; currentDay < numDays; currentDay++ {
		newFishCounts := make([]int, 9)

		for idx, count := range intSlice {
			if idx == 0 {
				newFishCounts[6] += count
				newFishCounts[8] += count
			} else {
				newFishCounts[idx-1] += count
			}
		}

		intSlice = newFishCounts
	}

	sumFishCounts := 0

	for _, val := range intSlice {
		sumFishCounts += val
	}

	return sumFishCounts
}

func stringSliceToIntSlice(strSlice []string) []int {
	intSlice := make([]int, 9)

	for _, val := range strSlice {
		intValue, _ := strconv.Atoi(string(val))
		intSlice[intValue] += 1
	}

	return intSlice
}

func appendFish(fishSlice []Fish, numFishToAdd int) []Fish {
	for i := 0; i < numFishToAdd; i++ {
		fishSlice = append(fishSlice, newFish())
	}

	return fishSlice
}

func newFish() Fish {
	newFish := Fish{
		DaysUntilNewOffspring: 8,
	}

	return newFish
}


func stringSliceToFishSlice(strSlice []string) []Fish {
	fishSlice := make([]Fish, len(strSlice))

	for idx, _ := range strSlice {
		intValue, _ := strconv.Atoi(string(strSlice[idx]))
		fish := Fish{
			DaysUntilNewOffspring: intValue,
		}
		fishSlice[idx] = fish
	}

	return fishSlice
}
