package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"jujhar.com/pkg/scanstring"
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

func main() {

	var files []string
	numOfFiles := 0
	sourceDir := "/home/jujhar/proj/o/todo-markdown-todo-interface/test/testData/"

	err := filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	for _, file := range files {
		if filepath.Ext((file)) == ".md" {
			numOfFiles++
			lines := scanFiles(file)
			if len(lines) > 0 {
				log.Printf("We have the following lines %v in file %v", lines, file)
				for _, line := range lines {
					lineNumber, output, _ := getStringAtLineInFile(file, line)
					log.Printf("%v: %v", lineNumber, output)
				}
			}

		}
	}

	log.Printf("Found %v markdown files to scan", numOfFiles)
}

func scanFiles(fileName string) []int {

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("File reading error %v\n", err)
	}
	return scanstring.ScanStringForReadmeItems(string(data))
}

func getStringAtLineInFile(filename string, lineNumber int) (lineNmbr int, todoString string, err error) {
	file, err := os.Open(filename)
	lastLine := 0
	if err != nil {
		log.Fatalf("File reading error %v\n", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lastLine++
		if lastLine == lineNumber {
			// you can return sc.Bytes() if you need output in []bytes
			return lineNumber, scanner.Text(), scanner.Err()
		}
	}
	return lineNumber, "", nil

}
