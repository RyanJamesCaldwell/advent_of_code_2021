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
	//absolutePath, _ := filepath.Abs("./Day6/input.txt")
  absolutePath, _ := filepath.Abs("./Day6/sample_input.txt")
	strDataSlice := strings.Split(f.ReadFile(absolutePath), ",")
	fishSlice := stringSliceToFishSlice(strDataSlice)

	fmt.Println("Part 1 solution:", part1(fishSlice, 80))
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
