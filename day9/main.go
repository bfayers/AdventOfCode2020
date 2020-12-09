package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Stolen - https://stackoverflow.com/a/18479916
func readLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		lines = append(lines, val)
	}
	return lines, scanner.Err()
}

//Stolen - https://stackoverflow.com/a/45976758
func minMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

var wrongValue int

func main() {
	// rawTestInput := `35
	// 20
	// 15
	// 25
	// 47
	// 40
	// 62
	// 55
	// 65
	// 95
	// 102
	// 117
	// 150
	// 182
	// 127
	// 219
	// 299
	// 277
	// 309
	// 576`
	// testInput := strings.Split(rawTestInput, "\n")
	// input := testInput
	input, _ := readLines("day9.txt")

	//preamble := input[0:5]

	//Part 1 - checks every number (after the preamble) to ensure that it is valid.
	for i, num := range input[25:] {
		adjustedIndex := i + 25
		//value, _ := strconv.Atoi(strings.TrimSpace(num))
		//Get last 25 numbers
		lastTwentyFive := input[adjustedIndex-25:]
		//Check if two of the last five can sum to current
		found := false
		for _, lNum := range lastTwentyFive {
			//num1, _ := strconv.Atoi(strings.TrimSpace(lNum))
			for _, l2Num := range lastTwentyFive {
				//num2, _ := strconv.Atoi(strings.TrimSpace(l2Num))
				if (lNum + l2Num) == num {
					found = true
					break
				}
			}
		}
		if found {
			continue
		} else {
			wrongValue = num
			fmt.Println("Incorrect value: " + strconv.Itoa(num))
			break
		}
	}

	//Part 2 - find >2 contiguous numbers that add to wrongValue
	for i, numm := range input {
		var numsUsed []int
		total := numm
		for _, numm2 := range input[i+1:] {
			//Add numbers we're using to array for storage
			numsUsed = append(numsUsed, numm)
			numsUsed = append(numsUsed, numm2)
			total += numm2
			if total == wrongValue {
				//Add smallest and largest numbers used
				smallest, biggest := minMax(numsUsed)
				fmt.Println("Weakness: " + strconv.Itoa(smallest+biggest))
			}
		}
	}
}
