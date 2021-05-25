package scanString

import (
	"regexp"
	"strings"
)

// ScanStringForReadmeItems returns any lines that have an open todo item in
// or returns 0 if none found
func ScanStringForReadmeItems(input string) []int {

	output := []int{}
	// NB not zero index as file lines start from 1 not 0
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
