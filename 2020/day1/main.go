package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var nums []int
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		if num > 2020 {
			continue
		}
		nums = append(nums, num)
	}
	for _, num1 := range nums {
		for _, num2 := range nums {
			for _, num3 := range nums {
				if num1+num2+num3 == 2020 {
					println(num1 * num2 * num3)
					return
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
