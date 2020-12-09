package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func find(slice []int, val int) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

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

func runBootCode(bootCode []string) (int, int) {
	accumulator := 0
	i := 0
	alreadyRun := []int{}
	for i < len(bootCode) {
		//Check if instruction has run before
		//fmt.Println(i)
		_, found := find(alreadyRun, i)
		if found {
			//Already ran this instruction
			return i, accumulator
		}

		instruction := strings.TrimSpace(bootCode[i])

		action := instruction[0:3]
		value := instruction[4:]
		//Check action
		switch action {
		case "acc":
			value, _ := strconv.Atoi(value)
			accumulator += value
			alreadyRun = append(alreadyRun, i)
			i++
		case "jmp":
			value, _ := strconv.Atoi(value)
			alreadyRun = append(alreadyRun, i)
			i += value
		case "nop":
			i++
		}
	}
	return i, accumulator
}

func main() {
	/*rawBootCode := `nop +0
	acc +1
	jmp +4
	acc +3
	jmp -3
	acc -99
	acc +1
	jmp -4
	acc +6`
	bootCode := strings.Split(rawBootCode, "\n")*/
	bootCode, _ := readLines("day8.txt")

	//Part 1
	alreadyRun := []int{}
	//Read every instruction
	accumulator := 0
	i := 0
	for i < len(bootCode) {
		//Check if instruction has run before
		//fmt.Println(i)
		_, found := find(alreadyRun, i)
		if found {
			//Already ran this instruction - print accumulator
			fmt.Println(accumulator)
			break
		}

		instruction := strings.TrimSpace(bootCode[i])

		action := instruction[0:3]
		value := instruction[4:]
		//Check action
		switch action {
		case "acc":
			value, _ := strconv.Atoi(value)
			accumulator += value
			alreadyRun = append(alreadyRun, i)
			i++
		case "jmp":
			value, _ := strconv.Atoi(value)
			alreadyRun = append(alreadyRun, i)
			i += value
		case "nop":
			i++
		}
	}

	//Part 2
	//Find every jmp index and every nop index
	var jmps []int
	var nops []int
	for i, line := range bootCode {
		if strings.Contains(line, "jmp") {
			jmps = append(jmps, i)
		} else if strings.Contains(line, "nop") {
			nops = append(nops, i)
		}
	}

	//Try every different jmp->nop
	for _, jmp := range jmps {
		//Generate new bootcode with jmp changed to nop
		//Apparently doing newBootCode := bootCode makes it a reference?!
		newBootCode := make([]string, len(bootCode))
		copy(newBootCode, bootCode)
		newLine := strings.Replace(newBootCode[jmp], "jmp", "nop", -1)
		newBootCode[jmp] = newLine
		index, accumulator := runBootCode(newBootCode)
		if index == len(newBootCode) {
			fmt.Println(accumulator)
		}
	}
	//Try every different nop->jmp
	for _, nop := range nops {
		//Generate new bootcode with jmp changed to nop
		//Apparently doing newBootCode := bootCode makes it a reference?!
		newBootCode := make([]string, len(bootCode))
		copy(newBootCode, bootCode)
		newLine := strings.Replace(newBootCode[nop], "nop", "jmp", -1)
		newBootCode[nop] = newLine
		index, accumulator := runBootCode(newBootCode)
		if index == len(newBootCode) {
			fmt.Println(accumulator)
		}
	}
}
