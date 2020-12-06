package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(step1())
	fmt.Println(step2())
}
func step1() int {
	// Get input
	input := inputData("input.txt")

	// Split into entries
	entrys := strings.Split(input, "\n\n")

	// Evaluate entries
	var validEntrys int
	for _, entry := range entrys {
		// Split again
		fields := strings.Fields(entry)

		var validEntry int
		for _, field := range fields {
			ft := strings.Split(field, ":")[0]
			switch ft {
			case "byr":
				validEntry++
			case "iyr":
				validEntry++
			case "eyr":
				validEntry++
			case "hgt":
				validEntry++
			case "hcl":
				validEntry++
			case "ecl":
				validEntry++
			case "pid":
				validEntry++
			default:
			}
		}
		if validEntry == 7 {
			validEntrys++
		}
	}
	return validEntrys
}

func step2() int {
	// Get input
	input := inputData("input.txt")

	// Split into entries
	entrys := strings.Split(input, "\n\n")

	// Evaluate entries
	var validEntrys int
	for _, entry := range entrys {
		// Split again
		fields := strings.Fields(entry)

		var validEntry int
		for _, field := range fields {
			ft := strings.Split(field, ":")[0]
			value := strings.Split(field, ":")[1]
			switch ft {
			case "byr":
				year, _ := strconv.Atoi(value)
				if year >= 1920 && year <= 2002 {
					validEntry++
				}
			case "iyr":
				year, _ := strconv.Atoi(value)
				if year >= 2010 && year <= 2020 {
					validEntry++
				}
			case "eyr":
				year, _ := strconv.Atoi(value)
				if year >= 2020 && year <= 2030 {
					validEntry++
				}
			case "hgt":
				ins := strings.Contains(value, "in")
				cms := strings.Contains(value, "cm")
				if ins {
					h, _ := strconv.Atoi(strings.TrimSuffix(value, "in"))
					if h >= 59 && h <= 76 {
						validEntry++
					}
				}
				if cms {
					h, _ := strconv.Atoi(strings.TrimSuffix(value, "cm"))
					if h >= 150 && h <= 193 {
						validEntry++
					}
				}
			case "hcl":
				if len(value) == 7 {
					if string(value[0]) == "#" {
						var isValid = regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString
						if isValid(value[1:]) {
							validEntry++
						}
					}
				}
			case "ecl":
				switch value {
				case "amb":
					validEntry++
				case "blu":
					validEntry++
				case "brn":
					validEntry++
				case "gry":
					validEntry++
				case "grn":
					validEntry++
				case "hzl":
					validEntry++
				case "oth":
					validEntry++
				}
			case "pid":
				if len(value) == 9 {
					validEntry++
				}
			default:
			}
		}
		if validEntry == 7 {
			validEntrys++
		}
	}
	return validEntrys
}

func inputData(filename string) string {
	bytes, _ := ioutil.ReadFile(filename)
	return string(bytes)
}
