package aoc

import (
	"io/ioutil"
	"log"
)

// InputAsString reads the data from the input file and returns it as a string
func InputAsString(filename string) string {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}
