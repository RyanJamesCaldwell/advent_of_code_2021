package Day1

import(
  f "adventOfCode/fileReader"
  "fmt"
  "path/filepath"
  "strings"
  "strconv"
)

func Header() {
  fmt.Println("AoC Day1 Puzzle Solution")
}

func Solve() {
  absolutePath, _ := filepath.Abs("./Day1/input.txt")
  strDataSlice := strings.Split(f.ReadFile(absolutePath), "\n")

  intDataSlice := make([]int, len(strDataSlice))
  for idx, val := range strDataSlice {
    intValue, _ := strconv.Atoi(val)
    intDataSlice[idx] = intValue
  }

  fmt.Println("Part 1 Solution:", part1(intDataSlice))
  fmt.Println("Part 2 Solution:", part2(intDataSlice))
}

func part1(dataSlice []int) int {
  count := 0
  for idx, value := range dataSlice {
    if idx == 0 {
      continue
    }

    if value > dataSlice[idx - 1] {
      count++
    }
  }

  return count
}

func part2(dataSlice []int) int {
  count := 0

  summedWindows := make([]int, len(dataSlice) - 2) // -2 to account for first block of three

  for i := 2; i < len(dataSlice); i++ {
    summedWindows[count] = sumSlice(dataSlice[i-2:i+1])
    count++
  }

  return part1(summedWindows)
}

func sumSlice(slice []int) int {
  result := 0
  for _, val := range slice {
    result += val
  }

  return result
}

