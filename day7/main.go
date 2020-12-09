package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	//"sort"
)

//Stolen - https://stackoverflow.com/a/18479916
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func main() {
	lines, _ := readLines("day7.txt")
	//bagsToFind := []string{"shiny gold"}
	bagsEventuallyContainSG := []string{}

	//Go over every line to find bags that can contain gold
	for _, line := range lines {
		line := strings.Split(line, "contain")
		if strings.Contains(line[1], "shiny gold") {
			bag := strings.Split(line[0], "bags")[0]
			bagsEventuallyContainSG = append(bagsEventuallyContainSG, bag)
		}
	}

	/*for _, line := range lines {
		line := strings.Split(line, "contain")
		for _, bag := range bagsEventuallyContainSG {
			if strings.Contains(line[1], bag) {
				bag := strings.Split(line[0], "bags")[0]
				bagsEventuallyContainSG = append(bagsEventuallyContainSG, bag)
			}
		}
	}
	fmt.Println(len(bagsEventuallyContainSG))*/

	fmt.Println(bagRecurser(bagsEventuallyContainSG, lines, []string{}, 0))
}

func bagRecurser(bags []string, lines []string, checked []string, count int) int {

	if len(bags) == 0 {
		return count
	}
	newBags := []string{}
	for _, bag := range bags {
		_, found := Find(checked, bag)
		if found {
			//We've already checked this bag
			continue
		}
		for _, line := range lines {
			line := strings.Split(line, "contain")
			if strings.Contains(line[1], bag) {
				bagName := strings.Split(line[0], " bags")[0]
				newBags = append(newBags, bagName)
				count++
				//fmt.Println(newBags)
			}
		}
		checked = append(checked, bag)
	}
	return bagRecurser(newBags, lines, checked, count)
	//return bags

}
