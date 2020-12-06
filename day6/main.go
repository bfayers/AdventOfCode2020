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
	lines, _ := readLines("day6.txt")

	//Combine group answers into one line each
	groupAnswers := []string{}
	for i := 0; i < len(lines); i++ {
		newLine := ""
		counter := 0
		for j := i; j < len(lines); j++ {
			//When line is blank
			if lines[j] != "" {
				counter++
			} else {
				break
			}
		}
		for k := i; k < i+counter; k++ {
			newLine += lines[k]
		}
		groupAnswers = append(groupAnswers, newLine)
		i += counter
	}

	//Count the answers per group.
	//Define the alphabet
	alphabet := strings.Split("abcdefghijklmnopqrstuvwxyz", "")
	totalSum := 0
	for _, groupAnswer := range groupAnswers {
		counter := 0
		for _, alpha := range alphabet {
			if strings.Contains(groupAnswer, alpha) {
				//Increase the counter
				counter++
			}
		}
		totalSum += counter
	}

	fmt.Println(totalSum)

	//Part 2 adaptation
	//EVERYONE must answer yes not ANYONE
	newGroupAnswers := [][]string{{}}
	for i := 0; i < len(lines); i++ {
		newLine := []string{}
		counter := 0
		for j := i; j < len(lines); j++ {
			//When line is blank
			if lines[j] != "" {
				counter++
			} else {
				break
			}
		}
		for k := i; k < i+counter; k++ {
			newLine = append(newLine, lines[k])
		}
		newGroupAnswers = append(newGroupAnswers, newLine)
		i += counter
	}
	totalSum = 0
	//Iterate over the answers array
	for _, groupAnswer := range newGroupAnswers {
		counter := 0
		//Go through the alphabet
		for _, alpha := range alphabet {
			//Check each persons answers
			count := 0
			for _, ans := range groupAnswer {
				if strings.Contains(ans, alpha) {
					count++
				}
				if count == len(groupAnswer) {
					//Enough answers for "everyone" to be accounted for
					counter++
				}
			}
		}
		totalSum += counter
	}
	fmt.Println(totalSum)

}
