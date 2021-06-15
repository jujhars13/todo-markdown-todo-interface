package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"text/tabwriter"

	"jujhar.com/pkg/scan"
)

type fileResultsType = struct {
	fileName string
	data     []scan.LineResultsType
}

// TODO Refactor to make less ugly
// TODO report if no open todos found

// Todo is the main entrypoint
func Todo(sourceDir string) error {

	var mdFiles []string
	err := filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext((path)) == ".md" {
			mdFiles = append(mdFiles, path)
		}
		return nil
	})
	if err != nil {
		log.Print(err)
		return errors.New("could not get file")
	}
	if len(mdFiles) == 0 {
		return errors.New("no .md files found in target dir")
	}

	var overallOutput []fileResultsType
	// look over each md file, scan for todos and return results
	for _, file := range mdFiles {
		lines, err := scanFileForContent(file)
		if err != nil {
			return err
		}
		if len(lines) > 0 {
			var results []scan.LineResultsType
			for _, line := range lines {
				tmp, _ := scan.GetStringAtSpecificLineInFile(file, line)
				results = append(results, tmp)
			}
			fileResults := fileResultsType{fileName: file, data: results}
			overallOutput = append(overallOutput, fileResults)
		}
	}
	printOutput(overallOutput)
	return nil
}

func printOutput(results []fileResultsType) {

	w := tabwriter.NewWriter(os.Stdout, 1, 4, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)
	for _, fileResults := range results {
		fmt.Fprintf(w, "%v:\n", fileResults.fileName)
		for _, el := range fileResults.data {
			fmt.Fprintf(w, "    %v %v\n", el.LineNumber, el.LineString)
		}
	}
	w.Flush()
}

func scanFileForContent(fileName string) ([]int, error) {

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return scan.StringForTodoItems(string(data)), nil
}
