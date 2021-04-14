package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Get Input
	getInput(2)

	// Solve Problem

	// Print Solution

}

func getInput(day int) {
	path := fmt.Sprintf("inputs/day%v.txt", day)
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	switch day {
	case 1:
		day1(scanner)
	case 2:
		day2(scanner)
	default:
	}

	err = scanner.Err()
	if err != nil {
		panic(err)
	}
}

/**************************************************/
// Day 2
func day2(input *bufio.Scanner) {
	// Get string input codes
	ok := input.Scan()
	if !ok {
		panic("no scan")
	}
	codeStr := input.Text()
	codeStrSlice := strings.Split(codeStr, ",")
	// Convert slice of strings to slice of int
	var codes []int
	for _, c := range codeStrSlice {
		i, err := strconv.Atoi(c)
		if err != nil {
			panic(err)
		}
		codes = append(codes, i)
	}

	// Replace position 1 and 2
	var noun, verb int
	for n := 0; n < 100; n++ {
		noun = n
		for v := 0; v < 100; v++ {
			verb = v
			codeCopy := make([]int, len(codes))
			copy(codeCopy, codes)
			output := execute(codeCopy, n, v)
			// Return answer, value of position 0
			if output == 19690720 {
				fmt.Printf("The noun and verb to result in %v are %v and %v", output, noun, verb)
				return
			}
		}
	}
}

func execute(codes []int, noun, verb int) int {
	codes[1] = noun
	codes[2] = verb

	// Execute program
	for i := 0; i < len(codes)-4; i += 4 {
		// Get Inputs and Replace values
		input1 := codes[codes[i+1]]
		input2 := codes[codes[i+2]]
		replaceIndex := codes[i+3]
		// Determine Code
		switch codes[i] {
		case 1:
			codes[replaceIndex] = input1 + input2
		case 2:
			codes[replaceIndex] = input1 * input2
		case 99:
			break
		default:
			fmt.Println("i,n,v", i, noun, verb)
			panic("no code??")
		}
	}
	return codes[0]
}

/**************************************************/
// Day 1
func day1(input *bufio.Scanner) {
	var totalFuel int
	for input.Scan() {
		mass, err := strconv.Atoi(input.Text())
		if err != nil {
			panic(err)
		}
		fuel := calcFuelUsage(mass)
		for fuel > 0 {
			totalFuel = totalFuel + fuel
			mass = fuel
			fuel = calcFuelUsage(mass)
		}
	}

	fmt.Println("Total Fuel Usage:", totalFuel)
}
func calcFuelUsage(mass int) int {
	return mass/3 - 2
}
