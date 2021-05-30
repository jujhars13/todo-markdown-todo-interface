package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"jujhar.com/pkg/scan"
)

type FileResultsType = struct {
	fileName string
	data     []scan.LineResultsType
}

// TODO Refactor to make less ugly
// TODO report if no open todos found

// Todo is the main entrypoint
func Todo(sourceDir string) {

	var mdFiles []string
	err := filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext((path)) == ".md" {
			mdFiles = append(mdFiles, path)
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	if len(mdFiles) == 0 {
		log.Fatal("No Markdown files found in target dir")
	}

	var overallOutput []FileResultsType

	// look over each files, find markdown ones and scan them saving the results
	// TODO refactor so we're not indented in so far
	for _, file := range mdFiles {
		lines := scanFileContent(file)
		if len(lines) > 0 {
			var results []scan.LineResultsType
			for _, line := range lines {
				tmp, _ := scan.GetStringAtSpecificLineInFile(file, line)
				results = append(results, tmp)
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
	return scan.ScanStringForReadmeItems(string(data))
}
