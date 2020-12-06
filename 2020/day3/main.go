package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var treeMatrix [][]int

func main() {
	// fmt.Println(step1())
	fmt.Println(step2())
}
func step1() int {
	genTreeMatrix()
	et := encountedTrees(3, 1)
	// if et != 1 {
	// 	log.Fatal("wrong number of trees", 1, et)
	// }
	return et
}
func step2() int {
	genTreeMatrix()
	et1_1 := encountedTrees(1, 1)
	if et1_1 != 4 {
		fmt.Println("wrong number of trees ", 4, et1_1)
	}
	et := et1_1
	et3_1 := encountedTrees(3, 1)
	if et3_1 != 5 {
		fmt.Println("wrong number of trees ", 5, et3_1)
	}
	et *= et3_1
	et5_1 := encountedTrees(5, 1)
	if et5_1 != 3 {
		fmt.Println("wrong number of trees ", 3, et5_1)
	}
	et *= et5_1
	et7_1 := encountedTrees(7, 1)
	if et7_1 != 2 {
		fmt.Println("wrong number of trees ", 2, et7_1)
	}
	et *= et7_1
	et1_2 := encountedTrees(1, 2)
	if et1_2 != 2 {
		fmt.Println("wrong number of trees ", 2, et1_2)
	}
	et *= et1_2
	return et
}

func genTreeMatrix() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		pattern := scanner.Text()
		var trees []int
		for _, p := range pattern {
			if string(p) == "#" {
				trees = append(trees, 1)
			} else {
				trees = append(trees, 0)
			}
		}
		treeMatrix = append(treeMatrix, trees)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
func encountedTrees(slope, angle int) int {
	treesEncountered := 0
	index := slope
	for i := 0; i < len(treeMatrix); i += angle {
		if i == 0 {
			continue
		}
		ti := index
		trees := treeMatrix[i]
		for ti >= len(trees) {
			ti -= len(trees)
		}
		if trees[ti] == 1 {
			// fmt.Println(i, ti)
			treesEncountered++
		}
		index += slope
	}

	return treesEncountered

}
