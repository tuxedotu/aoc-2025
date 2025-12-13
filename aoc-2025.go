package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Day 1 //
	// day1p1() ; works!
	// day1p2() // WIP

	// Day 2 //
	day2p1()
}

func day2p1() {
	inputFile, err := os.Open("./day2.input")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)
	var idRanges []string

	// 1. scan input file & 2. cut up data on ',' delim
	for scanner.Scan() {
		idRanges = strings.Split(scanner.Text(), ",")
	}

	// 3. loop through data & check for funky ids
	for _, idRange := range idRanges {
		fmt.Println(idRange)
	}
}
func day1p2() {
	inputFile, err := os.Open("./day1.input")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)
	const startDial = 50
	lineCount := 0
	zeroCounter := 0
	currentDial := startDial
	nextInstruction := ""

	for scanner.Scan() {
		lineCount += 1
		nextInstruction = scanner.Text()

		// log results before each line
		fmt.Printf("{{ %d }} Curr. Dial: %d; Next Instr.: %s\n", lineCount, currentDial, nextInstruction)

		// process next line
		changeValue, err := strconv.Atoi(nextInstruction[1:])
		if err != nil {
			continue
		}

		fmt.Printf("Val: %d\n", changeValue)
		if string(nextInstruction[0]) == "L" {
			currentDial = currentDial - changeValue
			if currentDial < 0 && changeValue < 100 {
				zeroCounter += 1
				fmt.Printf("Zeroes now at: %d\n", zeroCounter)
				currentDial = 100 + currentDial%100
			} else if currentDial < 0 {
				zeroCounter += (currentDial * -1) / 100
				fmt.Printf("Zeroes now at: %d\n", zeroCounter)
				currentDial = 100 + currentDial%100
			}
		} else {
			currentDial = currentDial + changeValue
			if currentDial > 99 && changeValue < 100 {
				zeroCounter += 1
				fmt.Printf("Zeroes now at: %d\n", zeroCounter)
				currentDial = currentDial % 100
			} else if currentDial > 99 {
				zeroCounter += currentDial / 100
				fmt.Printf("Zeroes now at: %d\n", zeroCounter)
				currentDial = currentDial % 100
			}
		}

		if currentDial%100 == 0 { // check for 0
			zeroCounter += 1
			fmt.Printf("Hit a straight zero! Now at: %d\n", zeroCounter)
		}
	}

	fmt.Printf("The counted zeroes, and thus the password contains: %d\n", zeroCounter)
}

func day1p1() {
	inputFile, err := os.Open("./day1.input")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)
	const startDial = 50
	lineCount := 0
	zeroCounter := 0
	currentDial := startDial
	nextInstruction := ""

	for scanner.Scan() {
		lineCount += 1
		nextInstruction = scanner.Text()

		// log results before each line
		fmt.Printf("{{ %d }} Curr. Dial: %d; Next Instr.: %s\n", lineCount, currentDial, nextInstruction)

		// process next line
		changeValue, err := strconv.Atoi(nextInstruction[1:])
		if err != nil {
			continue
		}

		if string(nextInstruction[0]) == "L" {
			currentDial = currentDial - changeValue
			if currentDial < 0 {
				currentDial = 100 + currentDial%100
			}
		} else {
			currentDial = currentDial + changeValue
			if currentDial > 99 {
				currentDial = currentDial % 100
			}
		}

		if currentDial%100 == 0 { // check for 0
			zeroCounter += 1
		}
	}

	fmt.Printf("The counted zeroes, and thus the password contains: %d\n", zeroCounter)
}
