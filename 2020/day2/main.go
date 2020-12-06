package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var valid int
	for scanner.Scan() {
		line := scanner.Text()
		elements := strings.Split(line, " ")
		if len(elements) != 3 {
			panic(len(elements))
		}

		// Grab max and min
		minMax := strings.Split(elements[0], "-")
		min, err := strconv.Atoi(string(minMax[0]))
		if err != nil {
			panic(err)
		}
		max, err := strconv.Atoi(string(minMax[1]))
		if err != nil {
			panic(err)
		}

		// Grab char to search for
		char := elements[1]
		char = strings.TrimSuffix(char, ":")

		// Grab password
		password := elements[2]

		if string(password[min-1]) == char && string(password[max-1]) == char {
			continue
		}
		if string(password[min-1]) == char || string(password[max-1]) == char {
			valid++
		}
	}
	fmt.Println(valid)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
