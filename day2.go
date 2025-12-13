package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
