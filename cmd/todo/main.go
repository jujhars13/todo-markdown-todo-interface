package main

import (
	"fmt"
	"regexp"
)

func main() {

	//sourceDir := "/home/jujhar/proj/o/todo-markdown-todo-interface/test/testData"

	//TODO check if dir exists and check if markdown files in dir
	//TODO for each markdown file scan for readme items
  //TODO 
	testString := `- [x] todo item 11
  - [ ] todo item 12 done
  - [x] todo item 13
  - [ ] todo item 14 done
  - [ ] todo item 15
  - [ ] todo item 16 done
  `
	findRegex := regexp.MustCompile(`\- \[ \]`)
	matches := findRegex.FindAllStringSubmatchIndex(testString, -1)
	fmt.Printf("%v", matches)
}

// func scanDirForMarkdownFiles(path string, d DirEntry, err error) bool {

// 	for _, f := range files {
// 		fmt.Println(f.Name())
// 	}

// 	return true
// }
