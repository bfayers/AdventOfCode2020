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
				fmt.Println(line)
			}
		}
	}*/
	fmt.Println(len(bagRecurser(bagsEventuallyContainSG, lines, []string{})))
}

func bagRecurser(bags []string, lines []string, prevBags []string) []string {
	fmt.Println(len(bags))

	if len(bags) == 0 {
		return prevBags
	}
	newBags := []string{}
	for _, bag := range bags {
		for _, line := range lines {
			line := strings.Split(line, "contain")
			if strings.Contains(line[1], bag) {
				bagName := strings.Split(line[0], "bags")[0]
				newBags = append(newBags, bagName)
				//fmt.Println(newBags)
			}
		}
	}
	return bagRecurser(newBags, lines, bags)
	//return bags

}
