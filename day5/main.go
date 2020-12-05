package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"sort"
)


//Stolen - https://stackoverflow.com/a/18479916
func readLines(path string) ([][]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines [][]string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, strings.Split(scanner.Text(), ""))
    }
    return lines, scanner.Err()
}

func main() {

	//Read file lines
	lines, _ := readLines("day5.txt")
	//Seat IDs
	seatIds := []int{}
	//Iterate
	for _, line := range lines {
		rows := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 125, 126, 127}   // Intialized with values
		cols := []int{0, 1, 2, 3, 4, 5, 6, 7}
		//First 7 elements define row
		for _, element := range line[0:7] {
			if element == "F" {
				//Split to lower
				rows = rows[0:len(rows)/2]
			} else if element == "B" {
				//Split to higher
				rows = rows[len(rows)/2:len(rows)]
			}
		}
		//Last 3 elements define col
		for _, element := range line[7:] {
			if element == "L" {
				//Split to lower
				cols = cols[0:len(cols)/2]
			} else if element == "R" {
				//Split to higher
				cols = cols[len(cols)/2:len(cols)]
			}
		}
		//Calculate seat id
		seatId := (rows[0] * 8) + cols[0]
		seatIds = append(seatIds, seatId)
	}
	//Find highest seatId
	//Sort the seats
	sort.Ints(seatIds)
	//Get largest
	fmt.Println(seatIds[len(seatIds)-1])

	//Do part 2
	
	for index, seatId := range seatIds[0:len(seatIds)-1] {
		difference := seatIds[index+1]-seatId
		if difference > 1 {
			//MISSING ONE
			fmt.Println(seatId+1)
			break
		}
	}

}