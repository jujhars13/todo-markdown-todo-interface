package main

import (
	"fmt"
	"regexp"
)

// func main() {

// 	//sourceDir := "/home/jujhar/proj/o/todo-markdown-todo-interface/test/testData"

// 	//TODO check if dir exists and check if markdown files in dir
// 	//TODO for each markdown file scan for readme items and return the line numbers for any finds
// 	//TODO
// 	testString := `- [ ] todo item 11
//   - [x] todo item 12 done
//   - [x] todo item 13
//   - [x] todo item 14 done
//   - [ ] todo item 15
//   - [ ] todo item 16 done
//   `
// 	lines := ScanFileForReadmeItems(testString)
// 	if len(lines) == 0 {
// 		log.Println("no todo items in this file")
// 	} else {
// 		log.Printf("Found the following lines %v", lines)
// 	}

// }

// ScanFileForReadmeItems returns any lines that have an open todo item in
// or returns 0 if none found
func ScanFileForReadmeItems(input string) []int {

	findRegex := regexp.MustCompile(`\- \[ \]`)
	matches := findRegex.FindAllStringSubmatchIndex(input, -1)
	if len(matches) == 0 {
		return make([]int, 0)
	}

	fmt.Printf("%v", matches)

	// return make([]int, 12)
	return matches[0]
}
