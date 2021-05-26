package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"jujhar.com/pkg/scanstring"
)

type LineResultsType = struct {
	lineNumber int
	lineString string
}
type FileResultsType = struct {
	fileName string
	data     []LineResultsType
}

// TODO Refactor to make less ugly
// TODO report if no open todos found

// Todo is the main entrypoint
func Todo(sourceDir string) {

	var files []string
	err := filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	var overallOutput []FileResultsType

	var mdFiles []string
	for _, file := range files {
		if filepath.Ext((file)) == ".md" {
			mdFiles = append(mdFiles, file)
		}
	}
	if len(mdFiles) == 0 {
		log.Fatal("No Markdown files found in target dir")
	}

	// look over each files, find markdown ones and scan them saving the results
	// TODO refactor so we're not indented in so far
	for _, file := range mdFiles {
		lines := scanFileContent(file)
		if len(lines) > 0 {
			var results []LineResultsType
			for _, line := range lines {
				lineNumber, retString, _ := getStringAtSpecificLineInFile(file, line)
				results = append(results, LineResultsType{lineNumber: lineNumber, lineString: retString})
			}
			fileResults := FileResultsType{fileName: file, data: results}
			overallOutput = append(overallOutput, fileResults)
		}
	}
	log.Printf("Results %v\n", overallOutput)
}

func scanFileContent(fileName string) []int {

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("File reading error %v\n", err)
	}
	return scanstring.ScanStringForReadmeItems(string(data))
}

func getStringAtSpecificLineInFile(filename string, lineNumber int) (lineNmbr int, todoString string, err error) {
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
