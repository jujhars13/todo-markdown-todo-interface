// Pacakge which contains funcs to scan strings
package scan

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

type LineResultsType = struct {
	lineNumber int
	lineString string
}

// ScanStringForReadmeItems returns any lines that have an open todo item in
// or returns 0 if none found
func ScanStringForReadmeItems(input string) []int {

	output := []int{}
	// NB not zero index as file line numbers start from 1
	line := 1
	temp := strings.Split(input, "\n")

	for _, item := range temp {
		findRegex := regexp.MustCompile(`\- \[ \]`)
		matches := findRegex.FindAllStringSubmatchIndex(item, -1)
		if len(matches) > 0 {
			output = append(output, line)
		}
		line++
	}

	return output
}

// GetStringAtSpecificLineInFile gets the string content of a specified line in a file
func GetStringAtSpecificLineInFile(filename string, lineNumber int) (results LineResultsType, err error) {
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
			return LineResultsType{lineNumber, scanner.Text()}, scanner.Err()
		}
	}
	return LineResultsType{lineNumber, ""}, nil
}
