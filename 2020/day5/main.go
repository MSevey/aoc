package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/MSevey/aoc"
)

func main() {
	// Get input
	input := aoc.InputAsString("input.txt")
	lines := strings.Split(input, "\n")
	fmt.Println(step1(lines))
	fmt.Println(step2(lines))
}

func step1(lines []string) int {
	max := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		id := uid(line)
		if id > max {
			max = id
		}
	}
	return max
}

func step2(lines []string) int {
	var seats []int
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		seats = append(seats, uid(line))
	}
	sort.Slice(seats, func(i, j int) bool {
		return seats[i] < seats[j]
	})
	for i, seat := range seats {
		if seat+1 != seats[i+1] {
			return seat + 1
		}
	}
	return 0
}

func row(str string) int {
	return binSearch(str, "F", "B", 0, 127)
}
func binSearch(str, maxChar, minChar string, min, max int) int {
	mid := newMid(min, max)
	val := 0
	elements := strings.Split(str, "")
	for i, el := range elements {
		if el == maxChar {
			max = mid - 1
			if i == len(elements)-1 {
				val = max
			}
			mid = newMid(min, max)
		}
		if el == minChar {
			min = mid
			if i == len(elements)-1 {
				val = min
			}
			mid = newMid(min, max)
		}
	}
	return val
}
func seat(str string) int {
	return binSearch(str, "L", "R", 0, 7)
}

func uid(str string) int {
	rowStr := str[:7]
	seatStr := str[7:]
	return row(rowStr)*8 + seat(seatStr)
}

func newMid(min, max int) int {
	return max - ((max - min) / int(2))
}
